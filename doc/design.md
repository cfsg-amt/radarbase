
# Packages

## mdb

### Testing

1. `mdb_test.go`: This file will contain the TestMain function that sets up and tears down the test database. It will also load the test data into the MongoDB collection.
2. `handler_test.go`: This file will test the GetAllStocksWithSelectedHeaders and GetAllHeadersForStock functions.`
3. `loader_test.go`: This file will print all items in the database for testing.


### Database

Wrapping the MongoDB client into a custom Database type is not necessary but is often a best practice in software development. Encapsulating the database operations into a specific type allows for **better code organization, reusability, and the ability to add custom behavior or state that might be necessary for application.** 

It also makes it easier to swap out the underlying database technology, if needed, **as the rest of your codebase is interacting with the `Database` type and not directly with MongoDB.** 

Lastly, it makes it easier to mock the Database type for unit testing.

