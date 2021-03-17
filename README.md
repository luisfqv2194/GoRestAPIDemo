# GoRestAPIDemo
Go Rest API using Gorm and Fiber V2, that allows to do CRUD operations in a SQLlite database for a movie catalog.

## Requisites
1. Go must be installed. (Download [here](https://golang.org/))
2. Postman must be installed. (Download [here](https://www.postman.com/downloads/))

## Installation
1. Clone this repo to a directory


## Introduction
Please import the file GoRestAPIDemo.postman_collection.json into your Postman.

## Avaible Postman Calls And Descriptions
### Ping Endpoint
Returns a greeting message
Endpoint call
```
"http://localhost:8080/"
```

### Get All Movies
Queries all the rows for the table Movies

Endpoint call
```
"http://localhost:8080/api/movies/"
```

### Get Movie By Parcial Name
Extracts a movie using a partial title in the search

Endpoint call
```
"http://localhost:8080/api/movies/{PARTIAL_TITLE}"
```
Replace PARTIAL_TITLE with the full name or a parical name of a movie, an error will be returned if no movie matches the title provided
### Insert New Movie
Inserts new movies into the database
Endpoint call
```
"http://localhost:8080/api/movies/"
```
Body must contain the following:
| Attribute  | Value  |
|---|---|
| Rating  |  string |
| Title  | string  |
| Duration  | int  |
| Year  | int  |
### Update Movie
Updates an existing movie record with a new movie object
```
"http://localhost:8080/api/movies/update"
```
Body must contain the following:
| Attribute  | Value  |
|---|---|
| Rating  |  string |
| Title  | string  |
| Duration  | int  |
| Year  | int  |
| ID  | int  |
### Delete Movie By ID
Deletes a movie by the specified ID on the URL
```
"http://localhost:8080/api/movies/update/{ID}"
```
Replace ID with an ID present in the database, if the ID doesn't exist in the datbase an error will be returned.

## Unit Test
To run the tests within the app_test.go file run the following in the root folder:
```
go test
```

