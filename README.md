# Snippets

A RESTful API example for snippets app with GO

## Installation and run
```sh
go get github.com/mbredikhin/snippets
cd snippets
# Build services
make build
# Create and start containers
make run
# Run database migrations
make migrate 
# API Endpoint: http://127.0.0.1:8001
```

## API

#### /auth/sign-up
* `POST` : Register a new user
#### /auth/sign-in
* `POST` : Login with username and password
#### /api/lists
* `POST` : Create a new list of snippets
* `GET` : Get all lists
#### /api/lists/:id
* `GET` : Get list
* `PUT` : Update list
* `DELETE` : Delete list
#### /api/lists/:id/snippets
* `POST` : Create new snippet in a list
* `GET` : Get snippets from a list
#### /api/snippets/:id
* `GET` : Get snippet
* `PUT` : Update snippet
* `DELETE` : Delete snippet
#### /api/favourite-snippets
* `GET` : Get favourite snippet's IDs
#### /api/favourite-snippets/:id
* `POST` : Add snippet to favourites
* `DELETE` : Delete snippet from favourites
#### /api/tags
* `POST` : Create new tag
* `GET` : Get tags
#### /api/tags/:id
* `GET` : Get tag
* `PUT` : Update tag
* `DELETE` : Delete tag
#### /api/syntaxes
* `GET` : Get list of syntaxes
