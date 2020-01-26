{{ define "packages" }}
---
title: S3::Bucket
description: From zero to docs deployed on Netlify in just a few minutes
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
