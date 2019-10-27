# GraphQL api 

### Dependencies 
gqlgen 
https://github.com/99designs/gqlgen
https://gqlgen.com/

Regenerate the support files from graphql dir.
```
$ go run github.com/99designs/gqlgen -v
```

### Running 
This api is meant to be run within a kube cluster. Locally you can run it within minikube.

```
make start
```