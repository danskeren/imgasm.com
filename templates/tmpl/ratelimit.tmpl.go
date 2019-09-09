package tmpl

const RateLimitHTML = `
{{ define "head" }}
<meta name="robots" content="noindex">
{{ end }}

{{ define "content" }}
<div class="content">
  <h1>Rate Limit!</h1>
  <p>
    To protect the server against abuse, you're only permitted to view {{ .ViewLimit }} pages every hour, and upload {{ .UploadLimit }} files every day. Your understanding is much appreciated.
  </p>
</div>
{{ end }}
`
