package global

{{- if .HasGlobal }}

import "github.com/wxshope/go-admin/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}