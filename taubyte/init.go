package taubyte

import (
	"context"
	"errors"

	"github.com/taubyte/go-interfaces/services/substrate/database"
	"github.com/taubyte/go-interfaces/services/substrate/ipfs"
	"github.com/taubyte/go-interfaces/services/substrate/p2p"
	"github.com/taubyte/go-interfaces/services/substrate/pubsub"
	"github.com/taubyte/go-interfaces/services/substrate/storage"
	"github.com/taubyte/go-interfaces/vm"
)

type plugin struct {
	ctx          context.Context
	ctxC         context.CancelFunc
	ipfsNode     ipfs.Service
	pubsubNode   pubsub.Service
	databaseNode database.Service
	storageNode  storage.Service
	p2pNode      p2p.Service
}

var _plugin *plugin

type Option func() error

func IpfsNode(node ipfs.Service) Option {
	return func() error {
		if _plugin == nil {
			return errors.New("ipfsNode option failed, plugin is nill")
		}
		if node == nil {
			return errors.New("ipfsNode option failed, node is nill")
		}

		_plugin.ipfsNode = node
		return nil
	}
}

func PubsubNode(node pubsub.Service) Option {
	return func() error {
		if _plugin == nil {
			return errors.New("pubsubNode option failed, plugin is nill")
		}
		if node == nil {
			return errors.New(" pubsubNode option failed, node is nill")
		}

		_plugin.pubsubNode = node
		return nil
	}
}

func DatabaseNode(node database.Service) Option {
	return func() error {
		if _plugin == nil {
			return errors.New("databaseNode option failed, plugin is nill")
		}

		if node == nil {
			return errors.New("databaseNode option failed, node is nill")
		}

		_plugin.databaseNode = node
		return nil
	}
}

func StorageNode(node storage.Service) Option {
	return func() error {
		if _plugin == nil {
			return errors.New("storageNode option failed, plugin is nill")
		}
		if node == nil {
			return errors.New("storageNode option failed, node is nill")
		}

		_plugin.storageNode = node
		return nil
	}
}

func P2PNode(node p2p.Service) Option {
	return func() error {
		if _plugin == nil {
			return errors.New("p2pNode option failed, plugin is nill")
		}
		if node == nil {
			return errors.New("p2pNode option failed, node is nill")
		}

		_plugin.p2pNode = node
		return nil
	}
}

func (p *plugin) Name() string {
	return "taubyte/sdk"
}

func Plugin() vm.Plugin {
	return _plugin
}

// First initialize the plugin
func Initialize(ctx context.Context, options ...Option) error {
	if _plugin != nil {
		return nil
	}

	_plugin = &plugin{}

	_plugin.ctx, _plugin.ctxC = context.WithCancel(ctx)

	for _, opt := range options {
		if err := opt(); err != nil {
			return err
		}
	}

	go func() {
		<-_plugin.ctx.Done()
		_plugin.ctxC()
		_plugin = nil
	}()

	return nil
}
