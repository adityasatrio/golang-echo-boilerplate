WIP project

<H3>Running apps</H3>
```go
go mod tidy
go run cmd/main.go
```
</br>

<H3>ORM</H3>
ORM using : https://entgo.io/docs </br>
Create model schema using ent:
1. get dependency for golang ent 
```go
go get entgo.io/ent/cmd/ent
```
2. create new model schema, the generated model located on `ent/schema/model_name.go` </br>
   https://entgo.io/docs/schema-def for schema documentation
```go
go run entgo.io/ent/cmd/ent init {model_name}
```
3. generate assets
```go
go generate ./ent

```

<H3>TO DO</H3>
1. create clean code structure : done
2. create 
3. create manual injection from controller -> user-case
4. implement repository + database connection
5. immplement DI google wire 
6. implement cache 