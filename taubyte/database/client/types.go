package client

import (
	"context"
	"sync"

	dbIface "github.com/taubyte/go-interfaces/services/substrate/database"
	"github.com/taubyte/go-interfaces/vm"
	"github.com/taubyte/vm-plugins/taubyte/helpers"
)

type Factory struct {
	helpers.Methods
	databaseNode     dbIface.Service
	parent           vm.Instance
	CurrentKeystore  string
	ctx              context.Context
	databaseLock     sync.RWMutex
	databaseIdToGrab uint32
	database         map[uint32]*Database
}

var _ vm.Factory = &Factory{}

type Database struct {
	dbIface.Database
	Id uint32
}
