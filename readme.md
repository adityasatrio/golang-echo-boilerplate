<H1>Golang CRUD boilerplate</h1>
- Golang 1.18
  - https://go.dev/play/
  - Video references from PZN bahasa indo : [PZN-golang-playlist](https://www.youtube.com/watch?v=JOXbresHhIk&list=PL-CtdCApEFH-0i9dzMzLw6FKVrFWv3QvQ)
  - Tutorial bahasa indonesia [dasar golang noval agung](https://dasarpemrogramangolang.novalagung.com/1-berkenalan-dengan-golang.html)
- echo framework v4
  - https://echo.labstack.com/
  - Not the the fastest, but on par with GIN with better documentation [benchmark discussion](https://github.com/labstack/echo/discussions/2143)
  - There are also a lot of tutorial on the net using bahasa indo, and easy for beginner ! [noval agung echo framework rest api](https://dasarpemrogramangolang.novalagung.com/C-echo-routing.html) 
- Viper 
  - https://github.com/spf13/viper
- Entgo
  - https://entgo.io/
    - New kids on the block, developed by facebook team. Not the fastest, but better than gorm and have generated query builder! 
    - See benchmark : [ent benchmark](https://github.com/efectn/go-orm-benchmarks/blob/master/results.md)
- Google wire - code gen for dependency injection
  - https://github.com/google/wire
  - Good tutorial for getting started with example [tutorial google DI with google wire](https://clavinjune.dev/en/blogs/golang-dependency-injection-using-wire/)
  - [Video references from PZN - golang DI with google wire](https://www.youtube.com/watch?v=dZ8Ir4Gc8D0&list=PL-CtdCApEFH-0i9dzMzLw6FKVrFWv3QvQ&index=14)

<H3>TO DO</H3>
- [x] create clean code structure
- [x] create interface
- [x] create manual DI
- [x] implement repository + database connection using ent
- [x] implement error handling
- [x] implement DI google wire
- [ ] implement cache https://github.com/eko/gocache#a-chained-cache
- [ ] implement outbound http calls
- [ ] implement migration files 
- [ ] implement test
- [ ] implement pub sub libs intergation
- [ ] integrate swagger or API docs
- [ ] dockerize project

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
Create model schema using ent:
1. Get dependency for golang ent 
```
go get entgo.io/ent/cmd/ent
```
2. Create new model schema, the generated model located on `ent/schema/model_name.go` </br>
   https://entgo.io/docs/schema-def for schema documentation
```
go run entgo.io/ent/cmd/ent init {model_name}
```
3. Generate assets
```
go generate ./ent
```
</br>

<H3>Dependency Injection</H3>
1. Install google wire CLI
```
go install github.com/google/wire/cmd/wire@latest
```
2. Add wire on your $PATH, so we can use wire CLI on every project
3. Create {domains}_injector.go in your feature directory
4. Run wire on the same directory of your injector file 
