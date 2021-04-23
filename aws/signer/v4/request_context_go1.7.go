// +build go1.7

package v4

import (
	"net/http"

	"github.com/aws/mateuszwojcikcc/aws"
)

func requestContext(r *http.Request) aws.Context {
	return r.Context()
}
