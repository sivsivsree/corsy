
# Corsy 
### CORS injector proxy for development
[![Build](https://github.com/sivsivsree/corsy/actions/workflows/unit-test.yaml/badge.svg)](https://github.com/sivsivsree/corsy/actions/workflows/unit-test.yaml) [![Go Report Card](https://goreportcard.com/badge/github.com/sivsivsree/corsy)](https://goreportcard.com/report/github.com/sivsivsree/corsy)

Corsy is a proxy injector for development to avoid CORS issues while building SPA applications.

[![asciicast](https://asciinema.org/a/425776.svg)](https://asciinema.org/a/425776)

<h4>Install</h4>

1.  Using Homebrew tap, [Install homebrew](https://brew.sh/)
```sh
 brew tap sivsivsree/corsy 
 brew install sivsivsree/corsy/corsy
```

If windows Download the latest exe artifact from the [Release Page](https://github.com/sivsivsree/corsy/releases).



<b>Usage: </b>

```
Options:
  -p, --proxy string        remote address to proxy with cors [REQUIRED]
  -a, --addr string         address:port to listen on :8080  (default ":8001")
  -b, --blacklist strings   Headers to remove from the request and response
  -h, --help                Show this message
  -r, --max-redirects int   Maximum number of redirects to follow (default 10)
  -t, --timeout int         Request timeout (default 15)
```

<b>Example</b>
 <br>

```
corsy --proxy https://prod.example.com/api/v1/resource

```
or 
```
corsy --addr ":8080" --proxy https://prod.example.com/api/v1/resource  \
--timeout 20 --max-redirects 2 \
--blacklist X-Remote-Url 

```
<br>
<br>

> Do contribute and <a href="https://github.com/sivsivsree/corsy/issues/new?assignees=&labels=&template=feature_request.md&title=">create issue</a> if you feel something needs to be changed. <br>
> <h7>Made out of need, will help some of you - Siv 🧑🏻‍💻 </h7>
