package android

import (
	"github.com/prongbang/localizegen/pkg/core"
	"github.com/prongbang/localizegen/pkg/csvx"
)

type UseCase interface {
	Generate(csv csvx.CsvList, locale string, target string, filename string, languages core.Languages, state chan core.Result, onAdd func(), onDone func())
}

type useCase struct {
	Repo Repository
}

func (u *useCase) Generate(csv csvx.CsvList, locale string, target string, filename string, languages core.Languages, state chan core.Result, onAdd func(), onDone func()) {
	if filename == "" {
		filename = "strings.xml"
	}
	if locale != "" {
		localeIndex := languages[locale].Index
		if localeIndex > 0 {
			onAdd()
			go func() {
				defer onDone()
				state <- u.Repo.GenerateXmlResourceFile(csv, locale, target, localeIndex, filename)
			}()
		} else {
			state <- core.Result{Message: "Language key " + locale + " not found.", Status: "Error"}
		}
	} else {
		for k, v := range languages {
			onAdd()
			go func(localeIndex int, locale string) {
				defer onDone()
				state <- u.Repo.GenerateXmlResourceFile(csv, locale, target, localeIndex, filename)
			}(v.Index, k)
		}
	}
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}
