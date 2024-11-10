package tools

import (
	"fmt"
	"getidex_api/internal/chain"
	"getidex_api/internal/utils"

	"github.com/starmvp/langchaingo/callbacks"
	"github.com/starmvp/langchaingo/chains"
	"github.com/starmvp/langchaingo/schema"
)

// As base Tool, the Call function of Tool not implement, like an abstract class

type Tool struct {
	ToolName        string
	ToolDescription string

	Chain   *chains.LLMChain
	Builder *chain.ChainBuilder
	Memory  *schema.Memory

	utils.IO

	CallbacksHandler callbacks.CombiningHandler
}

func NewTool(opts ...Option) *Tool {
	options := &Options{}
	for _, o := range opts {
		o(options)
	}

	if err := utils.ValidateIO(options.IO); err != nil {
		panic(fmt.Sprintf("Tool (%s): %v", options.ToolName, err))
	}

	handler := callbacks.CombiningHandler{}
	handler.Callbacks = append(handler.Callbacks, options.CallbacksHandler...)

	t := Tool{
		ToolName:         options.ToolName,
		ToolDescription:  options.ToolDescription,
		Chain:            options.Chain,
		Builder:          options.Builder,
		Memory:           options.Memory,
		IO:               options.IO,
		CallbacksHandler: handler,
	}

	return &t
}

func (p Tool) Name() string {
	return p.ToolName
}

func (p Tool) Description() string {
	return p.ToolDescription
}
