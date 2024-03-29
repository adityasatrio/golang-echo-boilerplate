package controller

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"myapp/internal/applications/system_parameter/dto"
	"myapp/internal/component/rabbitmq/config"
	"myapp/internal/component/rabbitmq/producer"
	"myapp/internal/helper/response"
)

type ExampleRabbitMQController struct {
	producer producer.Producer
}

func NewExampleRabbitMQController(producer producer.Producer) *ExampleRabbitMQController {
	return &ExampleRabbitMQController{producer: producer}
}

// PublishRabbitMQ is controller to create new example publish.
//
//	@summary		This just sample for publish don't use this for any feature
//	@description	This just sample for publish don't use this for any feature
//	@tags			example publish
//	@accept			json
//	@produce		json
//	@success		201	{object}	response.body{data=dto.SystemParameterResponse}
//	@failure		422	{object}	response.body
//	@failure		500	{object}	response.body
//	@router			/example/rabbitmq [get]
func (ctl *ExampleRabbitMQController) PublishRabbitMQ(c echo.Context) error {

	// example request:
	request := dto.SystemParameterCreateRequest{
		Key:   "example_rabbit_key",
		Value: "example_rabbit_values",
	}

	// parsing to json:
	messageBody, err := json.Marshal(request)
	if err != nil {
		log.Errorf("failed parsing data: %v", err)
		messageBody = []byte("")
	}

	// Run factory to publish rabbitmq:
	// Send Direct in level service please aware just sample:
	data := config.NewRabbitMQConfigExample()
	execute, err := ctl.producer.SendToDirect(data, messageBody)
	if err != nil {
		log.Errorf("failed sent to direct: %v", err)
		return err
	}

	log.Infof("Sent a message: %\n", execute)
	return response.Success(c, dto.SystemParameterResponse{})
}
