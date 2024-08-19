package ticket_retrieve

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestExtract(t *testing.T) {
	type args struct {
		app   *fiber.App
		route string
	}
	type expected struct {
		err        error
		id         int
		statudCode int
	}
	tests := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "not found case",
			args: args{
				app:   fiber.New(),
				route: "/tickets/10",
			},
			expected: expected{
				err:        nil,
				id:         10,
				statudCode: http.StatusNotFound,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.args.route, nil)

			resp, _ := tt.args.app.Test(req, 1)
			assert.Equal(t, tt.expected.statudCode, resp.StatusCode)
		})
	}
}
