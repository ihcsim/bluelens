> Please feel free to ask us questions.

Create a README.md file with simple instructions to run it.

The focus of this assignment is on the **web back end** (or **server side**). We will be looking specifically for good practices and your development process.

You are not supposed to solve this problem with the tools you have most skill at, rather it is recommended to use **Node.js with either Express.js or Koa.js** as frameworks. Alternatively, feel free to use Go language with a framework of your choice or lastly, Ruby with Sinatra.

Also, use **MongoDB** for this assignment. All the languages refered have libraries and drivers for this DB.

## Fictional MVP!

This fictional client has asked for a recommendation system for his social music player system.
He wants you to essentially take note of what music the user has listened to, which people they follow and from there recommend some songs. There is no like or dislike so far.

In this system there are few "elements"; 

- **music**: have an ID and a list of tags (see `music.json`)
- **users**: have an ID, follow N other users, have heard Y musics in the past. 

How to model or index this data is up to you.

### There should be 3 end points

##### `POST /follow`
Add one follow relationship (see `follows.json`)

the request body have 2 params:
- from: \<user ID\>
- to: \<user ID\>

##### `POST /listen`
Add one song as the user has just listened ( see `listen.json` )

the request body have 2 params:
- user: \<user ID\>
- music: \<music ID\>

##### `GET /recommendations`
Return 5 music recommendations to this user, they should be sorted by relevance

Query string has:
- user: \<user ID\>

response looks like:

```json
{
  "list": ["<music ID>", "<music ID>", "<music ID>", "<music ID>", "<music ID>"]
}
```

--

It's supposed to be a simplistic recommendation engine, which takes into account these main components:
- based on what music they heard before
- people who the user follow in first degree, and maybe even folowees of follows
- maximize for discovery of new songs

#### make it run!

We expect 2 parts:

1. a server that only has business logic (the endpoints) with the DB, it should load `musics.json` upon server start, but other files will be loaded by:
2. a series of commands that load the data through your endpoints and finally get a recommendation (see `script.md`)

Finally, make any type of runner that starts your server and runs the script. Whether the server will be running or not after the results finish is up to you. It's also ok to have one command to put the server up running and another to run script.

#### hints
- there isn't one right answer, but the modeling skills matters
- also, don't worry about finding a perfect solution to this, it's a MVP
- implement the script correctly
- make simple instructions to execute the server and the script

Before you start or shoot us questions (don't hesitate ;) please read and fill `QnA.md`
