package initialization

/*
import (
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() *validator.Validate {
	return validator.New()
}

func (cv *CustomValidator) validate(i interface{}) error {
	return cv.validator.Struct(i)
}*/

/*func ReqBody(c echo.Context, reqBody interface{}) error {
	//TODO need create interface and inject on main
	//e.Validator = &CustomValidator{validator: validator.New()}

	if err := c.Validate(reqBody); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}*/

/*


func ValidateReqParamInt(c echo.Context, paramName string, paramInt int, validatorTags string) *echo.HTTPError {
	fmt.Println("debug", paramName, paramInt, validatorTags)

	paramStr := c.Param(paramName)
	intVar, err := strconv.Atoi(paramStr)
	if err != nil {
		//error parse then param is string
		if err := echo.PathParamsBinder(c).String(paramName, &paramStr).BindError(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

	}
	//param is int
	if err := echo.PathParamsBinder(c).Int(paramName, &intVar).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	fmt.Println("debug parse", intVar)

	if "" != validatorTags {
		//validatorTags value example "gt=1,lt=10"
		if err := validate.Var(paramInt, validatorTags); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	return nil
}
*/

/*
func ValidateReqParamString(c echo.Context, paramName string, paramString string, validatorTags string) *echo.HTTPError {
	if err := echo.PathParamsBinder(c).String(paramName, &paramString).BindError(); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if "" != validatorTags {
		//validatorTags value example "gt=1,lt=10"
		if err := validate.Var(paramString, validatorTags); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
	}

	return nil
}
*/
