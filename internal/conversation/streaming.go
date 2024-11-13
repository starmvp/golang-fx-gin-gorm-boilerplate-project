package conversation

import (
	"context"

	"boilerplate/internal/utils"
)

type StreamingConversation struct {
	Ctx context.Context

	utils.IO
}

func NewStreamingConversation(opts ...Option) *StreamingConversation {
	options := options{}
	for _, opt := range opts {
		opt(&options)
	}

	ctx := options.Ctx
	if ctx == nil {
		ctx = context.Background()
	}

	io := options.IO
	if io.ByteInputChannel == nil {
		ic := make(chan byte, 1024*64)
		io.ByteInputChannel = &ic
	}
	if io.ByteOutputChannel == nil {
		oc := make(chan byte, 1024*64)
		io.ByteOutputChannel = &oc
	}
	return &StreamingConversation{
		Ctx: ctx,
		IO:  io,
	}
}
