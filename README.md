# gator

## A Boot.Dev project
The purpose of this project is to utilize Go and SQL to create an RSS feed aggregator in the cli.

## Setup

#### Dependencies
Requires Postgresql v15+ and Go v1.25+.



#### Config
This program expectes to find a .gatorconfig.json file in your home directory (~/.gatorconfig.son). Currently, it doesn't generate this file automatically. 

The config file should contain two fields: 
- "db_url": A link to your Postgresql database (i.e. "postgres://exampledb")
- "current_user_string": The current user. This can be anything to start.

An example config file is included in this repository.

#### Installation

You can install this program with the Go install command:
```go
go install github.com/neeeb1/gator@latest
```
## Usage

Gator is an rss feed aggregator. You can add users, follow specific feeds for each user, and browser the latest posts form the feed.

Commands:
```
register - adds a new user
login - sets current user to a registered user
users - lists all registered users

addfeed - adds an rss feed with a name and url
feeds - lists all available feeds
follow - follow an already added feed for the current user
following - list all feeds the current user is following

agg - fetches feeds on a provided frequency (e.g. 5s, 10s, 1m)
browse - allows you to browse the latest N feeds that the current user is following

reset - \*DANGER\* removes all data from the database!
```