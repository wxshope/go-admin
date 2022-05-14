package v1

import (
	"github.com/wxshope/go-admin/api/v1/example"
	"github.com/wxshope/go-admin/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
