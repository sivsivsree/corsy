
# Corsy 
### CORS injector proxy for development
[![Build](https://github.com/sivsivsree/corsy/actions/workflows/unit-test.yaml/badge.svg)](https://github.com/sivsivsree/corsy/actions/workflows/unit-test.yaml) [![Go Report Card](https://goreportcard.com/badge/github.com/sivsivsree/corsy)](https://goreportcard.com/report/github.com/sivsivsree/corsy)

Corsy is a proxy injector for development to avoid CORS issues while building SPA applicaitons.


<b>Install</b>

1.  Run the below to install corsy to the system
```sh
   curl https://i.jpillora.com/sivsivsree/corsy! | bash
```

Download the latest exe artifact from the Release.

<b>Usage: </b>

```
Options:
  -a, --addr string         address:port to listen on :8080  (default ":8001")
  -b, --blacklist strings   Headers to remove from the request and response
  -h, --help                Show this message
  -r, --max-redirects int   Maximum number of redirects to follow (default 10)
  -p, --proxy string        remote address to proxy with cors
  -t, --timeout int         Request timeout (default 15)
```

<b>Example</b>
 <br>

```
corsy -a ":8080" -p https://prod.example.com/api/v1/resource

```
or 
```
corsy --addr ":8080" --proxy https://prod.example.com/api/v1/resource  \
--timeout 20 --max-redirects 2 \
--blacklist X-Remote-Url 

```
