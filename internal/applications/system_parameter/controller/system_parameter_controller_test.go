package controller

type SystemParameterCaseStub struct{}

/*
func (caseStub *SystemParameterCaseStub) Hello(ctx context.Context) (string, error) {
	//TODO implement me
	return "hello from stub", nil
}

func (caseStub *SystemParameterCaseStub) CreateSystemParameter(ctx context.Context) (*system_parameter.SystemParameter, error) {
	//TODO implement me
	return nil, nil
}

func (caseStub *SystemParameterCaseStub) UpdateSystemParameter(ctx context.Context) (*system_parameter.SystemParameter, error) {
	//TODO implement me
	return nil, nil
}

func (caseStub *SystemParameterCaseStub) DeleteSystemParameter(ctx context.Context) error {
	//TODO implement me
	return nil
}

func (caseStub *SystemParameterCaseStub) GetSystemParameterById(ctx context.Context) (*system_parameter.SystemParameter, error) {
	//TODO implement me
	return nil, nil
}

func (caseStub *SystemParameterCaseStub) GetSystemParameterAll(ctx context.Context) ([]*system_parameter.SystemParameter, error) {
	//TODO implement me
	return nil, nil
}
*/
/*func TestHello(t *testing.T) {
	// SetupRouteHandler
	e := echo.New()
	request := httptest.NewRequest(http.MethodGet, "/system-parameter", nil)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	ctx := e.NewContext(request, recorder)

	//setup stub
	useCaseStub := &SystemParameterCaseStub{}
	h := NewSystemParameterController(useCaseStub)

	//test global_handler
	expected := "\"Hello, World! hello from stub\"\n"
	err := h.Hello(ctx)

	//assert global_handler result
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.ErrorCode)
		assert.Equal(t, expected, recorder.Body.String())
	}

	//assert use case stub
	actualResult, err := h.service.Hello(ctx.Request().Context())

	//assert use case stub
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.ErrorCode)
		assert.Equal(t, "hello from stub", actualResult)
	}

}*/
