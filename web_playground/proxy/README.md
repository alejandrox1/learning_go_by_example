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
