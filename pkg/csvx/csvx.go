package csvx

import (
	"encoding/csv"
	"strings"
)

type CsvList [][]string

type CsvX interface {
	ReadAll(text string) CsvList
}

type mtd struct {
}

func (c *mtd) ReadAll(text string) CsvList {
	r := csv.NewReader(strings.NewReader(text))
	records, err := r.ReadAll()
	if err != nil {
		return CsvList{}
	}
	return records
}

func New() CsvX {
	return &mtd{}
}
