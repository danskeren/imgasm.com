package tmpl

const ProtectYourPrivacyHTML = `
{{ define "content" }}
<div class="content">
  <h1>Protect Your Privacy</h1>
  <p>
    <a href="https://www.youtube.com/watch?v=3gJvABEi3wQ" target="_blank">"We kill people based on metadata"</a> - Michael Hayden, former NSA and CIA director.
  </p>
  <p>
    In China, innocent people are sent to <a href="https://en.wikipedia.org/wiki/Xinjiang_re-education_camps" target="_blank">Nazi-like concentration camps</a>, and <a href="https://en.wikipedia.org/wiki/Organ_harvesting_from_Falun_Gong_practitioners_in_China" target="_blank">political prisoners are being executed "on demand" to harvest their organs</a>.
  </p>
  <p>
    Governments and companies all over the world violate people's privacy, and abuse the data in some of the most gruesome ways imaginable. Most of this data is being obtained by companies who seek to "serve better ads", and/or companies that carelessly store everything they can, even if they have no need for it. This data is later hacked, sold, and/or handed over to governments.
  </p>
  <p>
    At this time, the European Union is our best bet to win back our right to privacy. I strongly encourage you to vote for political parties that supports EU, so that we can continue to obtain <a href="https://en.wikipedia.org/wiki/General_Data_Protection_Regulation" target="_blank">great laws like the General Data Protection Regulation (GDPR)</a>. Laws that seek to turn user data from a corporate asset into a legal liability. If we can get internet companies to treat user data the same way brick and mortar companies treat hazardous chemicals - occasionally necessary for some industrial process, but best avoided where possible, and always treated with respect - then it would be a big win for society.
  </p>
  <p>
    Besides voting, you can help protect yourself, and others, by making changes the software and extensions you use to browse the internet.
  </p>
  <h2>Upgrade your browser to <a href="https://www.mozilla.org/en-US/firefox/" target="_blank">FireFox</a></h2>
  <p>
    A lot of tech-savvy people made the grave mistake of recommending and installing Chrome everywhere. This has greatly contributed to the fact that Chrome is now, by far, the most popular browser with around 60-70% market share. To make matters worse, pretty much every browser besides FireFox and Safari is built upon Chromium. This is a problem because Google can continiously create new web standards and push them through W3C, or ignore web standards altogether, and thus ensure that certain websites, such as Google Hangouts, Inbox, Allo, YouTube TV, etc. only works with Chromium-based browsers. 
  </p>
  <p>
    This has finally come back to visibly bite people in the ass, as Google is making modifications to Chrome (and Chromium) in order to cripple ad blockers. Unfortunately, Google controls such a large portion of the web, and has many ways to make the browsing experience extremely unpleasant for those who seek to protect their privacy - primarily thanks to reCAPTCHA. Nonetheless, we shouldn't give up our right to privacy due to this inconvenience. We should instead encourage websites not to use reCAPTCHA, encourage people to build websites that work without JavaScript (such as this one), and encourage everyone to switch to FireFox and install extensions that help protect their privacy.
  </p>
  <h3>Install <a href="https://addons.mozilla.org/en-US/firefox/addon/adnauseam/" target="_blank">AdNauseam</a> or <a href="https://addons.mozilla.org/en-US/firefox/addon/ublock-origin/" target="_blank">uBlock Origin</a> Extension</h3>
  <p>
    AdNauseam is an extension that helps fight back against advertising surveillance. It's built on top of uBlock Origin - this means you still get the malware blocking and ad hiding capabilities of uBlock Origin. But AdNauseam goes a step further, it automatically clicks on ads in the background, polluting your data profile and frustrating the trackers who violate your privacy and facilitate bulk-surveillance agendas. It is, as far as I know, the only non-malware extension that Google has banned from the Chrome Web Store.
  </p>
  <p>
    Please do <strong>not</strong> use or recommend AdBlock, AdBlock Plus or uBlock. They allow companies that violate people's privacy to pay them to be whitelisted.
  </p>
  <h3>Install <a href="https://addons.mozilla.org/en-US/firefox/addon/cookie-autodelete/" target="_blank">Cookie AutoDelete</a> Extension</h3>
  <p>
    Cookie AutoDelete automatically deletes cookies that hasn't been whitelisted.
  </p>
  <h3>Install <a href="https://addons.mozilla.org/en-US/firefox/addon/privacy-badger17/" target="_blank">Privacy Badger</a> Extension</h3>
  <p>
    Privacy Badger helps stop advertisers and other third-party trackers from secretly tracking where you go and what pages you look at on the web.
  </p>
  <h3>Install <a href="https://addons.mozilla.org/en-US/firefox/addon/decentraleyes/" target="_blank">Decentraleyes</a> Extension</h3>
  <p>
    Decentraleyes protects against tracking by "free", centralized, content delivery networks (CDNs). The extension works by bundling files, such as jQuery and AngularJS, locally.
  </p>
  <p>
    When you combine AdNauseam with Cookie AutoDelete, Privacy Badger and Decentraleyes, as well as FireFox's built-in Strict Content Blocking, then you're causing annoyance to those violating your privacy.
  </p>
  <h3>Install <a href="https://addons.mozilla.org/en-US/firefox/addon/noscript/" target="_blank">NoScript</a> or <a href="https://addons.mozilla.org/en-US/firefox/addon/umatrix/" target="_blank">uMatrix</a> Extension (for tech-savvy users)</h3>
  <p>
    No matter what extensions you install, then you won't be able to prevent all trackers. Using FireFox, along with AdNauseam, Cookie AutoDelete, Privacy Badger and Decentraleyes will go a long way towards fighting back against those who seek to violate people's privacy, and you'll be able to achieve this without hurting your browsing experience.
  </p>
  <p>
    If you wish to go a step further, then you can also install the NoScript or uMatrix extension, which means that you have to manually allow JavaScript on every website you visit. This will visibly affect your browsing experience, but it will also go a long way towards protecting your privacy.
  </p>
  <h2>Search with <a href="https://duckduckgo.com" target="_blank">DuckDuckGo</a></h2>
  <p>
    Remember to set it to your default search engine.
  </p>
  <h2>Share and backup personal data with <a href="https://syncthing.net/" target="_blank">Syncthing</a></h2>
  <p>
    Syncthing is an open source alternative to proprietary sync and cloud services, such as Dropbox and Google Drive. The incredible thing about Syncthing is that it doesn't upload any of your files to the cloud. It instead sync files between the devices you choose, e.g. your desktop, laptop and phone, and thus make it possible for you to easily backup and transfer files between your devices.
  </p>
</div>
{{ end }}
`
