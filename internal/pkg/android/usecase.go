package android

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
	if configuration.Filename == "" {
		configuration.Filename = "strings.xml"
	}
	if configuration.Locale != "" {
		localeIndex := configuration.Languages[configuration.Locale].Index
		if localeIndex > 0 {
			configuration.OnAdd()
			go func() {
				defer configuration.OnDone()
				configuration.Result <- u.Repo.GenerateXmlResourceFile(configuration.Csv, configuration.Languages, configuration.Locale, configuration.Target, localeIndex, configuration.Filename)
			}()
		} else {
			configuration.Result <- core.Result{Message: "Language key " + configuration.Locale + " not found.", Status: "Error"}
		}
	} else {
		for k, v := range configuration.Languages {
			configuration.OnAdd()
			go func(localeIndex int, locale string) {
				defer configuration.OnDone()
				configuration.Result <- u.Repo.GenerateXmlResourceFile(configuration.Csv, configuration.Languages, locale, configuration.Target, localeIndex, configuration.Filename)
			}(v.Index, k)
		}
	}
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}
