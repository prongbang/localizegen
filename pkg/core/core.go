package core

type Result struct {
	Filename string
	Status   string
	Message  string
}

type Language struct {
	Name  string
	Key   string
	Index int
}

type Languages map[string]Language

const LocaleEn = "en"
const LocaleTh = "th"
