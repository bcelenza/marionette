# Mockingbird

Mockingbird is a lightweight HTTP server designed to simulate an upstream service that may not be behaving ideally. You can use Mockingbird to simulate various conditions from the perspective of a client in a controlled environment.

Mockingbird was built out of the need to test client resilience in a service mesh architecture.

Things you can do with it:

* Set the HTTP response body from the request body
* Set HTTP response code
* Make a response purposely latent

## Run It

```bash
make image
docker run -it --rm --net host --env HTTP_PORT=80 mockingbird
```

## Examples

Echo a request back to the caller:

```bash
curl -d '{"foo":"bar"}' http://localhost:8080/echo
{"foo":"bar"}
```

Echo a request back with a different response code:

```bash
curl -v -d '{"foo":"bar"}' "http://localhost:8080/echo?status=418"
< HTTP/1.1 418 I'm a teapot
< Date: Sat, 28 Sep 2019 17:41:41 GMT
< Content-Length: 13
< Content-Type: text/plain; charset=utf-8
{"foo":"bar"}
```

Echo a request back with induced latency:

```bash
curl -d '{"foo":"bar"}' "http://localhost:8080/echo?latency=5s"
# ... 5 seconds later
{"foo":"bar"}
```

Combine status codes and latency:

```bash
curl -v -d '{"foo":"bar"}' "http://localhost:8080/echo?status=418&latency=5s"
# ... 5 seconds later
< HTTP/1.1 418 I'm a teapot
{"foo":"bar"}
```
