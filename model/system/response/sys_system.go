package response

import "github.com/wxshope/go-admin/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
