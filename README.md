# migmig

**Migmig** is a tool for migrating records from a database to another one. 

It simply selects records from origin databse and inserts them into the destination database. 

> My first goal is that migmig support sql databases. 

## Config 

- Origin database connection
- Destination database connection
- Bulk or normal
- Origin table name
- Destination table name

The table name in the destination database is optional. If it is passed by the user, a table with that name is created but if it had empty value, the name will be the same as table name in the origin database. 
## Commands

- migrate
- connection_test

## Deployment option

It has deployments using `Helm` and it can be easily deployed to your Kubernetes environment using the command below in the root directory:
```bash
helm install migmig .deploy/helm
```
## Upcoming features

**Test connection**

This is a command that checks the connection to databases. 

**Chunkenized migrating**

If the process of migration becaome interrupted the tool must do something. 

Using `limit` & `offset` I should be able to complete the migration in multiple steps. 

## Architecture

It has an interface with two methods:
```go
type DatabaseAgent interface {
  Read() error
  BulkRead() error
  Insert() error
  BulkInsert() error 
}
```

This interface can be implemented for any sql databse. The first goal is **MariaDB**.
