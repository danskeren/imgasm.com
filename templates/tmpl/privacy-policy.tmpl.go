package tmpl

const PrivacyPolicyHTML = `
{{ define "content" }}
<div class="content">
  <h1>Privacy Policy</h1>
  <p>
    I care deeply about user privacy, which is why I don't store any user-related information. The website simply have two key-value buckets to store information related to the uploaded files. The file data stored is as follows:
  </p>
  <h3>'MD5 hash key-value bucket'</h3>
  <ul>
    <li>Key: MD5 hash of the uploaded file</li>
    <li>Value: MD5 hash of the file after being processed (if it wasn't converted/resized then the value will be the same as the key)</li>
  </ul>
  <p>
    The value is basically the MD5 hash of the file that ends up being stored on the server. The purpose of this bucket is to prevent duplicates and processing files that has already been processed.
  </p>
  <h3>'File data key-value bucket'</h3>
  <ul>
    <li>Key: A unique ID pointing to the file (this is the short ID you see in the URL)</li>
    <li>Value: Filename of the stored file (this is the MD5 hash + file extension), along with a randomly generated password that can be used to delete the data from the bucket</li>
  </ul>
  <p>
    It's worth noting that the file itself won't be deleted from the server, but no one will be able to view the file unless they know the MD5 hash. I can manually delete the file from the server, but will only do so if you can prove you are the copyright owner.
  </p>
  <p>
    That's it. There is no futher information being stored. The uploaded images are not associated with an IP or anything else. Moreover, the website doesn't use any analytics or tracking. I do however use Cloudflare, which means I can see number of unique visitors and requests for 1 month. This information also include the country from which the request was made, but contain no other information.
  </p>
  <p>
    In order to protect imgasm against abuse and attacks, I use a <a href="https://github.com/ulule/limiter" target="_blank">rate limiter package</a>, that will store your IP in memory for 1 hour. I also use CloudFlare, which will <a href="https://blog.cloudflare.com/what-cloudflare-logs/" target="_blank">discard access logs within 4 hours</a>. I do <strong>not</strong> log anything myself, as I feel confident that the rate limiter together with CloudFlare can mitigate any attacks that may occur.
  </p>
  <p>
    imgasm also use AES-256 encrypted cookies for the sole purpose of displaying potential error messages. You can easily delete and block the cookies without breaking anything (besides the fact that potential error messages won't be displayed if cookies are being blocked).
  </p>
  <p>
    If you encounter any problems, or have any questions, then feel free to <a href="https://github.com/danskeren/imgasm.com/issues/new" target="_blank">open an issue on GitHub</a>.
  </p>
</div>
{{ end }}
`
