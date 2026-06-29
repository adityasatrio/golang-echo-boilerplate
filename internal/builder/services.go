package builder

import (
	"myapp/internal/applications/cache"
	cacheApp "myapp/internal/applications/cache"
	exampleInbound "myapp/internal/applications/example/rabbitmq/repository/inbound"
	exampleService "myapp/internal/applications/example/rabbitmq/service"
	healthRepo "myapp/internal/applications/health/repository"
	healthService "myapp/internal/applications/health/service"
	quotesRepo "myapp/internal/applications/quotes/repository/outbound"
	quotesService "myapp/internal/applications/quotes/service"
	roleRepo "myapp/internal/applications/role/repository"
	roleService "myapp/internal/applications/role/service"
	roleUserRepo "myapp/internal/applications/role_user/repository"
	"myapp/internal/component/rabbitmq/channel"
	systemParamRepo "myapp/internal/applications/system_parameter/repository/db"
	systemParamService "myapp/internal/applications/system_parameter/service"
	userRepo "myapp/internal/applications/user/repository"
	userService "myapp/internal/applications/user/service"
)

// BuildUserService builds the user service with all dependencies.
func (c *Container) BuildUserService() userService.UserService {
	userRepository := userRepo.NewUserRepository(c.db)
	roleRepository := roleRepo.NewRoleRepository(c.db)
	roleUserRepository := roleUserRepo.NewRoleUserRepository(c.db)
	trx := c.BuildTrx()
	cacheDep := c.BuildCache()

	return userService.NewUserService(userRepository, roleRepository, roleUserRepository, trx, cacheDep)
}

// BuildHealthService builds the health service with all dependencies.
func (c *Container) BuildHealthService() healthService.HealthService {
	healthRepository := healthRepo.NewHealthRepository(c.db)
	cachingService := c.BuildCachingService()

	return healthService.NewHealthService(healthRepository, cachingService)
}

// BuildRoleService builds the role service with all dependencies.
func (c *Container) BuildRoleService() roleService.RoleService {
	roleRepository := roleRepo.NewRoleRepository(c.db)

	return roleService.NewRoleService(roleRepository)
}

// BuildSystemParameterService builds the system parameter service with all dependencies.
func (c *Container) BuildSystemParameterService() systemParamService.SystemParameterService {
	systemParameterRepository := systemParamRepo.NewSystemParameterRepository(c.db)
	cacheDep := c.BuildCache()

	return systemParamService.NewSystemParameterService(systemParameterRepository, cacheDep)
}

// BuildQuotesService builds the quotes service with all dependencies.
func (c *Container) BuildQuotesService() quotesService.QuotesService {
	quoteOutbound := quotesRepo.NewQuoteOutbound()

	return quotesService.NewQuotesService(quoteOutbound)
}

// BuildCachingService builds the caching service with all dependencies.
func (c *Container) BuildCachingService() cacheApp.CachingService {
	return cache.NewCachingService(c.redis)
}

// BuildExampleRabbitMQService builds the example rabbitmq service with all dependencies.
func (c *Container) BuildExampleRabbitMQService() exampleService.ExampleRabbitMQService {
	systemParameterRepository := systemParamRepo.NewSystemParameterRepository(c.db)

	return exampleService.NewExampleRabbitMQService(systemParameterRepository)
}

// BuildExampleRabbitMQInbound builds the example rabbitmq inbound with all dependencies.
func (c *Container) BuildExampleRabbitMQInbound() exampleInbound.ExampleRabbitMQInbound {
	wrappedChannel := channel.NewWrappedChannel(c.rabbit)
	exampleSvc := c.BuildExampleRabbitMQService()
	producer := c.BuildProducer()

	return exampleInbound.NewExampleRabbitMQInbound(wrappedChannel, exampleSvc, producer)
}
