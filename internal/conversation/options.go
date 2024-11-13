package conversation

import (
	"context"
	"getidex_api/internal/utils"
)

type options struct {
	Ctx context.Context
	IO  utils.IO
}

type Option func(*options)

func WithContext(ctx context.Context) Option {
	return func(o *options) {
		o.Ctx = ctx
	}
}

func WithIO(io utils.IO) Option {
	return func(o *options) {
		o.IO = io
	}
}
