package self

import (
	"context"

	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-core-plugins/taubyte/helpers"
)

type Factory struct {
	helpers.Methods
	parent vm.Instance
	ctx    context.Context
}

var _ vm.Factory = &Factory{}
