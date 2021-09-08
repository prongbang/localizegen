//+build wireinject

package localizegen

import (
	"github.com/google/wire"
	"github.com/prongbang/filex"
	"github.com/prongbang/localizegen/internal/pkg/android"
	"github.com/prongbang/localizegen/internal/pkg/ios"
	"github.com/prongbang/localizegen/pkg/csvx"
)

func New() UseCase {
	wire.Build(
		NewRepository,
		NewUseCase,
		NewCallX,
		csvx.New,
		filex.New,
		android.ProvideSet,
		ios.ProvideSet,
	)
	return nil
}
