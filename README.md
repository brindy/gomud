# godmud 

I'm learning Go (golang? whatever), so decided to write a MUD.

## Objective

* Learn the language basics
* Learn how to modularise 
* Learn how to read/write files
* Learn how to parse JSON
* Learn how to accept network connections
* Learn how to read/write to the network connections
* Learn how to deal with concurrency issues

## Architecture

* The layout of the dungeon will be defined in JSON. However, logic associated with areas in the map will be written into the code directly.
* Code will be structured into packages.

### Modules / Packages

* `main` - bootstraps the MUD server
* `dungeonmap` - code related to reading the dungeon map JSON file

More to come.

## Running

`go run main.go dungeonmap/dungeonmap.json`