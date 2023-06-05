# MDB Package Documentation

## Design

The `mdb` package provides a simple and convenient abstraction over MongoDB operations in Go, which makes it easier to interact with a MongoDB database.

The central struct is `MDB`, which encapsulates a MongoDB client, a database reference, and a context. This design makes it simple to perform operations on a specific MongoDB database by instantiating an `MDB` object.

Methods in the `MDB` struct are designed to perform typical database operations such as `Insert`, `Delete`, `Get`, and `GetAll`. These methods encapsulate the underlying MongoDB operations, providing a simpler and more intuitive interface.

The `GetHeaders` method is a specialized function designed to retrieve specific fields (headers) from all documents in a MongoDB collection. It provides a way to retrieve a subset of data from the database without needing to fetch all data.

## Trade-offs

- **Simplicity vs Flexibility**: The `mdb` package is designed with simplicity in mind, which means it does not expose all features of the underlying MongoDB driver. For example, it doesn't provide methods for aggregations, indexing, transactions, etc. If you need to perform more complex operations, you may need to extend the `mdb` package or use the MongoDB driver directly.

- **Error Handling**: The package provides minimal error handling, returning errors directly from the MongoDB driver. You may want to add more sophisticated error handling, such as retries, logging, or specialized error types, depending on your application's needs.

- **Context Management**: The `MDB` struct holds a context that is used for all operations. This design is simple, but it assumes that all operations will have the same context requirements (timeout, cancellation, etc.). If you need different context settings for different operations, you might need to modify the design to pass the context as a parameter to each method.

- **Data Modeling**: The package assumes that all data in the database can be represented as `bson.M` (an unordered map) or slices of `bson.M`. This is a flexible representation that can handle any JSON-like data, but it doesn't provide any type safety or support for more complex data models. If your application uses complex data models, you might need to add methods that operate on those models directly.
