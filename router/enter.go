package router

import (
	"github.com/wxshope/go-admin/router/example"
	"github.com/wxshope/go-admin/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
