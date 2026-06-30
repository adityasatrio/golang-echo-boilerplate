package web

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"html"
	"net/http"
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
	"myapp/configs/credential"
	"myapp/exceptions"
	"myapp/internal/applications/auth/dto"
	authService "myapp/internal/applications/auth/service"
	roleService "myapp/internal/applications/role/service"
	sysParamDto "myapp/internal/applications/system_parameter/dto"
	sysParamService "myapp/internal/applications/system_parameter/service"
	userDto "myapp/internal/applications/user/dto"
	userService "myapp/internal/applications/user/service"
	"myapp/internal/helper"
	"myapp/middleware"
)

const oauthStateCookieName = "oauth_state"

// WebHandler renders the HTMX + Alpine.js UI on top of the existing services,
// calling them in-process rather than over HTTP.
type WebHandler struct {
	authSvc     authService.AuthService
	userSvc     userService.UserService
	roleSvc     roleService.RoleService
	sysParamSvc sysParamService.SystemParameterService
}

func NewWebHandler(authSvc authService.AuthService, userSvc userService.UserService, roleSvc roleService.RoleService, sysParamSvc sysParamService.SystemParameterService) *WebHandler {
	return &WebHandler{authSvc: authSvc, userSvc: userSvc, roleSvc: roleSvc, sysParamSvc: sysParamSvc}
}

func businessErrorMessage(err error) string {
	var bizErr *exceptions.BusinessLogicError
	if errors.As(err, &bizErr) {
		return exceptions.BusinessLogicReason(bizErr.ErrorCode).Message
	}
	return "something went wrong, please try again"
}

func isAdminFromContext(c echo.Context) bool {
	roleID, ok := c.Get(middleware.ContextKeyRoleID).(uint64)
	return ok && roleID == middleware.AdminRoleID
}

func setSessionCookie(c echo.Context, token string) {
	expiryMinutes := credential.GetInt("jwt.expiryMinutes")
	if expiryMinutes <= 0 {
		expiryMinutes = 60
	}

	c.SetCookie(&http.Cookie{
		Name:     middleware.SessionCookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   c.Request().TLS != nil,
		MaxAge:   expiryMinutes * 60,
	})
}

func clearCookie(c echo.Context, name string) {
	c.SetCookie(&http.Cookie{
		Name:     name,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})
}

func htmlAlert(c echo.Context, message string) error {
	return c.HTML(http.StatusOK, `<div class="alert-error">`+html.EscapeString(message)+`</div>`)
}

// LoginPage renders GET /login.
func (h *WebHandler) LoginPage(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", LoginPageData{Error: c.QueryParam("error")})
}

// PostLogin handles POST /auth/login (username/password via Auth0 ROPC).
func (h *WebHandler) PostLogin(c echo.Context) error {
	request := new(dto.PasswordLoginRequest)
	if err := helper.BindAndValidate(c, request); err != nil {
		return htmlAlert(c, "email and password are required")
	}

	result, err := h.authSvc.LoginWithPassword(c.Request().Context(), request)
	if err != nil {
		return htmlAlert(c, businessErrorMessage(err))
	}

	setSessionCookie(c, result.Token)
	c.Response().Header().Set("HX-Redirect", "/system-parameters")
	return c.NoContent(http.StatusOK)
}

// GoogleLoginRedirect handles GET /auth/google, redirecting to Auth0's
// hosted authorize page for the Google connection.
func (h *WebHandler) GoogleLoginRedirect(c echo.Context) error {
	state, err := generateState()
	if err != nil {
		return c.Redirect(http.StatusFound, "/login?error="+url.QueryEscape("could not start google login"))
	}

	c.SetCookie(&http.Cookie{
		Name:     oauthStateCookieName,
		Value:    state,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   5 * 60,
	})

	return c.Redirect(http.StatusFound, h.authSvc.BuildGoogleAuthorizeURL(state))
}

// GoogleCallback handles GET /auth/callback after the user authenticates
// with Google on Auth0's hosted page.
func (h *WebHandler) GoogleCallback(c echo.Context) error {
	request := new(dto.GoogleLoginCallbackRequest)
	if err := helper.BindAndValidate(c, request); err != nil {
		return c.Redirect(http.StatusFound, "/login?error="+url.QueryEscape("google login failed"))
	}

	stateCookie, err := c.Cookie(oauthStateCookieName)
	if err != nil || stateCookie.Value == "" || stateCookie.Value != request.State {
		return c.Redirect(http.StatusFound, "/login?error="+url.QueryEscape("google login failed, please try again"))
	}
	clearCookie(c, oauthStateCookieName)

	result, err := h.authSvc.HandleGoogleCallback(c.Request().Context(), request.Code)
	if err != nil {
		return c.Redirect(http.StatusFound, "/login?error="+url.QueryEscape(businessErrorMessage(err)))
	}

	setSessionCookie(c, result.Token)
	return c.Redirect(http.StatusFound, "/system-parameters")
}

// Logout handles POST /auth/logout.
func (h *WebHandler) Logout(c echo.Context) error {
	clearCookie(c, middleware.SessionCookieName)
	return c.Redirect(http.StatusFound, "/login")
}

func generateState() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// roleNameByID builds an ID->name lookup used to render the role badge/select.
func (h *WebHandler) roleOptions(c echo.Context) ([]RoleOption, map[uint64]string, error) {
	roles, err := h.roleSvc.GetAll(c.Request().Context())
	if err != nil {
		return nil, nil, err
	}

	options := make([]RoleOption, 0, len(roles))
	names := make(map[uint64]string, len(roles))
	for _, r := range roles {
		options = append(options, RoleOption{ID: r.ID, Name: r.Name})
		names[r.ID] = r.Name
	}
	return options, names, nil
}

// UsersPage renders GET /users (Admin-only).
func (h *WebHandler) UsersPage(c echo.Context) error {
	ctx := c.Request().Context()

	users, err := h.userSvc.GetAll(ctx)
	if err != nil {
		return err
	}

	_, roleNames, err := h.roleOptions(c)
	if err != nil {
		return err
	}

	rows := make([]UserRow, 0, len(users))
	for _, u := range users {
		rows = append(rows, UserRow{ID: u.ID, Name: u.Name, Email: u.Email, RoleName: roleNames[u.RoleID]})
	}

	return c.Render(http.StatusOK, "users.html", UsersPageData{Active: "users", IsAdmin: true, Users: rows})
}

// UserFormNew renders GET /users/new (modal partial).
func (h *WebHandler) UserFormNew(c echo.Context) error {
	roles, _, err := h.roleOptions(c)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "user_form", UserFormData{IsEdit: false, Roles: roles})
}

// UserFormEdit renders GET /users/:id/edit (modal partial).
func (h *WebHandler) UserFormEdit(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	user, err := h.userSvc.GetById(c.Request().Context(), id)
	if err != nil {
		return err
	}

	roles, _, err := h.roleOptions(c)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "user_form", UserFormData{
		IsEdit: true,
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		RoleID: user.RoleID,
		Roles:  roles,
	})
}

// UserCreate handles POST /users.
func (h *WebHandler) UserCreate(c echo.Context) error {
	request := new(userDto.UserRequest)
	if err := helper.BindAndValidate(c, request); err != nil {
		return h.renderUserFormError(c, false, 0, "please fill in all fields correctly")
	}

	created, err := h.userSvc.Create(c.Request().Context(), request)
	if err != nil {
		return h.renderUserFormError(c, false, 0, businessErrorMessage(err))
	}

	role, err := h.roleSvc.GetById(c.Request().Context(), created.RoleID)
	roleName := ""
	if err == nil {
		roleName = role.Name
	}

	return c.Render(http.StatusOK, "user_row", UserRow{
		ID: created.ID, Name: created.Name, Email: created.Email, RoleName: roleName,
		OOBMode: "beforeend:#user-table-body",
	})
}

// UserUpdate handles PUT /users/:id.
func (h *WebHandler) UserUpdate(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	request := new(userDto.UserRequest)
	if err := helper.BindAndValidate(c, request); err != nil {
		return h.renderUserFormError(c, true, id, "please fill in all fields correctly")
	}

	updated, err := h.userSvc.Update(c.Request().Context(), id, request)
	if err != nil {
		return h.renderUserFormError(c, true, id, businessErrorMessage(err))
	}

	role, err := h.roleSvc.GetById(c.Request().Context(), updated.RoleID)
	roleName := ""
	if err == nil {
		roleName = role.Name
	}

	return c.Render(http.StatusOK, "user_row", UserRow{
		ID: updated.ID, Name: updated.Name, Email: updated.Email, RoleName: roleName,
		OOBMode: "true",
	})
}

func (h *WebHandler) renderUserFormError(c echo.Context, isEdit bool, id uint64, message string) error {
	roles, _, err := h.roleOptions(c)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "user_form", UserFormData{IsEdit: isEdit, ID: id, Roles: roles, GeneralError: message})
}

// UserDelete handles DELETE /users/:id.
func (h *WebHandler) UserDelete(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user id")
	}

	if _, err := h.userSvc.Delete(c.Request().Context(), id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

// SystemParametersPage renders GET /system-parameters (Admin + User).
func (h *WebHandler) SystemParametersPage(c echo.Context) error {
	params, err := h.sysParamSvc.GetAll(c.Request().Context())
	if err != nil {
		return err
	}

	isAdmin := isAdminFromContext(c)
	rows := make([]SystemParameterRow, 0, len(params))
	for _, p := range params {
		rows = append(rows, SystemParameterRow{ID: p.ID, Key: p.Key, Value: p.Value, IsAdmin: isAdmin})
	}

	return c.Render(http.StatusOK, "system_parameters.html", SystemParametersPageData{Active: "system-parameters", IsAdmin: isAdmin, Params: rows})
}

// SystemParameterFormNew renders GET /system-parameters/new (modal partial, Admin-only).
func (h *WebHandler) SystemParameterFormNew(c echo.Context) error {
	return c.Render(http.StatusOK, "system_parameter_form", SystemParameterFormData{IsEdit: false})
}

// SystemParameterFormEdit renders GET /system-parameters/:id/edit (modal partial, Admin-only).
func (h *WebHandler) SystemParameterFormEdit(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid parameter id")
	}

	param, err := h.sysParamSvc.GetById(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.Render(http.StatusOK, "system_parameter_form", SystemParameterFormData{IsEdit: true, ID: param.ID, Key: param.Key, Value: param.Value})
}

// SystemParameterCreate handles POST /system-parameters (Admin-only).
func (h *WebHandler) SystemParameterCreate(c echo.Context) error {
	request := new(sysParamDto.SystemParameterCreateRequest)
	if err := helper.BindAndValidate(c, request); err != nil {
		return c.Render(http.StatusOK, "system_parameter_form", SystemParameterFormData{IsEdit: false, GeneralError: "please fill in all fields correctly"})
	}

	created, err := h.sysParamSvc.Create(c.Request().Context(), request)
	if err != nil {
		return c.Render(http.StatusOK, "system_parameter_form", SystemParameterFormData{IsEdit: false, Key: request.Key, Value: request.Value, GeneralError: businessErrorMessage(err)})
	}

	return c.Render(http.StatusOK, "system_parameter_row", SystemParameterRow{
		ID: created.ID, Key: created.Key, Value: created.Value, IsAdmin: true,
		OOBMode: "beforeend:#system-parameter-table-body",
	})
}

// SystemParameterUpdate handles PUT /system-parameters/:id (Admin-only).
func (h *WebHandler) SystemParameterUpdate(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid parameter id")
	}

	request := new(sysParamDto.SystemParameterUpdateRequest)
	if err := helper.BindAndValidate(c, request); err != nil {
		return c.Render(http.StatusOK, "system_parameter_form", SystemParameterFormData{IsEdit: true, ID: id, GeneralError: "please fill in all fields correctly"})
	}

	updated, err := h.sysParamSvc.Update(c.Request().Context(), id, request)
	if err != nil {
		return c.Render(http.StatusOK, "system_parameter_form", SystemParameterFormData{IsEdit: true, ID: id, Key: request.Key, Value: request.Value, GeneralError: businessErrorMessage(err)})
	}

	return c.Render(http.StatusOK, "system_parameter_row", SystemParameterRow{
		ID: updated.ID, Key: updated.Key, Value: updated.Value, IsAdmin: true,
		OOBMode: "true",
	})
}

// SystemParameterDelete handles DELETE /system-parameters/:id (Admin-only).
func (h *WebHandler) SystemParameterDelete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid parameter id")
	}

	if _, err := h.sysParamSvc.Delete(c.Request().Context(), id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
