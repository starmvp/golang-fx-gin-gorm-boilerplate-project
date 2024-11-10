package utils

import "errors"

type IO struct {
	StringInputChannel  *chan string
	ByteInputChannel    *chan byte
	StringOutputChannel *chan string
	ByteOutputChannel   *chan byte
}

func ValidateIO(io IO) error {
	if io.StringInputChannel == nil && io.ByteInputChannel == nil {
		return errors.New("missing input channel")
	}
	if io.StringOutputChannel == nil && io.ByteOutputChannel == nil {
		return errors.New("missing output channel")
	}

	return nil
}

type IOOptions struct {
	IO
}

type IOOption func(*IOOptions)

func WithStringInputChannel(c *chan string) IOOption {
	return func(opts *IOOptions) {
		opts.StringInputChannel = c
	}
}

func WithStringOutputChannel(c *chan string) IOOption {
	return func(opts *IOOptions) {
		opts.StringOutputChannel = c
	}
}

func WithByteInputChannel(c *chan byte) IOOption {
	return func(opts *IOOptions) {
		opts.ByteInputChannel = c
	}
}

func WithByteOutputChannel(c *chan byte) IOOption {
	return func(opts *IOOptions) {
		opts.ByteOutputChannel = c
	}
}
