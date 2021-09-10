package ios

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
	filenameString := filename
	if filename == "" {
		filenameString = "Localizable.strings"
	}
	filenameSwift := "Localizables.swift"
	if locale != "" {
		localeIndex := languages[locale].Index
		if localeIndex > 0 {
			onAdd()
			go func() {
				defer onDone()
				state <- u.Repo.GenerateSwiftResourceFile(csv, target, localeIndex, filenameSwift)
				state <- u.Repo.GenerateStringsResourceFile(csv, locale, target, localeIndex, filenameString)
			}()
		} else {
			state <- core.Result{Message: "Language key " + locale + " not found.", Status: "Error"}
		}
	} else {
		first := true
		for k, v := range languages {
			onAdd()
			go func(localeIndex int, locale string) {
				defer onDone()
				if first {
					first = false
					state <- u.Repo.GenerateSwiftResourceFile(csv, target, localeIndex, filenameSwift)
				}
				state <- u.Repo.GenerateStringsResourceFile(csv, locale, target, localeIndex, filenameString)
			}(v.Index, k)
		}
	}
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}
