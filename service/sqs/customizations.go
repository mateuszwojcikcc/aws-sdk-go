package sqs

import "github.com/aws/mateuszwojcikcc/aws/request"

func init() {
	initRequest = func(r *request.Request) {
		setupChecksumValidation(r)
	}
}
