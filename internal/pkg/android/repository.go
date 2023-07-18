package android

import (
	"regexp"
	"strings"

	"github.com/prongbang/filex"
	"github.com/prongbang/localizegen/pkg/common"
	"github.com/prongbang/localizegen/pkg/core"
	"github.com/prongbang/localizegen/pkg/csvx"
)

type Repository interface {
	GenerateXmlResources(csv csvx.CsvList, languages core.Languages, localeIndex int) string
	GenerateXmlResourceFile(csv csvx.CsvList, languages core.Languages, locale string, target string, localeIndex int, filename string) core.Result
}

type repository struct {
	FileX filex.FileX
}

func (r *repository) GenerateXmlResourceFile(csv csvx.CsvList, languages core.Languages, locale string, target string, localeIndex int, filename string) core.Result {
	targets := target + "/values"
	if locale != core.LocaleEn && strings.Index(target, "/values") == -1 {
		targets = target + "/values-" + locale
	}
	targets = strings.ReplaceAll(targets, "/values/values", "/values")

	xmlX := r.GenerateXmlResources(csv, languages, localeIndex)
	fName, err := r.FileX.CreateFile(targets, filename, xmlX)
	result := core.Result{Filename: fName, Status: "Success"}
	if err != nil {
		result.Status = "Error"
		result.Message = err.Error()
	}
	return result
}

func (r *repository) GenerateXmlResources(csv csvx.CsvList, languages core.Languages, localeIndex int) string {
	content := "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n"
	content += "<resources>"
	localeLen := len(languages)
	for i := 1; i < len(csv); i++ {
		row := csv[i]

		rowLen := common.ColNotEmpty(row)
		if rowLen >= 0 && rowLen <= localeLen {
			continue
		}

		formatted := ""
		if strings.Index(row[1], "%s") > -1 || strings.Index(row[1], "%d") > -1 || strings.Index(row[1], "%@") > -1 {
			formatted = " formatted=\"false\""
		}
		match, _ := regexp.MatchString("{([0-9])}", row[1])
		if match {
			formatted = " formatted=\"true\""
		}

		// Pattern
		escapedKey, _ := regexp.Compile("\\s{1,}")
		escapedStringBinding, _ := regexp.Compile("%@")
		escapedDoubleQuote, _ := regexp.Compile("[\"]")
		escapedSingleQuote, _ := regexp.Compile("[']")
		escapedDots, _ := regexp.Compile("\\.\\.\\.")
		escapedAmp, _ := regexp.Compile("[&]")
		escapedPercent, _ := regexp.Compile("[%]")
		escapedLessThan, _ := regexp.Compile("[<]")
		escapedLessThanEquals, _ := regexp.Compile("[≤]")
		escapedMoreThan, _ := regexp.Compile("[>]")
		escapedMoreThanEquals, _ := regexp.Compile("[≥]")
		escapedStringAutoBinding, _ := regexp.Compile("{([0-9])}")

		// Prepare data
		escapedContent := row[localeIndex]
		escapedContent = escapedDoubleQuote.ReplaceAllString(escapedContent, `"`)
		escapedContent = escapedSingleQuote.ReplaceAllString(escapedContent, "\\'")
		if !strings.Contains(escapedContent, "&lt;") &&
			!strings.Contains(escapedContent, "&gt;") &&
			!strings.Contains(escapedContent, "&le;") &&
			!strings.Contains(escapedContent, "&ge;") {
			escapedContent = escapedAmp.ReplaceAllString(escapedContent, "&amp;")
		}
		escapedContent = escapedDots.ReplaceAllString(escapedContent, "&#8230;")
		escapedContent = escapedPercent.ReplaceAllString(escapedContent, "%")
		escapedContent = escapedStringBinding.ReplaceAllString(escapedContent, "%s")
		escapedContent = escapedLessThan.ReplaceAllString(escapedContent, "&lt;")
		escapedContent = escapedLessThanEquals.ReplaceAllString(escapedContent, "&le;")
		escapedContent = escapedMoreThan.ReplaceAllString(escapedContent, "&gt;")
		escapedContent = escapedMoreThanEquals.ReplaceAllString(escapedContent, "&ge;")
		escapedContent = escapedStringAutoBinding.ReplaceAllString(escapedContent, "%${1}$$s")

		// Create key name
		key := escapedKey.ReplaceAllString(row[0], "_")
		key = strings.ToLower(key)
		content += "\n\t<string name=\"" + key + "\"" + formatted + ">" + escapedContent + "</string>"
	}
	content += "\n</resources>"
	return content
}

func NewRepository(fileX filex.FileX) Repository {
	return &repository{
		FileX: fileX,
	}
}
