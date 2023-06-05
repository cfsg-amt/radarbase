# API Package Design

The `api` package provides an interface to interact with MongoDB through an HTTP API server. It uses the `gorilla/mux` package to set up the routing and handle HTTP requests.

```go
type API struct {
	db *mdb.MDB
}
```
The API struct holds a pointer to an `mdb.MDB` instance.

## API Functions

1. **`GET /api/v1/{collectionName}/item?headers={headers}`** - This route is used to retrieve data based on headers in a collection. Headers are comma-separated. The handler for this route is **`GetByHeadersHandler`**.
2. **`GET /api/v1/{collectionName}/item/{stockName}`** - This route is used to retrieve a single record from a collection. The handler for this route is **`GetSingleRecordHandler`**.
3. **`GET /api/v1/headers/{collectionName}`** - This route is used to retrieve all headers from a collection. The handler for this route is **`GetHeadersHandler`**.
