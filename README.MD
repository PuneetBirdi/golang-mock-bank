##### Summary
This is a work in progress project for me to implement things as I learn golang. The commits are pretty intermittent, in most cases when I would like to add a new feature I will spend some time learning, researching and trying to understand the concept before putting it into practice here.
________________________

##### Current State
I've got a Docker setup as well as successful queries and tables with a PostgreSQL database - currently using the SQLC code library to generate queries (I think I'd like to make another project where I can write my queries raw just for more exposure to it). I've also created tests, and built out some database transactions using the standard golang Tx functionality.

_________________________

#### Database Schema
This is the current database schema, it's quite simple. The next change I'll probably make is adding a users table when I started to work on authentication.

<img src="https://user-images.githubusercontent.com/50375451/235549458-33852dad-3204-4f97-92d2-428422528952.png" height="300"/>


