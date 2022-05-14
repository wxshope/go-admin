package request

import (
	"github.com/wxshope/go-admin/model/{{.Package}}"
	"github.com/wxshope/go-admin/model/common/request"
)

type {{.StructName}}Search struct{
    {{.Package}}.{{.StructName}}
    request.PageInfo
}
