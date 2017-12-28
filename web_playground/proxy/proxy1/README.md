# [HTTP(S) Proxy in Golang in less than 100 lines of code](https://medium.com/@mlowicki/http-s-proxy-in-golang-in-less-than-100-lines-of-code-6a51c2f2c38c)

1. Package [http](https://golang.org/pkg/net/http/) provides HTTP client and server implementations.
2. [HTTP CONNECT](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods/CONNECT) 
   method. It tells the proxy server to establish TCP connection
   with destination server and when done to proxy the TCP stream to and from
   the client. This way proxy server won’t terminate SSL but will simply pass
   data between client and destination server so these two parties can
   establish secure connection.


Presented code is not a production-grade solution. It lacks e.g.:
* Handling [hop-by-hop headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers#hbh)
* setting up [timeouts](The complete guide to Go net/http timeouts) while copying data between two connections or the ones
  exposed by net/http
  
  
In our server HTTP/2 support has been deliberately removed because then
[hijacking is not
possible](https://github.com/golang/go/issues/14797#issuecomment-196103814).



## [Anatomy of an HTTP Transaction](http://blog.catchpoint.com/2010/09/17/anatomyhttp/)
HTTP data rides above the TCP protocol, which guarantees reliability of
delivery, and breaks down large data requests and responses into
network-manageable chunks. TCP is a “connection” oriented protocol, which means
when a client starts a dialogue with a server the TCP protocol will open a
connection, over which the HTTP data will be reliably transferred, and when the
dialogue is complete that connection should be closed. All of the data in the
HTTP protocol is expressed in human-readable ASCII text.

1. DNS Lookup: The client tries to resolve the domain name for the request.
 * Client sends DNS Query to local ISP DNS server.
 * DNS server responds with the IP address for hostname.com 
2. Connect: Client establishes TCP connection with the IP address of hostname.com
 * Client sends SYN packet.
 * Web server sends SYN-ACK packet.
 * Client answers with ACK packet, concluding the three-way TCP connection establishment.
3. Send: Client sends the HTTP request to the web server.
4. Wait: Client waits for the server to respond to the request.
Web server processes the request, finds the resource, and sends the response to
the Client. Client receives the first byte of the first packet from the web
server, which contains the HTTP Response headers and content.
5. Load: Client loads the content of the response.
 * Web server sends second TCP segment with the PSH flag set.
 * Client sends ACK. (Client sends ACK every two segments it receives. from the host)
 * Web server sends third TCP segment with HTTP_Continue.
6. Close: Client sends a a FIN packet to close the TCP connection.

### Serial transactions
A Serial HTTP connection occurs when multiple requests are issued sequentially
to a server, and each request establishes a new connection. This method rarely
occurs today because all modern browsers support parallel connections to a
host. However, this may also happen when a browser or server supports only HTTP
1.0, without Keep Alive (or HTTP 1.0 +) and the first request is a blocking
request (for example an inline JavaScript Request)

### Persistent transactions
Persistent connections allow the browser / HTTP client to utilize the same
connection for different object requests to the same hostname. The HTTP 1.1
protocol supports persistent connections natively, and does not require any
specific HTTP header information. For HTTP 1.0, persistent connections are
controlled via the Keep-Alive HTTP header.


