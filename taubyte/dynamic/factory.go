package dynamic

import (
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

func New(i vm.Instance, helper helpers.Methods) *Factory {
	return &Factory{
		instance: i,
		ctx:      i.Context().Context(),
		Methods:  helper,
	}
}
