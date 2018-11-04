# oldiez
We have to go back. Back to the Future!!
(•_•)
( •_•)>⌐■-■
(⌐■_■)

### Prerequisites
Install goose. 
```
$ go get -u github.com/pressly/goose/cmd/goose
```
https://github.com/pressly/goose

Install go-swagger
```
$ brew tap go-swagger/go-swagger
$ brew install go-swagger
```
Refer to swagger markeup guide here: https://goswagger.io/generate/spec.html

### Create Postgres DB
1. Create the postgres db.
```
$ createdb oldiez_test
```
2. Change ownership of DB to postgres:
```
psql> alter database oldiez_test owner to postgres;
```
3. Apply the migrations from the "migrations" directory.
```
$ goose postgres "user=postgres dbname=oldiez_test sslmode=disable" up
```

Clean DB how-to:
```
$ goose postgres "user=postgres dbname=oldiez_dev sslmode=disable" down-to 0 
```

### Testing 
Apply DB schema to test database. Create dB oldiez_test if it does not exist. 

```
$ goose postgres "user=postgres dbname=oldiez_test sslmode=disable" up
```


### Generating the API docs
From within the /api project 
```
$ swagger generate spec -o ./oldiez-swagger.json --scan-models
$ swagger serve -F=swagger oldiez-swagger.json
```