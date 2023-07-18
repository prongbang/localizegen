package ios

import (
	"regexp"
	"strings"

	"github.com/prongbang/filex"
	"github.com/prongbang/localizegen/pkg/common"
	"github.com/prongbang/localizegen/pkg/core"
	"github.com/prongbang/localizegen/pkg/csvx"
)

type Repository interface {
	GenerateStringsResources(csv csvx.CsvList, languages core.Languages, localeIndex int) string
	GenerateSwiftResources(csv csvx.CsvList, languages core.Languages, localeIndex int) string
	GenerateStringsResourceFile(csv csvx.CsvList, languages core.Languages, locale string, target string, localeIndex int, filename string) core.Result
	GenerateSwiftResourceFile(csv csvx.CsvList, languages core.Languages, target string, localeIndex int, filename string) core.Result
}

type repository struct {
	FileX filex.FileX
}

func (r *repository) GenerateStringsResources(csv csvx.CsvList, languages core.Languages, localeIndex int) string {
	content := ""
	localeLen := len(languages)
	for i := 1; i < len(csv); i++ {
		row := csv[i]

		rowLen := common.ColNotEmpty(row)
		if rowLen >= 0 && rowLen <= localeLen {
			continue
		}

		// Pattern
		escapedKey, _ := regexp.Compile("\\s{1,}")
		escapedStringBinding, _ := regexp.Compile("%s")
		escapedDigitBinding, _ := regexp.Compile("{\\d}")
		escapedDoubleQuote, _ := regexp.Compile("[\"]")
		escapedSingleQuote, _ := regexp.Compile("[']")
		escapedNewLine, _ := regexp.Compile("(?:\\r\\n|\\r|\\n)")

		// Prepare data
		escapedContent := row[localeIndex]
		escapedContent = escapedStringBinding.ReplaceAllString(escapedContent, "%@")
		escapedContent = escapedDigitBinding.ReplaceAllString(escapedContent, "%@")
		escapedContent = escapedDoubleQuote.ReplaceAllString(escapedContent, `"`)
		escapedContent = escapedSingleQuote.ReplaceAllString(escapedContent, "\\'")
		escapedContent = escapedNewLine.ReplaceAllString(escapedContent, "\\n")

		// Create key
		key := escapedKey.ReplaceAllString(row[0], "_")
		key = strings.ToLower(key)
		content += "\n\"" + key + "\" = \"" + escapedContent + "\";"
	}
	return content
}

func (r *repository) GenerateSwiftResources(csv csvx.CsvList, languages core.Languages, localeIndex int) string {
	content := "import Foundation\n\npublic enum LocalizablesType: String {"
	localeLen := len(languages)
	for i := 1; i < len(csv); i++ {
		row := csv[i]

		rowLen := common.ColNotEmpty(row)
		if rowLen >= 0 && rowLen <= localeLen {
			continue
		}

		// Pattern
		escapedKey, _ := regexp.Compile("\\s{1,}")

		// Prepare data
		key := escapedKey.ReplaceAllString(row[0], "_")
		key = strings.ToLower(key)

		content += "\n\tcase " + key
	}
	content += "\n}"

	return content
}

func (r *repository) GenerateStringsResourceFile(csv csvx.CsvList, languages core.Languages, locale string, target string, localeIndex int, filename string) core.Result {
	targets := target + "/" + locale + ".lproj"
	stringS := r.GenerateStringsResources(csv, languages, localeIndex)
	fName, err := r.FileX.CreateFile(targets, filename, stringS)
	result := core.Result{Filename: fName, Status: "Success"}
	if err != nil {
		result.Status = "Error"
		result.Message = err.Error()
	}
	return result
}

func (r *repository) GenerateSwiftResourceFile(csv csvx.CsvList, languages core.Languages, target string, localeIndex int, filename string) core.Result {
	stringS := r.GenerateSwiftResources(csv, languages, localeIndex)
	fName, err := r.FileX.CreateFile(target, filename, stringS)
	result := core.Result{Filename: fName, Status: "Success"}
	if err != nil {
		result.Status = "Error"
		result.Message = err.Error()
	}
	return result
}

func NewRepository(fileX filex.FileX) Repository {
	return &repository{
		FileX: fileX,
	}
}
