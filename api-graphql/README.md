# GraphQL api 

### Dependencies 
gqlgen 
https://github.com/99designs/gqlgen
https://gqlgen.com/

### Generate the graphql related support files under graphql/

1. Update the model defs in gqlgen.yml

2. Regenerate the support files from graphql dir.
```
$ cd graphql
$ go run github.com/99designs/gqlgen -v
```

### Running 
This api is meant to be run within a kube cluster. Locally you can run it within minikube.

```
make start
```