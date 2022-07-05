//go:build wireinject
// +build wireinject

package user

import "github.com/google/wire"

func initService() Service {
	wire.Build(newService, newDao, newCache)
	return nil
}
