package ios

import (
	"github.com/google/wire"
)

var ProvideSet = wire.NewSet(
	NewRepository,
	NewUseCase,
)
