# bluelens

[ ![Codeship Status for ihcsim/bluelens](https://app.codeship.com/projects/78e180d0-e10a-0134-d4f7-5e506c0c41eb/status?branch=master)](https://app.codeship.com/projects/205421)

bluelens makes music recommendations to its users based on the users' social activities.

## Introduction
Our fictional client has asked us to build a music recommendation system for his customers. The system shall make music recommendations based on a user's listening history and who they follow.

The following is the list of rules to determine the recommendations:

* History of all the music a user has heard before.
* Followees of user. And maybe even followees of the followees.
* Preferences to be given to songs that are new to the user.

## Getting Started
Use the Makefile to build the server and client:
```sh
$ make
```
Run the `bluelens` binary to start the server:
```sh
$ ./bluelens
2017/03/16 16:22:18 [INFO] mount ctrl=Music action=Get route=GET /bluelens/music/:musicID
2017/03/16 16:22:18 [INFO] mount ctrl=Recommendations action=Recommend route=GET /bluelens/recommendations/:userID/:maxCount
2017/03/16 16:22:18 [INFO] mount ctrl=Swagger files=server/swagger/swagger.json route=GET /swagger.json
2017/03/16 16:22:18 [INFO] mount ctrl=User action=Follow route=POST /bluelens/user/:userID/follows/:followeeID
2017/03/16 16:22:18 [INFO] mount ctrl=User action=Get route=GET /bluelens/user/:userID
2017/03/16 16:22:18 [INFO] mount ctrl=User action=Listen route=POST /bluelens/user/:userID/listen/:musicID
2017/03/16 16:22:18 [INFO] listen transport=http addr=:8080
```
Use the `blue` client to interact with the server:
```sh
$ ./blue --help
CLI client for the bluelens service

Usage:
  bluelens-cli [command]

Available Commands:
  download    Download file with given path
  follow      Update a user's followees list with a new followee.
  get         get action
  help        Help about any command
  listen      Add a music to a user's history.
  recommend   Make music recommendations for a user.

Flags:
      --dump               Dump HTTP request and response.
  -H, --host string        API hostname (default "localhost:8080")
  -s, --scheme string      Set the requests scheme
  -t, --timeout duration   Set the request timeout (default 20s)

Use "bluelens-cli [command] --help" for more information about a command.
```

## API
The `design/design.go` contains all the endpoints specifications. The swagger doc is accessible at http://localhost:8080/swagger.json

## Data Model
This system is constrained by the following pre-defined data model:

* `music`: Has an ID and a list of tags (see `etc/music.json`),
* `user`: Has an ID, follows N other users (see `etc/follows.json`), has heard Y musics in the past (see `etc/listen.json`).

## Development
This project uses [glide](https://github.com/Masterminds/glide) to manage its dependencies, and [goa](https://goa.design/) to generate the server and client code.

Use the Makefile `all` target to install the dependencies, run the test, autogenerate and build the server an client code. All the generated code, controller code and `main.go` are found in the `server` folder.
```sh
$ make
```
To force a re-generation of the controller code and `server/main.go`, set the `$CODEGEN_MAIN_OPTS` environmental variable:
```sh
$ CODEGEN_MAIN_OPTS=--force make
```
Take a look at the Makefile for other convenient targets that help with daily development tasks such as `test`, `server/build`, `client/build` etc.

## LICENSE
Refer the [LICENSE](LICENSE) file.
