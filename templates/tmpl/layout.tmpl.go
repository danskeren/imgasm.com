package tmpl

const LayoutHTML = `
<!DOCTYPE html>
<html>
<head>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{ if .Common.MetaTitle }}{{ .Common.MetaTitle }} - imgasm.com{{ else }}imgasm.com - when image sharing reaches its climax{{ end }}</title>
  {{ if .Common.MetaDescription }}
		<meta name="description" content="{{ .Common.MetaDescription }}">
	{{ else }}
		<meta name="description" content="Upload and share anonymous images.">
	{{ end }}

  <link rel="apple-touch-icon" sizes="180x180" href="/static/images/icons/apple-touch-icon.png">
  <link rel="icon" type="image/png" sizes="32x32" href="/static/images/icons/favicon-32x32.png">
  <link rel="icon" type="image/png" sizes="16x16" href="/static/images/icons/favicon-16x16.png">
  <link rel="manifest" href="/static/images/icons/site.webmanifest">
  <link rel="mask-icon" href="/static/images/icons/safari-pinned-tab.svg" color="#5bbad5">
  <link rel="shortcut icon" href="/static/images/icons/favicon.ico">
  <meta name="msapplication-TileColor" content="#9f00a7">
  <meta name="msapplication-config" content="/static/images/icons/browserconfig.xml">
  <meta name="theme-color" content="#ffffff">  
  
	<link href="/static/main.min.css" rel="stylesheet" rel="stylesheet">
	
  {{ block "head" . }}{{ end }}
</head>
<body>
  <div class="sticky-footer">
    <header>
      <div class="wrapper">
        <div class="promo">Liberate Hong Kong</div>
        <nav>
          <a href="/">
            <img src="/static/images/logo.png">imgasm.com
          </a>
        </nav>
      </div>
    </header>
 
    <main>
      <div class="content flashes">
        {{ range .Common.Flashes }}
        <p>{{ . }}</p>
        {{ end }}
      </div>
      {{ block "content" . }}{{ end }}
    </main>

    <div class="wrapper">
      <footer>
        <div class="links">
          <div>
            No tracking. No bloat. <a href="https://github.com/danskeren/imgasm.com" target="_blank">Open Source</a>.
          </div>
          <div>
            <a href="/protect-your-privacy">Protect Your Privacy</a>
          </div>
          <div>
            <a href="/privacy-policy">Privacy Policy</a>
          </div>
        </div>
        <div class="copyright">
          <div>
            <svg class="code" fill="currentColor" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M.7 9.3l4.8-4.8 1.4 1.42L2.84 10l4.07 4.07-1.41 1.42L0 10l.7-.7zm18.6 1.4l.7-.7-5.49-5.49-1.4 1.42L17.16 10l-4.07 4.07 1.41 1.42 4.78-4.78z"/></svg>
            with
            <svg class="heart" fill="currentColor" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M10 3.22l-.61-.6a5.5 5.5 0 0 0-7.78 7.77L10 18.78l8.39-8.4a5.5 5.5 0 0 0-7.78-7.77l-.61.61z"/></svg>
            in Copenhagen.
          </div>
          <div>© imgasm, 2013-2019. All rights reserved.</div>
        </div>
      </footer>
      </div>
  </div>
</body> 
</html>
` 
