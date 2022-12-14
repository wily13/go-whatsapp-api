package validations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wily13/go-whatsapp-api/structs"
)

func TestValidateSendMessage(t *testing.T) {
	type args struct {
		request structs.SendMessageRequest
	}
	tests := []struct {
		name string
		args args
		err  interface{}
	}{
		{
			name: "success phone & message normal",
			args: args{request: structs.SendMessageRequest{
				Phone:   "6289685024091",
				Message: "Hello this is testing",
			}},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err == nil {
				ValidateSendMessage(tt.args.request)
			} else {
				assert.PanicsWithValue(t, tt.err, func() {
					ValidateSendMessage(tt.args.request)
				})
			}

		})
	}
}
