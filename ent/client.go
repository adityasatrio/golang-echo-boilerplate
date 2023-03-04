// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"myapp/ent/migrate"

	"myapp/ent/pet"
	"myapp/ent/post"
	"myapp/ent/system_parameter"
	"myapp/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Pet is the client for interacting with the Pet builders.
	Pet *PetClient
	// Post is the client for interacting with the Post builders.
	Post *PostClient
	// System_parameter is the client for interacting with the System_parameter builders.
	System_parameter *SystemParameterClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Pet = NewPetClient(c.config)
	c.Post = NewPostClient(c.config)
	c.System_parameter = NewSystemParameterClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transaction client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		Pet:              NewPetClient(cfg),
		Post:             NewPostClient(cfg),
		System_parameter: NewSystemParameterClient(cfg),
		User:             NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transaction client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:              ctx,
		config:           cfg,
		Pet:              NewPetClient(cfg),
		Post:             NewPostClient(cfg),
		System_parameter: NewSystemParameterClient(cfg),
		User:             NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Pet.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Pet.Use(hooks...)
	c.Post.Use(hooks...)
	c.System_parameter.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Pet.Intercept(interceptors...)
	c.Post.Intercept(interceptors...)
	c.System_parameter.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *PetMutation:
		return c.Pet.mutate(ctx, m)
	case *PostMutation:
		return c.Post.mutate(ctx, m)
	case *SystemParameterMutation:
		return c.System_parameter.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// PetClient is a client for the Pet schema.
type PetClient struct {
	config
}

// NewPetClient returns a client for the Pet from the given config.
func NewPetClient(c config) *PetClient {
	return &PetClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `pet.Hooks(f(g(h())))`.
func (c *PetClient) Use(hooks ...Hook) {
	c.hooks.Pet = append(c.hooks.Pet, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `pet.Intercept(f(g(h())))`.
func (c *PetClient) Intercept(interceptors ...Interceptor) {
	c.inters.Pet = append(c.inters.Pet, interceptors...)
}

// Create returns a builder for creating a Pet entity.
func (c *PetClient) Create() *PetCreate {
	mutation := newPetMutation(c.config, OpCreate)
	return &PetCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Pet entities.
func (c *PetClient) CreateBulk(builders ...*PetCreate) *PetCreateBulk {
	return &PetCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Pet.
func (c *PetClient) Update() *PetUpdate {
	mutation := newPetMutation(c.config, OpUpdate)
	return &PetUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PetClient) UpdateOne(pe *Pet) *PetUpdateOne {
	mutation := newPetMutation(c.config, OpUpdateOne, withPet(pe))
	return &PetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PetClient) UpdateOneID(id uuid.UUID) *PetUpdateOne {
	mutation := newPetMutation(c.config, OpUpdateOne, withPetID(id))
	return &PetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Pet.
func (c *PetClient) Delete() *PetDelete {
	mutation := newPetMutation(c.config, OpDelete)
	return &PetDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PetClient) DeleteOne(pe *Pet) *PetDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PetClient) DeleteOneID(id uuid.UUID) *PetDeleteOne {
	builder := c.Delete().Where(pet.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PetDeleteOne{builder}
}

// Query returns a query builder for Pet.
func (c *PetClient) Query() *PetQuery {
	return &PetQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePet},
		inters: c.Interceptors(),
	}
}

// Get returns a Pet entity by its id.
func (c *PetClient) Get(ctx context.Context, id uuid.UUID) (*Pet, error) {
	return c.Query().Where(pet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PetClient) GetX(ctx context.Context, id uuid.UUID) *Pet {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PetClient) Hooks() []Hook {
	return c.hooks.Pet
}

// Interceptors returns the client interceptors.
func (c *PetClient) Interceptors() []Interceptor {
	return c.inters.Pet
}

func (c *PetClient) mutate(ctx context.Context, m *PetMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PetCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PetUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PetDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Pet mutation op: %q", m.Op())
	}
}

// PostClient is a client for the Post schema.
type PostClient struct {
	config
}

// NewPostClient returns a client for the Post from the given config.
func NewPostClient(c config) *PostClient {
	return &PostClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `post.Hooks(f(g(h())))`.
func (c *PostClient) Use(hooks ...Hook) {
	c.hooks.Post = append(c.hooks.Post, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `post.Intercept(f(g(h())))`.
func (c *PostClient) Intercept(interceptors ...Interceptor) {
	c.inters.Post = append(c.inters.Post, interceptors...)
}

// Create returns a builder for creating a Post entity.
func (c *PostClient) Create() *PostCreate {
	mutation := newPostMutation(c.config, OpCreate)
	return &PostCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Post entities.
func (c *PostClient) CreateBulk(builders ...*PostCreate) *PostCreateBulk {
	return &PostCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Post.
func (c *PostClient) Update() *PostUpdate {
	mutation := newPostMutation(c.config, OpUpdate)
	return &PostUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PostClient) UpdateOne(po *Post) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPost(po))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PostClient) UpdateOneID(id int) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPostID(id))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Post.
func (c *PostClient) Delete() *PostDelete {
	mutation := newPostMutation(c.config, OpDelete)
	return &PostDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PostClient) DeleteOne(po *Post) *PostDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PostClient) DeleteOneID(id int) *PostDeleteOne {
	builder := c.Delete().Where(post.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PostDeleteOne{builder}
}

// Query returns a query builder for Post.
func (c *PostClient) Query() *PostQuery {
	return &PostQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePost},
		inters: c.Interceptors(),
	}
}

// Get returns a Post entity by its id.
func (c *PostClient) Get(ctx context.Context, id int) (*Post, error) {
	return c.Query().Where(post.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PostClient) GetX(ctx context.Context, id int) *Post {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PostClient) Hooks() []Hook {
	return c.hooks.Post
}

// Interceptors returns the client interceptors.
func (c *PostClient) Interceptors() []Interceptor {
	return c.inters.Post
}

func (c *PostClient) mutate(ctx context.Context, m *PostMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PostCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PostUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PostDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Post mutation op: %q", m.Op())
	}
}

// SystemParameterClient is a client for the System_parameter schema.
type SystemParameterClient struct {
	config
}

// NewSystemParameterClient returns a client for the System_parameter from the given config.
func NewSystemParameterClient(c config) *SystemParameterClient {
	return &SystemParameterClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `system_parameter.Hooks(f(g(h())))`.
func (c *SystemParameterClient) Use(hooks ...Hook) {
	c.hooks.System_parameter = append(c.hooks.System_parameter, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `system_parameter.Intercept(f(g(h())))`.
func (c *SystemParameterClient) Intercept(interceptors ...Interceptor) {
	c.inters.System_parameter = append(c.inters.System_parameter, interceptors...)
}

// Create returns a builder for creating a System_parameter entity.
func (c *SystemParameterClient) Create() *SystemParameterCreate {
	mutation := newSystemParameterMutation(c.config, OpCreate)
	return &SystemParameterCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of System_parameter entities.
func (c *SystemParameterClient) CreateBulk(builders ...*SystemParameterCreate) *SystemParameterCreateBulk {
	return &SystemParameterCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for System_parameter.
func (c *SystemParameterClient) Update() *SystemParameterUpdate {
	mutation := newSystemParameterMutation(c.config, OpUpdate)
	return &SystemParameterUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SystemParameterClient) UpdateOne(sp *System_parameter) *SystemParameterUpdateOne {
	mutation := newSystemParameterMutation(c.config, OpUpdateOne, withSystem_parameter(sp))
	return &SystemParameterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SystemParameterClient) UpdateOneID(id int) *SystemParameterUpdateOne {
	mutation := newSystemParameterMutation(c.config, OpUpdateOne, withSystem_parameterID(id))
	return &SystemParameterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for System_parameter.
func (c *SystemParameterClient) Delete() *SystemParameterDelete {
	mutation := newSystemParameterMutation(c.config, OpDelete)
	return &SystemParameterDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SystemParameterClient) DeleteOne(sp *System_parameter) *SystemParameterDeleteOne {
	return c.DeleteOneID(sp.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SystemParameterClient) DeleteOneID(id int) *SystemParameterDeleteOne {
	builder := c.Delete().Where(system_parameter.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SystemParameterDeleteOne{builder}
}

// Query returns a query builder for System_parameter.
func (c *SystemParameterClient) Query() *SystemParameterQuery {
	return &SystemParameterQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSystemParameter},
		inters: c.Interceptors(),
	}
}

// Get returns a System_parameter entity by its id.
func (c *SystemParameterClient) Get(ctx context.Context, id int) (*System_parameter, error) {
	return c.Query().Where(system_parameter.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SystemParameterClient) GetX(ctx context.Context, id int) *System_parameter {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SystemParameterClient) Hooks() []Hook {
	return c.hooks.System_parameter
}

// Interceptors returns the client interceptors.
func (c *SystemParameterClient) Interceptors() []Interceptor {
	return c.inters.System_parameter
}

func (c *SystemParameterClient) mutate(ctx context.Context, m *SystemParameterMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SystemParameterCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SystemParameterUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SystemParameterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SystemParameterDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown System_parameter mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Pet, Post, System_parameter, User []ent.Hook
	}
	inters struct {
		Pet, Post, System_parameter, User []ent.Interceptor
	}
)
