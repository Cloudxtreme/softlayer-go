# softlayer-go

Golang bindings for SoftLayer's API. This project is a work in progress. Interfaces can (and will) break.

##Building
###Generating Sources
This client is generated dynamically from API metadata. That way, the type information is known and users have pre-defined types to use for API arguments and responses. To regenerate the sources, use go generate.

```
go generate ./...
```

##TODOs

* Also generate service calls
* Evaluate options for handling filters
* Evaluate options for handling masks
* Examples
* Tests
