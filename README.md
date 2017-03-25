# bluelens

[ ![Codeship Status for ihcsim/bluelens](https://app.codeship.com/projects/78e180d0-e10a-0134-d4f7-5e506c0c41eb/status?branch=master)](https://app.codeship.com/projects/205421)

bluelens makes music recommendations to its users based on the users' social activities.

## Table of Content

* [Introduction](#introduction)
* [Getting Started](#getting-started)
* [API Design](#api-design)
* [Data Model](#data-model)
* [Development](#development)
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
* Encrypt data transmission over HTTPS.

## Getting Started
The following is the list of prerequisites to build this project:

1. Install [glide](http://glide.sh/)
1. Set the `${PACKAGE_ROOT}` in the Makefile to match your project package.
1. Install goagen using this command: `make goagen`.

To build the server and client, use the `all` target of the Makefile:
```sh
$ make
```

Run the `blued` binary to start the server:
```sh
$ blued -apikey=<api_key> -user=<user> -password=<password>
2017/03/24 22:43:05 [INFO] mount ctrl=Music action=Create route=POST /bluelens/music security=APIKey
2017/03/24 22:43:05 [INFO] mount ctrl=Music action=List route=GET /bluelens/music security=APIKey
2017/03/24 22:43:05 [INFO] mount ctrl=Music action=Show route=GET /bluelens/music/:id security=APIKey
2017/03/24 22:43:05 [INFO] mount ctrl=Recommendations action=Recommend route=GET /bluelens/recommendations/:userID/:limit security=APIKey
2017/03/24 22:43:05 [INFO] mount ctrl=Swagger files=cmd/blued/swagger/swagger.json route=GET /bluelens/swagger.json security=BasicAuth
2017/03/24 22:43:05 [INFO] mount ctrl=Swagger files=cmd/blued/swagger/swagger.yaml route=GET /bluelens/swagger.yaml security=BasicAuth
2017/03/24 22:43:05 [INFO] mount ctrl=User action=Create route=POST /bluelens/user security=APIKey
2017/03/24 22:43:05 [INFO] mount ctrl=User action=Follow route=POST /bluelens/user/:id/follows/:followeeID security=APIKey
2017/03/24 22:43:05 [INFO] mount ctrl=User action=List route=GET /bluelens/user security=APIKey
2017/03/24 22:43:05 [INFO] mount ctrl=User action=Listen route=POST /bluelens/user/:id/listen/:musicID security=APIKey
2017/03/24 22:43:05 [INFO] mount ctrl=User action=Show route=GET /bluelens/user/:id security=APIKey
2017/03/24 22:43:05 [INFO] listen transport=http addr=:8080
```
For a list of all the flags supported by `blued`, see `$ blued -h`.

Use the `blue` client to interact with the server:
```sh
$ blue -h
CLI client for the bluelens service (http://localhost:8080/swagger.json)

Usage:
  bluelens-cli [command]

Available Commands:
  create      create action
  download    Download file with given path
  follow      Update a user's followees list with a new followee.
  help        Help about any command
  list        list action
  listen      Add a music to a user's history.
  recommend   Make music recommendations for a user.
  show        show action

Flags:
      --dump               Dump HTTP request and response.
      --format string      Format used to create auth header or query from key (default "Bearer %s")
  -H, --host string        API hostname (default "localhost:8080")
      --key string         API key used for authentication
      --pass string        Password used for authentication
  -s, --scheme string      Set the requests scheme
  -t, --timeout duration   Set the request timeout (default 20s)
      --user string        Username used for authentication
```

## API Design
The `design/design.go` contains all the endpoints specifications. The Swagger doc is accessible at http://localhost:8080/bluelens/swagger.json or http://localhost:8080/bluelens/swagger.yaml

All the non-Swagger endpoints are secured by the API key scheme. The Swagger doc endpoints are secured by the HTTP Basic Authentication scheme. `blued` provides the following flags to inject the API key, username and password into the server on start:

* `apikey`: API key used to secure all the non-Swagger endpoints.
* `user`: Basic username used to secure the Swagger endpoints.
* `password`: Basic password used to secure the Swagger endpoints.

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

## LICENSE
Refer the [LICENSE](LICENSE) file.
