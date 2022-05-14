package service

import (
	"github.com/wxshope/go-admin/service/example"
	"github.com/wxshope/go-admin/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
