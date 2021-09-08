package localizegen

import (
	"fmt"
	"github.com/prongbang/filex"

	"github.com/prongbang/callx"
	"github.com/prongbang/localizegen/pkg/csvx"
)

type Repository interface {
	GetLocalize(documentID string, sheetID string) callx.Response
	ReadCSV(text string) csvx.CsvList
}

type repository struct {
	CallX callx.CallX
	CsvX  csvx.CsvX
}

func (r *repository) GetLocalize(documentID string, sheetID string) callx.Response {
	return r.CallX.Get(fmt.Sprintf("/%s/export?format=csv&id=%s&gid=%s", documentID, documentID, sheetID))
}

func (r *repository) ReadCSV(text string) csvx.CsvList {
	return r.CsvX.ReadAll(text)
}

func NewRepository(callX callx.CallX, csvX csvx.CsvX, fileX filex.FileX) Repository {
	return &repository{
		CallX: callX,
		CsvX:  csvX,
	}
}
