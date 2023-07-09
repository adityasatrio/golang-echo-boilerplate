package controller

import (
	"github.com/labstack/echo/v4"
	"myapp/internal/applications/user/dto"
	"myapp/internal/applications/user/service"
	"myapp/internal/apputils"
	"myapp/internal/apputils/response"
	"strconv"
)

type UserController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{service: service}
}

// CreateUser godoc
//
//	@summary		Create new user
//	@description	Create new user
//	@tags			user
//	@accept			json
//	@produce		json
//	@param			body	body		dto.UserRequest	true	"Create User DTO"
//	@success		201		{object}	response.body{data=dto.UserResponse}
//	@failure		422		{object}	response.body
//	@failure		500		{object}	response.body
//	@router			/user [post]
func (c *UserController) Create(ctx echo.Context) error {
	request := new(dto.UserRequest)
	err := apputils.BindAndValidate(ctx, request)
	if err != nil {
		//return apputils.BadRequest(ctx, err)
		return err
	}

	created, err := c.service.Create(ctx.Request().Context(), request)
	if err != nil {
		return err
	}

	var responseDto = new(dto.UserResponse)
	err = apputils.Mapper(&responseDto, created)
	if err != nil {
		return err
	}

	return response.Created(ctx, responseDto)
}

// UpdateUser godoc
//
//	@summary		Update a user
//	@description	Update a user
//	@tags			user
//	@accept			json
//	@produce		json
//	@param			id		path		int				true	"User's ID"
//	@param			body	body		dto.UserRequest	true	"Update User DTO"
//	@success		200		{object}	response.body{data=dto.UserResponse}
//	@failure		404		{object}	response.body
//	@failure		422		{object}	response.body
//	@failure		500		{object}	response.body
//	@router			/user/{id} [put]
func (c *UserController) Update(ctx echo.Context) error {
	request := new(dto.UserRequest)
	err := apputils.BindAndValidate(ctx, request)
	if err != nil {
		return err
	}

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	updated, err := c.service.Update(ctx.Request().Context(), uint64(id), request)
	if err != nil {
		return err
	}

	var responseDto = new(dto.UserResponse)
	err = apputils.Mapper(&responseDto, updated)
	if err != nil {
		return err
	}

	return response.Success(ctx, responseDto)
}

// DeleteUser godoc
//
//	@summary		Delete a user
//	@description	Delete a user
//	@tags			user
//	@accept			json
//	@produce		json
//	@param			id	path		int	true	"User's ID"
//	@success		200	{object}	response.body{data=dto.UserResponse}
//	@failure		422	{object}	response.body
//	@failure		500	{object}	response.body
//	@router			/user/{id} [delete]
func (c *UserController) Delete(ctx echo.Context) error {

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	deleted, err := c.service.Delete(ctx.Request().Context(), uint64(id))
	if err != nil {
		return err
	}

	var responseDto = new(dto.UserResponse)
	err = apputils.Mapper(&responseDto, deleted)
	if err != nil {
		return err
	}

	return response.Success(ctx, responseDto)
}

// GetUserById godoc
//
//	@summary		Get a user by id
//	@description	Get a user by id
//	@tags			user
//	@accept			json
//	@produce		json
//	@param			id	path		int	true	"User's ID"
//	@success		200	{object}	response.body{data=dto.UserResponse}
//	@failure		422	{object}	response.body
//	@failure		500	{object}	response.body
//	@router			/user/{id} [get]
func (c *UserController) GetById(ctx echo.Context) error {

	idString := ctx.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	result, err := c.service.GetById(ctx.Request().Context(), uint64(id))
	if err != nil {
		return err
	}

	var responseDto = new(dto.UserResponse)
	err = apputils.Mapper(&responseDto, result)
	if err != nil {
		return err
	}

	return response.Success(ctx, responseDto)
}

// GetAllUsers godoc
//
//	@summary		Get all users
//	@description	Get all users
//	@tags			user
//	@accept			json
//	@produce		json
//	@success		200	{object}	response.body{data=[]dto.UserResponse}
//	@failure		422	{object}	response.body
//	@failure		500	{object}	response.body
//	@router			/user [get]
func (c *UserController) GetAll(ctx echo.Context) error {
	results, err := c.service.GetAll(ctx.Request().Context())
	if err != nil {
		return err
	}

	var responseDtos []*dto.UserResponse
	for _, result := range results {
		responseDto := new(dto.UserResponse)
		err = apputils.Mapper(responseDto, result)
		if err != nil {
			return err
		}
		responseDtos = append(responseDtos, responseDto)
	}

	return response.Success(ctx, results)
}
