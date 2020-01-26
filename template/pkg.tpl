{{ define "packages" }}
---
title: {{ .Name }}
---


{{ range .packages }}
    {{ range (visibleTypes (sortedTypes .Types))}}
        {{ template "type" .  }}
    {{ end }}
    ---
{{ end }}

{{ if .sensitiveData}}
## Sensitive Values
    | Name | Type | Description |
    |------|------|-------------|
    {{ range .sensitiveData }}
        | `{{- .Name -}}` | ***{{- .Type -}}*** | {{- .Description -}} |
    {{- end -}}
{{ end }}

{{ end }}
