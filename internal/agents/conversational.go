package agents

import (
	_ "embed"
	"fmt"

	"boilerplate/internal/agents/workflows"
	"boilerplate/internal/utils"

	"github.com/starmvp/langchaingo/agents"
	"github.com/starmvp/langchaingo/callbacks"
	"github.com/starmvp/langchaingo/chains"
	"github.com/starmvp/langchaingo/prompts"
	"github.com/starmvp/langchaingo/tools"
)

type ConversationalWorkflowAgent struct {
	Agent

	Workflows []workflows.Workflow
}

func NewConversationalWorkflowAgent(
	opts ...Option,
) (*ConversationalWorkflowAgent, error) {
	options := DefaultRootConversationalAgentOptions
	for _, opt := range opts {
		opt(&options)
	}

	laOptions := agents.Options{}
	for _, opt := range options.LangChainOptions {
		opt(&laOptions)
	}

	handler := callbacks.CombiningHandler{}
	handler.Callbacks = append(handler.Callbacks, options.CallbacksHandlers...)

	prompt := getConversationalPrompt(
		options,
		options.Workflows,
		options.Tools,
	)
	fmt.Println("NewConversationalAgent: prompt=", prompt)

	llm := options.LLM
	if llm == nil {
		return nil, ErrNilLLM
	}

	lca := agents.ConversationalAgent{
		Chain: chains.NewLLMChain(
			llm,
			prompt,
			chains.WithCallback(handler),
		),
		Tools:            options.Tools,
		OutputKey:        options.OutputKey,
		CallbacksHandler: handler, // DO NOT use callbacks to avoid streaming
		UseStreamingMode: false,
	}

	opts = append(opts, WithLangChainAgent(lca))
	a, err := NewAgent(opts...)
	if err != nil {
		return nil, err
	}

	return &ConversationalWorkflowAgent{
		Agent:     *a,
		Workflows: options.Workflows,
	}, nil
}

//go:embed prompts/conversational_prefix.txt
var _defaultConversationalPrefix string //nolint:gochecknoglobals

//go:embed prompts/conversational_format_instructions.txt
var _defaultConversationalFormatInstructions string //nolint:gochecknoglobals

//go:embed prompts/conversational_suffix.txt
var _defaultConversationalSuffix string //nolint:gochecknoglobals

var DefaultRootConversationalAgentOptions = Options{
	PromptPrefix:             _defaultConversationalPrefix,
	PromptFormatInstructions: _defaultConversationalFormatInstructions,
	PromptSuffix:             _defaultConversationalSuffix,
}

func getConversationalPrompt(
	options Options,
	workflows []workflows.Workflow,
	t []tools.Tool,
) prompts.PromptTemplate {
	if options.PromptTemplate.Template != "" {
		return options.PromptTemplate
	}

	tl := make([]tools.Tool, 0)
	tl = append(tl, t...)
	for _, w := range workflows {
		if w.InPrompt() {
			tl = append(tl, w)
		}
	}
	for _, w := range workflows {
		ts := w.ToolsInPrompt()
		tl = append(tl, ts...)
	}

	fmt.Println("getConversationalPrompt: options.PromptPrefix=", options.PromptPrefix)

	prompt := utils.CreateConversationalPrompt(
		tl,
		options.PromptPrefix,
		options.PromptFormatInstructions,
		options.PromptSuffix,
	)

	fmt.Printf("getConversationalPrompt: prompt=%v\n", prompt)

	return prompt
}
