package apigateway

import (
	"github.com/aws/mateuszwojcikcc/aws/client"
	"github.com/aws/mateuszwojcikcc/aws/request"
)

func init() {
	initClient = func(c *client.Client) {
		c.Handlers.Build.PushBack(func(r *request.Request) {
			r.HTTPRequest.Header.Add("Accept", "application/json")
		})
	}
}
