package globals

import (
	"context"

	dbIface "github.com/taubyte/go-interfaces/services/substrate/database"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

type Factory struct {
	helpers.Methods
	parent       vm.Instance
	databaseNode dbIface.Service
	ctx          context.Context

	databaseInstance dbIface.Database
}

var _ vm.Factory = &Factory{}
