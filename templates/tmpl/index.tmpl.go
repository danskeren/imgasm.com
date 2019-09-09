package tmpl

const IndexHTML = `
{{ define "content" }}

<div class="content">
  <form class="upload" action="/" method="post" enctype="multipart/form-data" onsubmit="upload.disabled = true; return true;">
    <input type="file" accept="image/*" name="file">
    <div class="submit">
      <button type="submit" name="upload">Upload Image</button>
    </div>
  </form>
</div>

{{ end }}
`