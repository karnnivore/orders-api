# Backend API for Order Application: Built with Go

## Project Description
Build an application which is able to:
1. Show to the user all orders sorted by creation date, showing the photo and some information nearby
2. Allow the user to filter by the attributes: Status, Size, Condition and Type
3. Create an api endpoint that will allow a user or machine to create, and update an order
## Technologies used:
* [GoLang](https://go.dev)
* [Gin Web Framework](https://github.com/gin-gonic/gin)

## Setup
* [Go version 1.18.1](https://go.dev/dl/)

To run the api use the command in the main.go folder
```
go run main.go
```

## Considerations/Future additions
* Update data model to match future SQL tables
* Add Database connectivity 
* Add Authentication capabilties (OAuth or Auth0)
* Create proper controller/repository pattern for endpoints/database interactivity
* Add environment variables