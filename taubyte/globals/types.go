package globals

import (
	"context"

	dbIface "github.com/taubyte/go-interfaces/services/substrate/components/database"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-core-plugins/taubyte/helpers"
)

type Factory struct {
	helpers.Methods
	parent       vm.Instance
	databaseNode dbIface.Service
	ctx          context.Context

	databaseInstance dbIface.Database
}

var _ vm.Factory = &Factory{}
