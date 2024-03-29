package utils

import (
	"github.com/labstack/gommon/log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func CheckLimitRetry(msg amqp.Delivery) int64 {

	// Read the "x-death" property from the message header:
	xDeath, ok := msg.Headers["x-death"].([]interface{})
	if !ok || len(xDeath) == 0 {
		return 0
	}

	// Get the "x-death" information from the first element:
	firstDeathInfo, ok := xDeath[0].(amqp.Table)
	if !ok {
		log.Warn("Unable to decipher 'x-death' information.")
		return 0
	}

	return firstDeathInfo["count"].(int64)
}

func IsHasExceeded(limit int64, count int64, msg amqp.Delivery) bool {
	if count >= limit {
		err := msg.Ack(false)
		if err != nil {
			log.Errorf("Failed to Ack a message: %v", err)
			return false
		}
		return true
	} else {
		err := msg.Reject(false)
		if err != nil {
			log.Errorf("Failed to Reject a message: %v", err)
			return false
		}
		return false
	}
}

func GetContentType() string {
	return "application/json"
}
