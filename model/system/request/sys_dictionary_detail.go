package request

import (
	"github.com/wxshope/go-admin/model/common/request"
	"github.com/wxshope/go-admin/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
