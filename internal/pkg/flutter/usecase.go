package flutter

import (
	"github.com/prongbang/localizegen/pkg/core"
)

type UseCase interface {
	Generate(configuration core.Configuration)
}

type useCase struct {
	Repo Repository
}

func (u *useCase) Generate(configuration core.Configuration) {
	filenameKeys := "keys_localizations.dart"
	filenameSources := "sources_localizations.dart"
	if configuration.Locale != "" {
		localeIndex := configuration.Languages[configuration.Locale].Index
		if localeIndex > 0 {
			configuration.OnAdd()
			go func() {
				defer configuration.OnDone()
				configuration.Result <- u.Repo.GenerateSourcesFile(configuration.Csv, configuration.Languages, configuration.Target, filenameSources)
				configuration.Result <- u.Repo.GenerateKeysFile(configuration.Csv, configuration.Languages, configuration.Target, filenameKeys)
			}()
		} else {
			configuration.Result <- core.Result{Message: "Language key " + configuration.Locale + " not found.", Status: "Error"}
		}
	} else {
		configuration.OnAdd()
		go func() {
			defer configuration.OnDone()
			configuration.Result <- u.Repo.GenerateKeysFile(configuration.Csv, configuration.Languages, configuration.Target, filenameKeys)
			configuration.Result <- u.Repo.GenerateSourcesFile(configuration.Csv, configuration.Languages, configuration.Target, filenameSources)
		}()
	}
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}
