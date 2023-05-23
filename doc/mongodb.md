# MongoDB

## NoSQL Database

NoSQL is a type of database that provides a mechanism for storage and retrieval of data that is modeled differently than the tabular relations used in relational databases (like MySQL, PostgreSQL). 

NoSQL databases are particularly useful for dealing with large sets of distributed data. 

Imagine we have data about users in a music streaming application. Each user can have their personal information and a list of songs they have listened to.

### Relational Database

In a relational database, such as MySQL, you would typically have at least two tables: one for the users and one for the songs. The Users table might look like this:

```sql
Users
+--------+----------+----------+------------+
| UserID | FirstName| LastName | Email      |
+--------+----------+----------+------------+
| 1      | John     | Doe      | jdoe@email.com |
| 2      | Jane     | Smith    | jsmith@email.com |
+--------+----------+----------+------------+
```

And the Songs table might look like this:

```sql
Songs
+--------+------------+-------------+
| UserID | SongName   | Artist      |
+--------+------------+-------------+
| 1      | Song A     | Artist A    |
| 1      | Song B     | Artist B    |
| 2      | Song C     | Artist C    |
| 2      | Song B     | Artist B    |
+--------+------------+-------------+
```
To find all the songs a user has listened to, you would have to join the two tables on the UserID field.

### Non-relational Database

In MongoDB, you could store all the data for each user in a single document in a Users collection, like so:

```
Users Collection
{
  {
    "UserID": 1,
    "FirstName": "John",
    "LastName": "Doe",
    "Email": "jdoe@email.com",
    "Songs": [
      {"SongName": "Song A", "Artist": "Artist A"},
      {"SongName": "Song B", "Artist": "Artist B"}
    ]
  },
  {
    "UserID": 2,
    "FirstName": "Jane",
    "LastName": "Smith",
    "Email": "jsmith@email.com",
    "Songs": [
      {"SongName": "Song C", "Artist": "Artist C"},
      {"SongName": "Song B", "Artist": "Artist B"}
    ]
  }
}
```

Each user's song history is directly embedded within their document. To find the songs a user has listened to, you just need to access the Songs array in their document. This eliminates the need for joins and can make read operations faster.

## Scaling

When the application becomes very popular and there's a large amount of data, you may need to distribute the data across multiple servers to handle the load. This is called **sharding.**

In a SQL database, the Users and Songs tables are separate. If you're trying to distribute your data across multiple servers, you have to carefully decide how to divide each table and maintain the relationships between them. This can get complex.

However, in MongoDB, each user and their songs are stored in one document. **You can easily distribute the documents across multiple servers without breaking any relationships, since all related data is contained within each document.**

For example, you can shard data by UserID, so all data for a particular user resides on a single shard. This can greatly improve performance by localizing data.

## Flexible

MongoDB is flexible and schema-less, meaning that it can adapt to changes in the data structure over time. This is one of the key advantages of MongoDB and NoSQL databases in general.

In a relational database like MySQL, the schema (structure of the data) is fixed when a table is created. If you want to add a new column to a table, you need to alter the table schema, which can be complex and time-consuming for large tables.

In contrast, MongoDB is a document-oriented database and does not enforce a fixed schema. Each document (analogous to a "row" in a relational database) in a MongoDB collection (analogous to a "table") can have its own unique structure. One document can have fields that another document doesn't. So, if your team members change some headers in the Excel files, MongoDB can handle this with no changes needed on the database side.


## Types of DB

* **NoSQL databases** are often well-suited to big data and real-time web applications, largely due to their **flexibility, scalability, and speed.** They are particularly good when you have data that doesn't fit well into table structures, such as hierarchical or graph data, or when you want to store, retrieve, and analyze large volumes of data quickly without needing complex transactions or joins.
* **Relational databases** are excellent when data integrity and consistency are critical. They are well-suited to applications that involve complex transactions or require complex queries involving multiple tables, like many traditional business applications. The structured nature of SQL databases is great for enforcing data consistency and for situations where you need your database to enforce business rules.






