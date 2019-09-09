package tmpl

const NotFoundHTML = `
{{ define "head" }}
<meta name="robots" content="noindex">
{{ end }}

{{ define "content" }}
<div class="content">
  <h1>404! The page you're looking for does not exist.</h1>
  <p>
    Please double-check that you typed the URL correctly. If you're certain it's the correct URL, then the image has most likely been deleted.
  </p>
  <p>
    Return to homepage to <a href="/">upload a new image</a>.
  </p>
</div>
{{ end }}
`
