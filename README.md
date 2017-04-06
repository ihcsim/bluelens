# bluelens

[ ![Codeship Status for ihcsim/bluelens](https://app.codeship.com/projects/78e180d0-e10a-0134-d4f7-5e506c0c41eb/status?branch=master)](https://app.codeship.com/projects/205421)

bluelens makes music recommendations to its users based on the users' social activities.

## Table of Content

* [Introduction](#introduction)
* [Getting Started](#getting-started)
* [API Design](#api-design)
* [Security](#security)
* [Data Model](#data-model)
* [Development](#development)
* [ACI Image](#aci-image-with-rkt)
* [License](#license)

## Introduction
Our fictional client has asked us to build a music recommendation system for his customers. The system shall make music recommendations based on a user's listening history and who they follow.

The following is the list of rules to determine the recommendations:

* History of all the music a user has heard before.
* Followees of user. And maybe even followees of the followees.
* Preferences to be given to songs that are new to the user.

The objectives of this project is to experiment with [goa](https://goa.design/). In particular, we will be using goa to:

* Auto-generate server and client code,
* Define custom media types,
* Secure API endpoints using [Basic Authentication](https://en.wikipedia.org/wiki/Basic_access_authentication) and API key,
* Generate API swagger documents,
* Logging HTTP requests,
* Build websocket endpoints,
* Timeout middleware with [context](https://golang.org/pkg/context/),
* Encrypt data transmission over HTTPS.

The [`acbuild`](https://github.com/containers/build) is used to build an ACI container image of `blued`, which can be run with [rkt](https://github.com/rkt/rkt). Refer to the [ACI Image](#aci-image) section.

## Getting Started
The following is the list of prerequisites to build this project:

1. Install [glide](http://glide.sh/)
1. Set the `${PACKAGE_ROOT}` in the Makefile to match your project package.
1. Install goagen using this command: `make goagen`
1. Create a self-signed certificate. This content is ignored by git: `make tls`

To build the server and client, use the `all` target of the Makefile:
```sh
$ make
```

Run the `blued` binary to start the server:
```sh
$ sudo ./blued -apikey=<apikey> -user=<user> -password=<password> -private tls/localhost.key -cert tls/localhost.crt
2017/03/28 11:14:13 [INFO] mount ctrl=Music action=Create route=POST /bluelens/music security=APIKey
2017/03/28 11:14:13 [INFO] mount ctrl=Music action=List route=GET /bluelens/music security=APIKey
2017/03/28 11:14:13 [INFO] mount ctrl=Music action=Show route=GET /bluelens/music/:id security=APIKey
2017/03/28 11:14:13 [INFO] mount ctrl=Recommendations action=Recommend route=GET /bluelens/recommendations/:userID/:limit security=APIKey
2017/03/28 11:14:13 [INFO] mount ctrl=Swagger files=cmd/blued/swagger/swagger.json route=GET /bluelens/swagger.json
2017/03/28 11:14:13 [INFO] mount ctrl=Swagger files=cmd/blued/swagger/swagger.yaml route=GET /bluelens/swagger.yaml
2017/03/28 11:14:13 [INFO] mount ctrl=User action=Create route=POST /bluelens/user security=APIKey
2017/03/28 11:14:13 [INFO] mount ctrl=User action=Follow route=POST /bluelens/user/:id/follows/:followeeID security=APIKey
2017/03/28 11:14:13 [INFO] mount ctrl=User action=List route=GET /bluelens/user security=APIKey
2017/03/28 11:14:13 [INFO] mount ctrl=User action=Listen route=POST /bluelens/user/:id/listen/:musicID security=APIKey
2017/03/28 11:14:13 [INFO] mount ctrl=User action=Show route=GET /bluelens/user/:id security=APIKey
2017/03/28 11:14:13 [INFO] listen transport=https addr=:443
```
Note that the server is started with `sudo` in order to bind the process to port 443. The Swagger docs for the endpoints can be found in https://localhost/bluelens/swagger.json and https://localhost/bluelens/swagger.yaml

`blued` supports the following flags:
```sh
$ blued -h
  -apikey string
        Key used for API key authentication
  -cert string
        Path to the TLS cert file (default "tls/cert.pem")
  -followees string
        Path to read user's followees data from (default "etc/followees.json")
  -history string
        Path to read user's history data from (default "etc/history.json")
  -music string
        Path to read music data from (default "etc/music.json")
  -password string
        Password used for HTTP Basic Authentication
  -private string
        Path to the TLS private key file (default "tls/key.pem")
  -timeout duration
        Request timeout in seconds. Default to 10s. (default 10s)
  -user string
        Username used for HTTP Basic Authentication
```

**In goa 1.1.0, the auto-generated client doesn't work with TLS.**

To interact with the `blued` server, use `curl`:
```sh
# view all music resources
$ curl --cacert tls/localhost.crt -H "Authorization: Bearer <apikey>" https://localhost/bluelens/music

# view all user resources
$ curl --cacert tls/localhost.crt -H "Authorization: Bearer <apikey>" https://localhost/bluelens/user

# make 10 recommendations for user-01
$ curl --cacert tls/localhost.crt -H "Authorization: Bearer <apikey>" https://localhost/bluelens/recommendations/user-01/10
```

## API Design
The `design/design.go` contains all the endpoints specifications. The Swagger doc is accessible at https://localhost/bluelens/swagger.json and https://localhost/bluelens/swagger.yaml

## Security
All request and responses messages are transmitted over HTTPS. The `make tls` target can be used to generate a self-signed certificate and private key. To change the transporatation scheme to HTTP, update the `design/design.go` specification, and update the server's service (in `cmd/blued/main.go`) to use the Golang [`http.ListenAndServer()`](https://golang.org/pkg/net/http/#ListenAndServe) API.

In addition, the endpoints authorization scheme are detailed below:

Endpoints                    | Authorization Schemes
---------------------------- | ---------------------
`/bluelens/music`            | API key
`/bluelens/user`             | API key
`/bluelens/recommendations`  | API key
`/bluelens/swagger.json`     | Basic Authorization
`/bluelens/swagger.yaml`     | Basic Authorization

## Data Model
This system is constrained by the following pre-defined data model:

* `music`: Has an ID and a list of tags (see `etc/music.json`),
* `user`: Has an ID, follows N other users (see `etc/follows.json`), has heard Y musics in the past (see `etc/listen.json`).

## Development
This project uses [glide](https://github.com/Masterminds/glide) to manage its dependencies, and [goa](https://goa.design/) to generate the server and client code.

Use the Makefile `all` target to install the dependencies, run the test, autogenerate and build the server an client code.
```sh
$ make
```
Take a look at the Makefile for other convenient targets that help with daily development tasks such as `test`, `codegen`, `server/build`, `client/build` etc.

### Code Layout

#### Server
The server code are found in the `cmd/blued`.  The controller code and `main.go` are generated once by `goagen`. As long as they remain in the `cmd/blued` folder, subsequent `goagen` execution will not re-generate them. All the code found in the `cmd/blued/app` sub-folder are re-generated by `goagen` in subsequent execution. To force the controller code or `main.go` to be re-generated, run the Make `server/codegen` target with the `CODEGEN_MAIN_OPTS` variable set to `--force`:
```sh
$ CODEGEN_MAIN_OPTS=--force make server/codegen
```

#### Client
The client code are found in the `cmd/blue` folder. The CLI `main.go` is located in the `cmd/blue/tool/blue` sub-folder. To force the `main.go` to be re-generated, delete it then run the Make `client/codengen` target:
```sh
$ make client/codegen
```

## ACI Image With rkt
The `acbuild.sh` script uses [`acbuild`](https://github.com/containers/build) to build an ACI container image for `blued`. This image can then be run with [`rkt`](https://github.com/rkt/rkt).
```sh
$ API_KEY=<api_key> BASIC_AUTH_USERNAME=<username> BASIC_AUTH_PASSWORD=<password> make aci
$ sudo rkt run --insecure-options=image --net=host build/blued-1.0.0-linux-amd64.aci
[18762.694938] blued[5]: 2017/04/06 19:59:51 [INFO] mount ctrl=Music action=Create route=POST /bluelens/music security=APIKey
[18762.695191] blued[5]: 2017/04/06 19:59:51 [INFO] mount ctrl=Music action=List route=GET /bluelens/music security=APIKey
[18762.695349] blued[5]: 2017/04/06 19:59:51 [INFO] mount ctrl=Music action=Show route=GET /bluelens/music/:id security=APIKey
[18762.695516] blued[5]: 2017/04/06 19:59:51 [INFO] mount ctrl=Recommendations action=Recommend route=GET /bluelens/recommendations/:userID/:limit security=APIKey
[18762.695691] blued[5]: 2017/04/06 19:59:51 [INFO] mount ctrl=Swagger files=cmd/blued/swagger/swagger.json route=GET /bluelens/swagger.json
[18762.695881] blued[5]: 2017/04/06 19:59:51 [INFO] mount ctrl=Swagger files=cmd/blued/swagger/swagger.yaml route=GET /bluelens/swagger.yaml
[18762.696132] blued[5]: 2017/04/06 19:59:51 [INFO] mount ctrl=User action=Create route=POST /bluelens/user security=APIKey
[18762.696279] blued[5]: 2017/04/06 19:59:51 [INFO] mount ctrl=User action=Follow route=POST /bluelens/user/:id/follows/:followeeID security=APIKey
[18762.696486] blued[5]: 2017/04/06 19:59:51 [INFO] mount ctrl=User action=List route=GET /bluelens/user security=APIKey
[18762.696658] blued[5]: 2017/04/06 19:59:51 [INFO] mount ctrl=User action=Listen route=POST /bluelens/user/:id/listen/:musicID security=APIKey
[18762.696804] blued[5]: 2017/04/06 19:59:51 [INFO] mount ctrl=User action=Show route=GET /bluelens/user/:id security=APIKey
[18762.696953] blued[5]: 2017/04/06 19:59:51 [INFO] listen transport=https addr=:443
```
Notice that the rkt pod is started with a host bridge network so that the endpoints are accessible at https://localhost. Also, the self-signed cert generated by the `make tls` target uses localhost as the CN.

Use `curl` to access the endpoints. For example,
```sh
$ curl -H "Authorization: Bearer <api_key>" --cacert tls/localhost.crt https://localhost/bluelens/music
[{"href":"/bluelens/music/song-01","id":"song-01","tags":["jazz","old school","instrumental"]},{"href":"/bluelens/music/song-02","id":"song-02","tags":["samba","60s"]},{"href":"/bluelens/music/song-03","id":"song-03","tags":["rock","alternative"]},{"href":"/bluelens/music/song-04","id":"song-04","tags":["rock","alternative"]},{"href":"/bluelens/music/song-05","id":"song-05","tags":["folk","instrumental"]},{"href":"/bluelens/music/song-06","id":"song-06","tags":["60s","rock","old school"]},{"href":"/bluelens/music/song-07","id":"song-07","tags":["alternative","dance"]},{"href":"/bluelens/music/song-08","id":"song-08","tags":["electronic","pop"]},{"href":"/bluelens/music/song-09","id":"song-09","tags":["60s","rock"]},{"href":"/bluelens/music/song-10","id":"song-10","tags":["60s","jazz"]},{"href":"/bluelens/music/song-11","id":"song-11","tags":["samba"]},{"href":"/bluelens/music/song-12","id":"song-12","tags":["jazz","instrumental"]},{"href":"/bluelens/music/song-13","id":"song-13","tags":["80s","old school","dance"]},{"href":"/bluelens/music/song-14","id":"song-14","tags":[""]},{"href":"/bluelens/music/song-15","id":"song-15","tags":["soft rock","90s","international"]},{"href":"/bluelens/music/song-16","id":"song-16","tags":["pop","soundtrack"]},{"href":"/bluelens/music/song-17","id":"song-17","tags":["jazz","billboard","90s"]},{"href":"/bluelens/music/song-18","id":"song-18","tags":["pop","80s","international"]},{"href":"/bluelens/music/song-19","id":"song-19","tags":["soundtrack","old school"]},{"href":"/bluelens/music/song-20","id":"song-20","tags":["90s","jazz"]}]
```
To access the swagger docs,
```sh
curl -H "Authorization: Bearer <api_key>" --cacert tls/localhost.crt https://admin:admin@localhost/bluelens/swagger.json
{"swagger":"2.0","info":{"title":"The bluelens API","description":"This API provides a set of endpoints to manage users' followees, music history and recommendations.","license":{"name":"MIT","url":"https://github.com/ihcsim/bluelens/blob/master/LICENSE"},"version":""},"host":"localhost","schemes":["https"],"consumes":["application/json"],"produces":["application/json"],"paths":{"/bluelens/music":{"get":{"tags":["music"],"summary":"list music","description":"List up to N music resources. N can be adjusted using the 'limit' and 'offset' parameters.","operationId":"music#list","parameters":[{"name":"limit","in":"query","required":false,"type":"integer","default":20},{"name":"offset","in":"query","required":false,"type":"integer","default":0}],"responses":{"200":{"description":"OK","schema":{"$ref":"#/definitions/BluelensMusicCollection"}},"401":{"description":"Unauthorized"},"500":{"description":"Internal Server Error","schema":{"$ref":"#/definitions/error"}}}
...
```

## LICENSE
Refer the [LICENSE](LICENSE) file.
