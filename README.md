# appa
We have to go back. Back, ...to the Future!!
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

### Apply the local-config map to your local kubernetes cluster. 
Apply dev environment variables for DB_URL:
```
kubectl create -f configs/local-config-map.yaml
```
Create cluster role to minikube:
```
kubectl create -f configs/cluster-role-admin.yml
```

**Edit your pg_hba.conf file to trust kubernetes connections if necessary. 
```
host    appa_dev      postgres        192.168.1.0/24          trust
```
Restart postgres via brew services:
```
brew services restart postgres
```


### Testing 
Apply DB schema to test database. Create dB appa_test if it does not exist. 

```
$ goose postgres "user=postgres dbname=appa_test sslmode=disable" up
```


### Generating the API docs
From within the /api project 
```
$ swagger generate spec -o ./appa-swagger.json --scan-models
$ swagger serve -F=swagger appa-swagger.json
```
