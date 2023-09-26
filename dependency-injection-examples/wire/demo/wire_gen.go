// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"demo/business"
	"demo/database"
	"demo/service"
)

// Injectors from wire.go:

func InitializeService() service.Service {
	databaseDatabase := database.NewDatabase()
	businessBusiness := business.NewBusiness(databaseDatabase)
	serviceImpl := service.NewService(businessBusiness)
	return serviceImpl
}
