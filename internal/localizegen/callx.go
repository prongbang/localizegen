package localizegen

import "github.com/prongbang/callx"

func NewCallX() callx.CallX {
	c := callx.Config{
		BaseURL: "https://docs.google.com/spreadsheets/d",
		Timeout: 60,
	}
	return callx.New(c)
}
