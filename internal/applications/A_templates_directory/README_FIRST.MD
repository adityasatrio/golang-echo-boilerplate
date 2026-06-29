This directory and files are only example,
you can use to copy and paste to your new task or feature that you will be developed

This directory and files works as guidelines, and you can change your future directory feature / task based on your need

## Dependency Injection Setup (Builder Pattern)

To wire this service into the application:

### 1. Add a builder method in `internal/builder/services.go`

```go
// BuildTemplateService builds the template service with all dependencies.
func (c *Container) BuildTemplateService() service.TemplateService {
    repo := repository.NewTemplateRepository(c.db)
    cacheDep := c.BuildCache()  // if your service needs cache
    
    return service.NewTemplateService(repo, cacheDep)
}
```

### 2. Register routes in `internal/adapter/rest/routes_setup.go`

```go
// Template service
templateService := container.BuildTemplateService()
controller.NewTemplateController(templateService).AddRoutes(e, appName)
```

### 3. Run `make mocks` to generate repository mocks

```shell
make mocks
```

That's it! No build commands or code generation needed. The builder pattern provides explicit, type-safe dependency injection.