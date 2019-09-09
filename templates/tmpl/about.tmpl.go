package tmpl

const AboutHTML = `
{{ define "content" }}
<div class="content">
  <h1>About</h1>
  <p>
    imgasm (pronounced image-gasm) is a free, open source image sharing service that respects people's privacy.
  </p>
  <p>
    I built imgasm because I got fed up with the incredible poor user experience and evil behaviour of imgur. To name a few of my many complaints:
  </p>
  <ul>
    <li>Violate users privacy. To add insult to injury they even have a popup saying "we value your privacy", which uses dark patterns to make it harder, if not impossible, to prevent them from selling your data to their hundred plus different partners</li>
    <li>Loading tiny resolution images while making it impossible to request the original image</li>
    <li>Redirection hell. Direct links are often redirected to pages filled with user tracking, ads and popups. The same redirection hell makes it impossible to load a higher resolution of the image when you for some reason are served with a 50px resolution image</li>
  </ul>
  <p>
    imgasm will go in the opposite direction. I believe digital privacy and anonymity is a human right. imgasm does not and will not violate people's privacy. To get a better understanding of how imgasm works and how little information is stored, check out the <a href="/privacy-policy">Privacy Policy</a> page.
  </p>
  <p>
    Profitability is something a lot of people worry and are curious about when it comes to file hosting/sharing websites. Which is understandable considering all the good websites tend to shut down or become evil once the hosting costs become too high. 
  </p>
  <p>
    I want to avoid being forced to shut down, which is why the website will initially focus exclusively on images with a relatively low size limit, and with no durability promises. I believe that good code, along with smart hosting, storage and cacheing will allow me to handle millions of users with ease, while serving images at a speed that's faster than the competitors regardless of your location.
  </p>
  <p>
    I hope to cover the hosting costs, and perhaps even be profitable, through <a href="https://github.com/danskeren">GitHub Donations</a>, as well as hardcoded promotions that will appear in the header and footer. These promotions will not involve any user tracking.
  </p>
  <p>
    If you wish to support me, then I would appreciate if you share the website, <a href="https://github.com/danskeren/imgasm.com">star imgasm on GitHub</a> (as well as my other projects), and donate.
  </p>
  <p>
    With enough support then I will increase the size limit, and add support for video files, since Reddit Videos and Gfycat are also evil.
  </p>
</div>
{{ end }}
`
