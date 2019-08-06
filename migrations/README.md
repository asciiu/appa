# goose migrations 

### Prerequisites
Install goose. 
```
$ go get -u github.com/pressly/goose/cmd/goose
```
https://github.com/pressly/goose

### Postgres database setup
1. Create the postgres dbs.
```
$ createdb appa_test
$ createdb appa_dev
```
2. Change ownership of DB to postgres:
```
psql> alter database appa_test owner to postgres;
psql> alter database appa_dev owner to postgres;
```
3. Apply the migrations from the "migrations" directory.
```
$ goose postgres "user=postgres dbname=appa_test sslmode=disable" up
$ goose postgres "user=postgres dbname=appa_dev sslmode=disable" up
```

Clean DB how-to:
```
$ goose postgres "user=postgres dbname=appa_dev sslmode=disable" down-to 0 
```

### How to generate a new migration
```
$ goose create add_some_column sql
```

