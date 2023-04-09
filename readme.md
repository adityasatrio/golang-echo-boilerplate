# Golang CRUD boilerplate
  
- Golang 1.18
  - https://go.dev/play/
  - Video references from PZN bahasa indo : [PZN-golang-playlist](https://www.youtube.com/watch?v=JOXbresHhIk&list=PL-CtdCApEFH-0i9dzMzLw6FKVrFWv3QvQ)
  - Tutorial bahasa indonesia [dasar golang noval agung](https://dasarpemrogramangolang.novalagung.com/1-berkenalan-dengan-golang.html)
- echo framework v4
  - https://echo.labstack.com/, Not the the fastest, but on par with GIN with better documentation [benchmark discussion](https://github.com/labstack/echo/discussions/2143)
  - There are also a lot of tutorial on the net using bahasa indo, and easy for beginner ! [noval agung echo framework rest api](https://dasarpemrogramangolang.novalagung.com/C-echo-routing.html) 
- Viper 
  - https://github.com/spf13/viper, commonly used and powerfull configuration libs
- Entgo
  - https://entgo.io/, New kids on the block, developed by facebook team. Not the fastest, but better than gorm and have generated query builder!
  - See benchmark : [ent benchmark](https://github.com/efectn/go-orm-benchmarks/blob/master/results.md)
- Google wire - code gen for dependency injection
  - https://github.com/google/wire, DI code generator
  - Good tutorial for getting started with example [tutorial google DI with google wire](https://clavinjune.dev/en/blogs/golang-dependency-injection-using-wire/)
  - [Video references from PZN - golang DI with google wire](https://www.youtube.com/watch?v=dZ8Ir4Gc8D0&list=PL-CtdCApEFH-0i9dzMzLw6FKVrFWv3QvQ&index=14)
- testify : {to add explanation later}
- mockery : {to add explanation later}
- Cache : {to add explanation later}
- Logging : {to add explanation later}
- Message brooker : {to add explanation later}
- API documentation : {to add explanation later}
- Containerization : {to add explanation later}

### Todo
- [x] create clean code structure
- [x] create interface with example domains system param
- [x] create manual DI on hello worlds example domains
- [x] implement repository + database connection using ent in system param example domains
- [ ] implement optimistic locking https://github.com/ent/ent/blob/master/examples/version/README.md 
- [x] implement global error handling
- [x] implement DI google wire
- [ ] implement cache, use go cache as cache abstraction https://github.com/eko/gocache#a-chained-cache
- [ ] implement migration files, instead of using ent / atlas we decide to use https://pkg.go.dev/github.com/golang-migrate/migrate/v4 for easiness and simplicity
- [ ] Example - implement test for CRUD example
  - great talk by imre ! [Writing Better Test in Go | Go Srilanka Meetup June 25th, 2021](https://www.youtube.com/watch?v=xTQI_4EKB8Y)
  - using testify for the sake concise code, no needs write many if for assertion
  - using mockery
- [ ] implement cache https://github.com/eko/gocache#a-chained-cache
- [ ] Example - implement outbound http calls - support config timeout, retry, circuit breaker
- [ ] Example - implement transaction examples - https://entgo.io/docs/transactions
- [ ] Example - implement migration files - https://entgo.io/docs/data-migrations, https://atlasgo.io/
- [ ] implement logger
- [ ] integrate live reload
- [ ] implement message brooker integration - rabbitmq
- [ ] integrate swagger or API docs - https://github.com/swaggo/echo-swagger
- [ ] dockerize project / Colima 

### Example
- Hello world CRUD example using MVC pattern - [add link]
- System parameter for feature flag as CRUD example using MVC & repository pattern - [add link]
- Transaction CRUD example using ent go - [add link]
- Publish and subscribe using rabbitmq - [add link]

### Running apps
```
go mod tidy
go run cmd/main.go
```
</br>

### Build Apps
```
go build cmd/main.go
./main
```
</br>

### ORM
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

### Dependency Injection
1. Install google wire CLI
```
go install github.com/google/wire/cmd/wire@latest
```
2. Add wire on your $PATH, so we can use wire CLI on every project
3. Create {domains}_injector.go in your feature directory
4. Run wire on the same directory of your injector file 

### Mockery
generate mock file for all interface in domains
```
 mockery --all --dir internal/applications --output mocks --keeptree --packageprefix mock_
```