package mocks

type MockAcknowledger struct{}

// Ack simulates acknowledgment
func (a *MockAcknowledger) Ack(tag uint64, multiple bool) error {
	// Simulated acknowledgment
	return nil
}

// Nack simulates negative acknowledgment
func (a *MockAcknowledger) Nack(tag uint64, multiple, requeue bool) error {
	// Simulated negative acknowledgment
	return nil
}

// Reject simulates rejecting a delivery
func (a *MockAcknowledger) Reject(tag uint64, requeue bool) error {
	// Simulated rejection
	return nil
}
