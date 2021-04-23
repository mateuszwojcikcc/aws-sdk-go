package awstesting

import (
	"github.com/aws/mateuszwojcikcc/aws"
	"github.com/aws/mateuszwojcikcc/aws/client"
	"github.com/aws/mateuszwojcikcc/aws/client/metadata"
	"github.com/aws/mateuszwojcikcc/aws/defaults"
)

// NewClient creates and initializes a generic service client for testing.
func NewClient(cfgs ...*aws.Config) *client.Client {
	info := metadata.ClientInfo{
		Endpoint:    "http://endpoint",
		SigningName: "",
	}
	def := defaults.Get()
	def.Config.MergeIn(cfgs...)

	if v := aws.StringValue(def.Config.Endpoint); len(v) > 0 {
		info.Endpoint = v
	}

	return client.New(*def.Config, info, def.Handlers)
}
