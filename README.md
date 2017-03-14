# bluelens

[ ![Codeship Status for ihcsim/bluelens](https://app.codeship.com/projects/78e180d0-e10a-0134-d4f7-5e506c0c41eb/status?branch=master)](https://app.codeship.com/projects/205421)

bluelens makes music recommendations to its users based on the users' social activities.

## Introduction
Our fictional client has asked us to build a music recommendation system for his customers. The system has to take note of what music a user has listened to, which people they follow and from there recommend some songs.

The following is the list of rules to determine the recommendations:

* History of all the music a user has heard before
* Followees of user. And maybe even followees of the followees.
* Preferences to be given to songs that are new to the user.

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
