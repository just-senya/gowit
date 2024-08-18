package ticket_create

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
)

type errHandler func(err error) bool
type bodySetter func(ctx *fiber.Ctx, b []byte)

func TestExtract(t *testing.T) {
	type args struct {
		app  *fiber.App
		body []byte
	}
	type expected struct {
		err        error
		name       string
		desc       string
		allocation uint32
	}
	tests := []struct {
		name       string
		args       args
		expected   expected
		errHandler errHandler
	}{
		// {
		// 	name: "empty req-body",
		// 	args: args{app: fiber.New()},
		// 	expected: expected{
		// 		err: errEmptyReqBody,
		// 	},
		// },
		// {
		// 	name: "invalid req-body",
		// 	args: args{
		// 		app:  fiber.New(),
		// 		body: []byte(`{"invalid-json"}`),
		// 	},
		// 	expected: expected{
		// 		err: nil,
		// 	},
		// },
		{
			name: "full-body",
			args: args{
				app:  fiber.New(),
				body: []byte(`{"name":"mock-name","desc":"mock-desc","allocation":123456}`),
			},
			expected: expected{
				err:        nil,
				name:       "mock-name",
				desc:       "mock-desc",
				allocation: 123456,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := tt.args.app.AcquireCtx(&fasthttp.RequestCtx{})
			ctx.Context().Request.SetBody(tt.args.body)
			extr := NewCreateExtractor()

			res, err := extr.Extract(ctx)

			assert.Equal(t, tt.expected.err, err)
			if tt.expected.err == nil {
				assert.Equal(t, tt.expected.name, res.GetName())
				assert.Equal(t, tt.expected.desc, res.GetDesc())
				assert.Equal(t, tt.expected.allocation, res.GetAllocation())
			}
		})
	}
}
