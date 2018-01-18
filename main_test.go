package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request SNSMessage
		err     error
	}{
		{
			request: SNSMessage{},
			err:     ErrInvalidRecords,
		},
		{
			request: SNSMessage{
				Records: []Record{
					{
						SNS: SNS{
							Message:   "hello",
							MessageID: "1",
							Subject:   "welcoming",
						}},
				},
			},
			err: nil,
		},
	}

	for _, test := range tests {
		err := Handler(test.request)
		assert.IsType(t, test.err, err)
	}
}
