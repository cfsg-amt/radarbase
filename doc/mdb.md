## API methods

* `RowCollection(collectionName string) *mongo.Collection`:
    * Description: Returns the specified collection from the row database.
    * Parameters:
        * `collectionName`: Name of the collection in the row database.

* `ColCollection(collectionName string) *mongo.Collection`:
    * Description: Returns the specified collection from the columnar database.
    * Parameters:
        * `collectionName`: Name of the collection in the columnar database.

* `NewMDB(uri string, rowDBName string, colDBName string) (*MDB, error)`:
    * Description: Creates a new instance of the MDB struct.
    * Parameters:
        * `uri`: The MongoDB URI.
        * `rowDBName`: Name of the row database.
        * `colDBName`: Name of the columnar database.

* `Disconnect(ctx context.Context) error`:
    * Description: Disconnects from MongoDB.
    * Parameters:
        * `ctx`: A context with a deadline for the operation.

* `Drop(ctx context.Context) error`:
    * Description: Drops both the row and columnar databases.
    * Parameters:
        * `ctx`: A context with a deadline for the operation.

* `Ping() error`:
    * Description: Pings MongoDB.

* `InsertOneRow(ctx context.Context, collectionName string, document interface{}) error`:
    * Description: Inserts a single record into the row database.
    * Parameters:
        * `ctx`: A context with a deadline for the operation.
        * `collectionName`: Name of the collection in the row database.
        * `document`: The document to be inserted.

* `InsertOneCol(ctx context.Context, collectionName string, document interface{}) error`:
    * Description: Inserts a single record into the columnar database.
    * Parameters:
        * `ctx`: A context with a deadline for the operation.
        * `collectionName`: Name of the collection in the columnar database.
        * `document`: The document to be inserted.
