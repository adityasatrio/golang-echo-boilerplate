package utils

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"testing"
)

func TestCheckLimitRetry(t *testing.T) {

	tests := []struct {
		name      string
		delivery  amqp.Delivery
		wantCount int64
	}{
		{
			name: "Valid x-death header",
			delivery: amqp.Delivery{
				Headers: amqp.Table{
					"x-death": []interface{}{
						amqp.Table{
							"count": int64(3),
						},
					},
				},
			},
			wantCount: 3,
		},
		{
			name:      "Invalid x-death header",
			delivery:  amqp.Delivery{},
			wantCount: 0,
		},
		{
			name: "Unparseable x-death header",
			delivery: amqp.Delivery{
				Headers: amqp.Table{
					"x-death": []interface{}{
						"invalid_value",
					},
				},
			},
			wantCount: 0,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := CheckLimitRetry(tt.delivery); gotCount != tt.wantCount {
				t.Errorf("CheckLimitRetry() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}

}

func TestIsHasExceeded(t *testing.T) {

	tests := []struct {
		name        string
		limit       int64
		count       int64
		delivery    amqp.Delivery
		expectation bool
	}{
		{
			name:  "Valid x-death header",
			limit: 3,
			count: 3,
			delivery: amqp.Delivery{
				Headers: amqp.Table{
					"x-death": []interface{}{
						amqp.Table{
							"count": int64(3),
						},
					},
				},
			},
			expectation: true,
		},
		{
			name:  "Valid x-death header",
			limit: 3,
			count: 1,
			delivery: amqp.Delivery{
				Headers: amqp.Table{
					"x-death": []interface{}{
						amqp.Table{
							"count": int64(1),
						},
					},
				},
			},
			expectation: false,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isHasExceeded := IsHasExceeded(tt.limit, tt.count, tt.delivery)
			if isHasExceeded && tt.expectation {
				t.Errorf("TestIsHasExceeded() = %v, want %v", tt.expectation, isHasExceeded)
			}
		})
	}

}

func TestGetContentType(t *testing.T) {
	GetContentType()
}
