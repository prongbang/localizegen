package core

import (
	"github.com/prongbang/localizegen/pkg/csvx"
)

type Configuration struct {
	Csv       csvx.CsvList
	Locale    string
	Target    string
	Filename  string
	Languages Languages
	Result    chan Result
	OnAdd     func()
	OnDone    func()
}
