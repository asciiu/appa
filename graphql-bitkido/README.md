# This is Bit-ki-do. It is a disciple, an art, a gamble. 

The base project was generated via:
```
go run github.com/99designs/gqlgen init
```

This created the following package layout:

├── go.mod
├── go.sum
├── gqlgen.yml               - The gqlgen config file, knobs for controlling the generated code.
├── graph
│   ├── generated            - A package that only contains the generated runtime
│   │   └── generated.go
│   ├── model                - A package for all your graph models, generated or otherwise
│   │   └── models_gen.go
│   ├── resolver.go          - The root graph resolver type. This file wont get regenerated
│   ├── schema.graphqls      - Some schema. You can split the schema into as many graphql files as you like
│   └── schema.resolvers.go  - the resolver implementation for schema.graphql
└── server.go                - The entry point to your app. Customize it however you see fit


Note: server.go was renamed to main.go to follow convention. 

## To regen the files from an updated schema.graphqls
At the top of our resolver.go, between package and import, add the following line:

//go:generate go run github.com/99designs/gqlgen
This magic comment tells go generate what command to run when we want to regenerate our code. To run go generate recursively over your entire project, use this command:

```
go generate ./...
```