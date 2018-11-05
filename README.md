# oldiez
We have to go back. Back to the Future!!
```
(•_•)
( •_•)>⌐■-■
(⌐■_■)
```

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
1. Create the postgres dbs.
```
$ createdb oldiez_test
$ createdb oldiez_dev
```
2. Change ownership of DB to postgres:
```
psql> alter database oldiez_test owner to postgres;
psql> alter database oldiez_dev owner to postgres;
```
3. Apply the migrations from the "migrations" directory.
```
$ goose postgres "user=postgres dbname=oldiez_test sslmode=disable" up
$ goose postgres "user=postgres dbname=oldiez_dev sslmode=disable" up
```

Clean DB how-to:
```
$ goose postgres "user=postgres dbname=oldiez_dev sslmode=disable" down-to 0 
```

### Apply the local-config map to your local kubernetes cluster. 
Apply dev environment variables for DB_URL:
```
kubectl create -f local-config-map.yaml
```

**Edit your pg_hba.conf file to trust kubernetes connections if necessary. 
```
host    oldiez_dev      postgres        192.168.1.0/24          trust
```
Restart postgres via brew services:
```
brew services restart postgres
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