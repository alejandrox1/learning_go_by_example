# [Proxies](https://mauricio.github.io/golang-proxies/#/)
## Types
* Database proxies 
* TCP proxies
* HTTP proxies

## Uses
* SSL termination
* Connection management
* Protocol upgrade/downgrade
* Security, auditing, load balancing, caching, compression, etc.

Correctly handle hop-by-hop headers:
* Proxy-\*
* Upgrade
* Keep-alive
* Transfer encoding

Forward original client headers
* x-forwarded-proto
* x-forwarded-host

Streaming:
* Transfer-encoding: chuncked
* Large content-length values

You cannot have a content-length and transfer-encoding: chuncked at the
same time.

Respect (try) cache-control headers.

Be careful with buffers: do not read data without bounds from request bodies,
make sure you are buffering and have clear limits on how much memory or
connections you can use.

Be specific on errors: No chencked support? return 411 Length Required. Request
does not contain authentication details? 401 Unauthorized. Request contains
authentication but creds are invalid? 403 Forbidden.

Use a Via/Server header to define the source responses.

Log everything/nothing:
Some headers are sensitive.
Some request bodies are sensitive.

Routing, athenticating, logging, rate limiting, health checks.

### Health checking
Don't reuse connections.
Set user-agent header.
Set timeouts.
Allow for header overwrites.
Log the actual error if it fails.

### Limit Connections
Go's standard dialer and HTTP client do not limit connections, you can run off
file handles if you do not limit them.
