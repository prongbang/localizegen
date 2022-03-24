package ios

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
	filenameString := configuration.Filename
	if configuration.Filename == "" {
		filenameString = "Localizable.strings"
	}
	filenameSwift := "Localizables.swift"
	if configuration.Locale != "" {
		localeIndex := configuration.Languages[configuration.Locale].Index
		if localeIndex > 0 {
			configuration.OnAdd()
			go func() {
				defer configuration.OnDone()
				configuration.Result <- u.Repo.GenerateSwiftResourceFile(configuration.Csv, configuration.Target, localeIndex, filenameSwift)
				configuration.Result <- u.Repo.GenerateStringsResourceFile(configuration.Csv, configuration.Locale, configuration.Target, localeIndex, filenameString)
			}()
		} else {
			configuration.Result <- core.Result{Message: "Language key " + configuration.Locale + " not found.", Status: "Error"}
		}
	} else {
		first := true
		for k, v := range configuration.Languages {
			configuration.OnAdd()
			go func(localeIndex int, locale string) {
				defer configuration.OnDone()
				if first {
					first = false
					configuration.Result <- u.Repo.GenerateSwiftResourceFile(configuration.Csv, configuration.Target, localeIndex, filenameSwift)
				}
				configuration.Result <- u.Repo.GenerateStringsResourceFile(configuration.Csv, locale, configuration.Target, localeIndex, filenameString)
			}(v.Index, k)
		}
	}
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}
