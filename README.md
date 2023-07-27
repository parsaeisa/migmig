# migmig

**Migmig** is a tool for migrating records from a database to another one. 

It simply selects records from origin databse and inserts them into the destination database. 

> My first goal is that migmig support sql databases. 

## Config 

- Origin database connection
- Destination database connection
- Bulk or normal

## Commands

### Table name
It takes name of the table that it wants to read records from. 

The table name in the destination database is optional. If it is passed by the user, a table with that name is created but if it had empty value, the name will be the same as table name in the origin database. 

## Deployment option

It has deployments using `Helm` and it can be easily deployed to your Kubernetes environment using the command below in the root directory:
```bash
helm install migmig .deploy/helm
```
