package example

import "github.com/wxshope/go-admin/model/system"

type ExcelInfo struct {
	FileName string               `json:"fileName"` // 文件名
	InfoList []system.SysBaseMenu `json:"infoList"`
}
