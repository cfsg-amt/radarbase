# Packages

## pkg/mdb

`mdb` package is a layer above MongoDB, allowing it to handle requests, transform them into MongoDB queries, and return processed data to the user. 

### Data Structure

#### Database

```go
type Database struct {
	client   *mongo.Client
	database *mongo.Database
}
```

The client field of the Database type is used to interact with MongoDB at a higher level than the database or collection. Some operations in MongoDB are performed at the client level rather than the database or collection level.

### Testing

1. `mdb_test.go`: This file will contain the TestMain function that sets up and tears down the test database. It will also load the test data into the MongoDB collection.
2. `handler_test.go`: This file will test the GetAllStocksWithSelectedHeaders and GetAllHeadersForStock functions.`
3. `loader_test.go`: This file will print all items in the database for testing.


### Database

Wrapping the MongoDB client into a custom Database type is not necessary but is often a best practice in software development. Encapsulating the database operations into a specific type allows for **better code organization, reusability, and the ability to add custom behavior or state that might be necessary for application.** 

It also makes it easier to swap out the underlying database technology, if needed, **as the rest of your codebase is interacting with the `Database` type and not directly with MongoDB.** 

Lastly, it makes it easier to mock the Database type for unit testing.

## pkg/excel

The excel package only contains a parser for parsing the excel file and transform it to `[]map[string]interface{}` format.

The parser is built above a go package called **[excelize](https://github.com/qax-os/excelize)**

## pkg/api

The api package provides the main HTTP API interfaces for the radarbase application. It is responsible for handling the incoming HTTP requests and responding with appropriate data fetched from the database.

This package provides the primary means of interaction for clients to access and manipulate the stock data stored in the MongoDB database.

Key components of this package include:

* `Handler`: The `Handler` struct contains the MongoDB database instance and provides methods to handle HTTP requests. The methttps://github.com/gorilla/muxhods include `GetStocksHandler`,  `GetStockByIDHandler` and `GetHeadersHandler` for fetching all stocks, a specific stock and all the headers, respectively, from a particular collection.
* `NewRouter`: This function sets up the routes for the API server. It uses the **[gorilla/mux](https://github.com/gorilla/mux)** package to create routes that map to the Handler's methods.
* `StartAPIServer`: This function starts the HTTP server on a given port and uses the router created by NewRouter. It's responsible for starting the HTTP server and listening for incoming requests.

### Example Usage:
The api package is typically used by creating an instance of the Handler, setting up routes with NewRouter, and then starting the server with StartAPIServer.

```go
db, _ := mdb.NewDatabase("mongodb://localhost:27017", "testdb")

handler := &api.Handler{DB: db}

r := api.NewRouter(handler)

api.StartAPIServer(r)
```

This will start the server, and the Handler methods can be accessed via HTTP at the routes defined by NewRouter.
