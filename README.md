# FBLA Quiz App

Quiz website about interesting facts from FBLA

# Run
To run this program locally follow these steps:
1. Make sure you have the most recent version of Go installed.
1. Make sure you have NPM installed
1. Make sure you have MongoDB installed and copied the most recent version of the database over.
1. Run the following Go command: `go run backend/main.go`
1. Cd into the frontent directory and then run `ng serve --open`

# Tests
Not implemented yet

# Documentation
Not implement yet

# Structure
This is how everything is stuctured
## Backend
Corresponds to Uncle Bob's clean archeticture. Every component (e.g. users, questions, etc) has the following components.
* driver (third party implementations, for example MongoDB interactions)
* entities (structs or entites, for example the struct User)
* usescases (business logic)
* http.go (routing logic)

Errors are handled through Slugs in the errorCodes folder.

## Frontent
Not implemented yet.
