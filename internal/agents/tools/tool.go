package tools

import (
	"fmt"

	"boilerplate/internal/chain"
	"boilerplate/internal/utils"

	"github.com/starmvp/langchaingo/callbacks"
	"github.com/starmvp/langchaingo/chains"
	"github.com/starmvp/langchaingo/schema"
	"github.com/starmvp/langchaingo/vectorstores"
)

// As base Tool, the Call function of Tool not implement, like an abstract class

type Tool struct {
	ToolName        string
	ToolDescription string

	// for langchain
	Chain            *chains.LLMChain
	Builder          *chain.ChainBuilder
	Memory           schema.Memory
	CallbacksHandler callbacks.CombiningHandler

	// for logging
	Prompt                string
	PromptTemplate        string
	Input                 string
	Output                string
	UsagePromptTokens     int
	UsageCompletionTokens int

	// vector store
	VectorStore           vectorstores.VectorStore
	RetrieverNumDocuments int
	Retriever             schema.Retriever

	// to be used as a long-running agent
	utils.IO
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
	handler.Callbacks = append(handler.Callbacks, options.CallbacksHandlers...)

	t := Tool{
		ToolName:              options.ToolName,
		ToolDescription:       options.ToolDescription,
		Chain:                 options.Chain,
		Builder:               options.Builder,
		Memory:                options.Memory,
		CallbacksHandler:      handler,
		VectorStore:           options.VectorStore,
		RetrieverNumDocuments: options.RetrieverNumDocuments,
		Retriever:             options.Retriever,
		IO:                    options.IO,
	}

	return &t
}

func (p Tool) Name() string {
	return p.ToolName
}

func (p Tool) Description() string {
	return p.ToolDescription
}
