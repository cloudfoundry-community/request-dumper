### Summary

`request-dumper` is an app designed to run in CloudFoundry (and potentially other PaaSes),
which dumps out a summary of the request it received. This is useful for ensuring that apps
you have deployed are getting the correct headers and data passed through to them. It should
handle any URIs you pass to it, and any HTTP methods.

### Running Locally

To run locally, simply `go get github.com/cloudfoundry-community/request-dumper`, and run
`request-dumper`. It will launch in the foreground, and log messages about requests it's
processing, what interface/port it's listening on, and any errors it encounters.

The default port is `8080`, but can be overridden by setting the PORT environment variable.
To access it, you can `curl localhost:8080` (or replace 8080 with whatever PORT you choose).

### Running on CF

1. `git clone https://github.com/cloudfoundry-community/request-dumper`
2. `cd request-dumper`
3. `cf push request-dumper`
4. `curl request-dumper.<cf-domain>`

### Example Output

Spawning the server locally:

```
$ ./request-dumper
request_dumper: 2016/03/29 09:27:16 INFO: Listening on ':8080'
request_dumper: 2016/03/29 09:27:23 INFO: Processing Request: HTTP/1.1 GET /
request_dumper: 2016/03/29 09:29:44 INFO: Processing Request: HTTP/1.1 POST /
^C
$
```

Querying the app:

```
$ curl localhost:8080
GET / HTTP/1.1
Host: localhost:8080
Accept: */*
User-Agent: curl/7.43.0

$ curl localhost:8080 -d asdf
POST / HTTP/1.1
Host: localhost:8080
Accept: */*
Content-Type: application/x-www-form-urlencoded
User-Agent: curl/7.43.0

asdf
$
```
