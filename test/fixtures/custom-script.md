## Custom Script Test

```
> POST /post HTTP/1.1
> Accept: text/plain
> Content-Type: application/json
{{ custom }}
< HTTP/1.1 200 OK
{
  "args": {},
  "data": "this is custom\n",
  "files": {},
  "form": {},
  "headers": {
    "Accept": "text/plain",
    "Content-Length": "15",
    "Content-Type": "application/json",
    "Host": "localhost:7357",
    "User-Agent": "Go-http-client/1.1"
  },
  "json": null,
  "origin": "172.17.0.1",
  "url": "http://localhost:7357/post"
}
```