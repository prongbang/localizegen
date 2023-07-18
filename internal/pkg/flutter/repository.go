package flutter

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/prongbang/filex"
	"github.com/prongbang/localizegen/pkg/common"
	"github.com/prongbang/localizegen/pkg/core"
	"github.com/prongbang/localizegen/pkg/csvx"
)

type Repository interface {
	GenerateKeys(csv csvx.CsvList, languages core.Languages) string
	GenerateSources(csv csvx.CsvList, languages core.Languages) string
	GenerateKeysFile(csv csvx.CsvList, languages core.Languages, target string, filename string) core.Result
	GenerateSourcesFile(csv csvx.CsvList, languages core.Languages, target string, filename string) core.Result
}

type repository struct {
	FileX filex.FileX
}

func (r *repository) GenerateKeys(csv csvx.CsvList, languages core.Languages) string {
	content := "import 'package:localization/localization.dart';\n\n"
	content += "class KeysLocalizations extends TranslateLocalizations {"
	localeLen := len(languages)
	for i := 1; i < len(csv); i++ {
		row := csv[i]

		rowLen := common.ColNotEmpty(row)
		if rowLen >= 0 && rowLen <= localeLen {
			continue
		}

		// Pattern
		escapedKey, _ := regexp.Compile("\\s{1,}")
		escapedAdBinding, _ := regexp.Compile("%@")
		escapedDigitBinding, _ := regexp.Compile("{\\d}")

		// Prepare key
		key := escapedKey.ReplaceAllString(row[0], "_")
		key = strings.ToLower(key)

		// Prepare value
		value := row[1]
		value = escapedAdBinding.ReplaceAllString(value, "%s")
		value = escapedDigitBinding.ReplaceAllString(value, "%s")

		count := strings.Count(value, "%s")
		if count > 0 {
			parameters := ""
			comma := ", "
			arguments := "["
			for c := 1; c <= count; c++ {
				if c == count {
					comma = ""
				}
				args := fmt.Sprintf("arg%d", c)
				parameters += fmt.Sprintf("String %s%s", args, comma)
				arguments += fmt.Sprintf("%s%s", args, comma)
			}
			arguments += "]"
			parameters += ""
			content += "\n\tString " + common.ToCamelVariable(key) + "(" + parameters + ") => sprintf(translate('" + key + "'), " + arguments + ");"
		} else {
			content += "\n\tString get " + common.ToCamelVariable(key) + " => translate('" + key + "');"
		}
	}
	content += "\n}"

	return content
}

func (r *repository) GenerateSources(csv csvx.CsvList, languages core.Languages) string {
	content := "import 'dart:ui';\n\n"
	content += "class SourcesLocalizations {\n"
	content += "\tMap<String, String> loadLocale(Locale locale) => _supported[locale.languageCode] ?? {};\n\n"

	// Sort locale keys
	localeKeys := make([]string, 0, len(languages))
	for k := range languages {
		localeKeys = append(localeKeys, k)
	}
	sort.Strings(localeKeys)

	// Generate supported locale
	supported := "\tMap<String, Map<String, String>> get _supported => {\n"
	sources := map[string]string{}
	localeLen := len(languages)
	for _, key := range localeKeys {
		locale := languages[key]
		sources[locale.Key] += "\tMap<String, String> get _" + locale.Key + " => {\n"
		supported += "\t\t\t'" + locale.Key + "'" + ": _" + locale.Key + ",\n"
	}
	supported += "\t\t\t};\n\n"

	for i := 1; i < len(csv); i++ {
		row := csv[i]

		rowLen := common.ColNotEmpty(row)
		if rowLen >= 0 && rowLen <= localeLen {
			continue
		}

		// Pattern
		escapedKey, _ := regexp.Compile("\\s{1,}")
		escapedAdBinding, _ := regexp.Compile("%@")
		escapedDigitBinding, _ := regexp.Compile("{\\d}")
		escapedDoubleQuote, _ := regexp.Compile("[\"]")
		escapedSingleQuote, _ := regexp.Compile("[']")
		escapedNewLine, _ := regexp.Compile("(?:\\r\\n|\\r|\\n)")

		// Prepare data
		for _, k := range localeKeys {
			locale := languages[k]
			escapedContent := row[locale.Index]
			escapedContent = escapedAdBinding.ReplaceAllString(escapedContent, "%s")
			escapedContent = escapedDigitBinding.ReplaceAllString(escapedContent, "%s")
			escapedContent = escapedDoubleQuote.ReplaceAllString(escapedContent, `"`)
			escapedContent = escapedSingleQuote.ReplaceAllString(escapedContent, "\\'")
			escapedContent = escapedNewLine.ReplaceAllString(escapedContent, "\\n")

			// Create key
			key := escapedKey.ReplaceAllString(row[0], "_")
			key = strings.ToLower(key)

			sources[locale.Key] += "\t\t'" + key + "': '" + escapedContent + "',\n"

			// last row
			if i == len(csv)-1 {
				sources[locale.Key] += "\t};\n\n"
			}
		}
	}

	// Append supported
	content += supported

	// Append language
	for _, k := range localeKeys {
		locale := sources[k]
		content += locale
	}

	content += "}"
	return content
}

func (r *repository) GenerateKeysFile(csv csvx.CsvList, languages core.Languages, target string, filename string) core.Result {
	stringS := r.GenerateKeys(csv, languages)
	fName, err := r.FileX.CreateFile(target, filename, stringS)
	result := core.Result{Filename: fName, Status: "Success"}
	if err != nil {
		result.Status = "Error"
		result.Message = err.Error()
	}
	return result
}

func (r *repository) GenerateSourcesFile(csv csvx.CsvList, languages core.Languages, target string, filename string) core.Result {
	stringS := r.GenerateSources(csv, languages)
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
