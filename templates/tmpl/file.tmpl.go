package tmpl

const FileHTML = `
{{ define "content" }}

<img src="https://file.imgasm.com/{{ .Filename }}" class="file">

{{ end }}
`