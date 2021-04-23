// +build go1.6

package request_test

import (
	"errors"
	"testing"

	"github.com/aws/mateuszwojcikcc/aws"
	"github.com/aws/mateuszwojcikcc/aws/awserr"
	"github.com/aws/mateuszwojcikcc/aws/client"
	"github.com/aws/mateuszwojcikcc/aws/client/metadata"
	"github.com/aws/mateuszwojcikcc/aws/defaults"
	"github.com/aws/mateuszwojcikcc/aws/request"
)

// go version 1.4 and 1.5 do not return an error. Version 1.5 will url encode
// the uri while 1.4 will not
func TestRequestInvalidEndpoint(t *testing.T) {
	endpoint := "http://localhost:90 "

	r := request.New(
		aws.Config{},
		metadata.ClientInfo{Endpoint: endpoint},
		defaults.Handlers(),
		client.DefaultRetryer{},
		&request.Operation{},
		nil,
		nil,
	)

	if r.Error == nil {
		t.Errorf("expect error")
	}
}

type timeoutErr struct {
	error
}

var errTimeout = awserr.New("foo", "bar", &timeoutErr{
	errors.New("net/http: request canceled"),
})

func (e *timeoutErr) Timeout() bool {
	return true
}

func (e *timeoutErr) Temporary() bool {
	return true
}
