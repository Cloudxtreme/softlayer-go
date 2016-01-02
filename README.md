# softlayer-go

Golang bindings for SoftLayer's API. This project is a work in progress. Interfaces might break.

##Getting
To get this library, simply `go get` it:

```
go get github.com/sudorandom/softlayer-go
```

##Building
###Generating Sources
This client is generated dynamically from API metadata. That way, the type information is known and users have pre-defined types to use for API arguments and responses. To regenerate the sources, use go generate.

```
go generate ./...
```
