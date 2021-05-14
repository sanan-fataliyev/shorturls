# URL Shortener
(Live coding task from a tech interview)

Create a URL shortening service, similar to http://bit.ly/ or other.

A user visiting your shortening website, which you can just run at `localhost:4567`, is prompted to enter the URL they wish to shorten. They are returned their new shortened URL which looks something like `http://localhost:4567/a1b2c3`. When visiting that URL they will be redirected to the original URL they submitted, or render a 404 if the URL does not exist.

The generated URL fragment should be short and unique.