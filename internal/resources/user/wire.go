//go:build wireinject
// +build wireinject

package user

import "github.com/google/wire"

func initController() *controller {
	wire.Build(newController, newService, newDao, newCache)
	return nil
}
