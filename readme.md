WIP project

<H3>Running apps</H3>
```
go mod tidy
go run cmd/main.go
```
</br>

<H3>Build Apps</H3>
```
go build cmd/main.go
./main
```
</br>

<H3>ORM</H3>
ORM using : https://entgo.io/docs </br>
Create model schema using ent:
1. get dependency for golang ent 
```
go get entgo.io/ent/cmd/ent
```
2. create new model schema, the generated model located on `ent/schema/model_name.go` </br>
   https://entgo.io/docs/schema-def for schema documentation
```
go run entgo.io/ent/cmd/ent init {model_name}
```
3. generate assets
```
go generate ./ent
```
</br>

<H3>Dependency Injection</H3>
DI using : https://github.com/google/wire </br>
1. install google wire CLI
```
go install github.com/google/wire/cmd/wire@latest
```
2. add wire on your $PATH, so we can use wire CLI on every project

<H3>TO DO</H3>
1. create clean code structure : done
2. create interface
3. create manual injection from controller -> user-case
4. implement repository + database connection
5. immplement DI google wire 
6. implement cache 