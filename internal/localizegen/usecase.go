package localizegen

import (
	"github.com/prongbang/localizegen/internal/pkg/android"
	"github.com/prongbang/localizegen/internal/pkg/flutter"
	"github.com/prongbang/localizegen/internal/pkg/ios"
	"github.com/prongbang/localizegen/pkg/core"
	"github.com/prongbang/localizegen/pkg/csvx"
	"github.com/prongbang/localizegen/pkg/platform"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

type UseCase interface {
	Process(param Localization) chan core.Result
	GetAvailableLanguages(csv csvx.CsvList) core.Languages
	GenerateResource(csv csvx.CsvList, param Localization) chan core.Result
}

type usecase struct {
	Repo      Repository
	AndroidUc android.UseCase
	IosUc     ios.UseCase
	FlutterUc flutter.UseCase
}

func (u *usecase) Process(param Localization) chan core.Result {
	state := make(chan core.Result, 0)
	rs := u.Repo.GetLocalize(param.DocumentID, param.SheetID)
	if rs.Code == http.StatusOK {
		csv := u.Repo.ReadCSV(string(rs.Data))
		state = u.GenerateResource(csv, param)
	}
	return state
}

func (u *usecase) GenerateResource(csv csvx.CsvList, param Localization) chan core.Result {
	languages := u.GetAvailableLanguages(csv)

	chanLen := 0
	if param.Locale != "" {
		chanLen = 1
	} else {
		chanLen = len(languages)
	}
	if param.Platform == platform.Ios {
		chanLen += 1
	}

	var wg sync.WaitGroup
	state := make(chan core.Result, chanLen)

	config := core.Configuration{
		Csv:       csv,
		Locale:    param.Locale,
		Target:    param.Target,
		Filename:  param.Filename,
		Languages: languages,
		Result:    state,
		OnAdd:     func() { wg.Add(1) },
		OnDone:    func() { wg.Done() },
	}
	switch param.Platform {
	case platform.Android:
		u.AndroidUc.Generate(config)
		break
	case platform.Ios:
		u.IosUc.Generate(config)
		break
	case platform.Flutter:
		u.FlutterUc.Generate(config)
		break
	default:
		state <- core.Result{Message: "Platform " + param.Platform + " not supported.", Status: "Error"}
	}

	go func() {
		defer close(state)
		wg.Wait()
	}()

	return state
}

func (u *usecase) GetAvailableLanguages(csv csvx.CsvList) core.Languages {
	languages := core.Languages{}
	for i := 1; i < len(csv[0]); i++ {
		name, err := regexp.Compile("\\w+ \\(")
		key, err := regexp.Compile("\\((\\w\\w*)\\)")
		if err == nil {
			matchName := name.FindString(csv[0][i])
			matchKey := key.FindString(csv[0][i])

			// Match: xx
			if matchKey == "" || matchName == "" {
				matchName = csv[0][i]
				matchKey = csv[0][i]
			}

			// Match: Name (xx)
			if matchKey != "" {
				matchName = strings.ReplaceAll(matchName, " (", "")
				matchKey = strings.ReplaceAll(matchKey, "(", "")
				matchKey = strings.ReplaceAll(matchKey, ")", "")
				languages[matchKey] = core.Language{
					Name:  matchName,
					Key:   matchKey,
					Index: i,
				}
			}
		}
	}
	return languages
}

func NewUseCase(
	repo Repository,
	androidUc android.UseCase,
	iosUc ios.UseCase,
	flutterUc flutter.UseCase,
) UseCase {
	return &usecase{
		Repo:      repo,
		AndroidUc: androidUc,
		IosUc:     iosUc,
		FlutterUc: flutterUc,
	}
}
