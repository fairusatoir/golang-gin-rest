//go:build wireinject
// +build wireinject

package app

import (
	"github.com/fairusatoir/golang-gin-rest/cmd/api/config"
	"github.com/fairusatoir/golang-gin-rest/internal/clients"
	"github.com/fairusatoir/golang-gin-rest/internal/controllers/handlers"
	"github.com/fairusatoir/golang-gin-rest/internal/repositories"
	"github.com/fairusatoir/golang-gin-rest/internal/services"
	"github.com/google/wire"
)

func NewServerAPI() (*Server, error) {
	wire.Build(
		config.Load,
		clients.Connect,
		repositories.NewProductRepo,
		services.NewProductService,
		handlers.NewProductHandler,
		NewServer,
	)

	return nil, nil
}
