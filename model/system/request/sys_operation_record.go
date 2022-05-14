package request

import (
	"github.com/wxshope/go-admin/model/common/request"
	"github.com/wxshope/go-admin/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
