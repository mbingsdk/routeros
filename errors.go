package routeros

import (
	"errors"
	"fmt"

	"github.com/go-routeros/routeros/proto"
)

var (
	errAlreadyConnected = errors.New("Connect() or ConnectTLS() has already been called")
	errAlreadyAsync     = errors.New("Async() has already been called")
)

// UnknownReplyError records the sentence whose Word is unknown.
type UnknownReplyError struct {
	Sentence *proto.Sentence
}

func (err *UnknownReplyError) Error() string {
	return fmt.Sprintf("unknown RouterOS reply word: %s", err.Sentence.Word)
}

// DeviceError records the sentence containing the error received from the device.
// The sentence may have Word !trap or !fatal.
type DeviceError struct {
	Sentence *proto.Sentence
}

func (err *DeviceError) Error() string {
	m := err.Sentence.Map["message"]
	if m == "" {
		m = fmt.Sprintf("unknown error: %s", err.Sentence)
	}
	return fmt.Sprintf("RouterOS: %s", m)
}