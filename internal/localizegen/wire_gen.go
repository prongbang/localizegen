// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//go:build !wireinject
// +build !wireinject

package localizegen

import (
	"github.com/prongbang/filex"
	"github.com/prongbang/localizegen/internal/pkg/android"
	"github.com/prongbang/localizegen/internal/pkg/flutter"
	"github.com/prongbang/localizegen/internal/pkg/ios"
	"github.com/prongbang/localizegen/pkg/csvx"
)

// Injectors from wire.go:

func New() UseCase {
	callX := NewCallX()
	csvX := csvx.New()
	fileX := filex.New()
	localizegenRepository := NewRepository(callX, csvX, fileX)
	androidRepository := android.NewRepository(fileX)
	useCase := android.NewUseCase(androidRepository)
	iosRepository := ios.NewRepository(fileX)
	iosUseCase := ios.NewUseCase(iosRepository)
	flutterRepository := flutter.NewRepository(fileX)
	flutterUseCase := flutter.NewUseCase(flutterRepository)
	localizegenUseCase := NewUseCase(localizegenRepository, useCase, iosUseCase, flutterUseCase)
	return localizegenUseCase
}
