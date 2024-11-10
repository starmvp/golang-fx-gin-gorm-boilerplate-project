package agents

import (
	"context"
	"errors"
	"fmt"
	"getidex_api/internal/chain"
	"getidex_api/internal/conversation"
	"getidex_api/internal/utils"

	"github.com/starmvp/langchaingo/agents"
	"github.com/starmvp/langchaingo/callbacks"
	"github.com/starmvp/langchaingo/chains"
	"github.com/starmvp/langchaingo/llms"
	"github.com/starmvp/langchaingo/schema"
	"github.com/starmvp/langchaingo/tools"
)

var (
	ErrNilChainBuilder = errors.New("nil chain builder")
	ErrNilContext      = errors.New("nil context")
	ErrNilLLM          = errors.New("nil llm")
)

type Agent struct {
	LLM   llms.Model
	Agent agents.Agent

	Ctx              context.Context
	Builder          *chain.ChainBuilder
	Chains           []chains.Chain
	Tools            []tools.Tool
	Memory           *schema.Memory
	CallbacksHandler callbacks.Handler
	Conversation     conversation.Conversation

	utils.IO
}

func NewAgent(opts ...Option) (*Agent, error) {
	options := Options{}
	for _, opt := range opts {
		opt(&options)
	}

	if options.Ctx == nil {
		return nil, ErrNilContext
	}

	if options.Builder == nil {
		return nil, ErrNilChainBuilder
	}

	c := options.Chains

	fmt.Println("NewAgent: options.Tools=", options.Tools)

	memory := options.Memory
	handler := callbacks.CombiningHandler{}
	handler.Callbacks = append(handler.Callbacks, options.CallbacksHandler...)

	io := utils.IO{
		StringInputChannel:  options.StringInputChannel,
		StringOutputChannel: options.StringOutputChannel,
		ByteInputChannel:    options.ByteInputChannel,
		ByteOutputChannel:   options.ByteOutputChannel,
	}
	if io.StringInputChannel == nil && options.Conversation.GetStringInputChannel() != nil {
		io.StringInputChannel = options.Conversation.GetStringInputChannel()
	}
	if io.StringOutputChannel == nil && options.Conversation.GetStringOutputChannel() != nil {
		io.StringOutputChannel = options.Conversation.GetStringOutputChannel()
	}
	if io.ByteInputChannel == nil && options.Conversation.GetByteInputChannel() != nil {
		io.ByteInputChannel = options.Conversation.GetByteInputChannel()
	}
	if io.ByteOutputChannel == nil && options.Conversation.GetByteOutputChannel() != nil {
		io.ByteOutputChannel = options.Conversation.GetByteOutputChannel()
	}

	if err := utils.ValidateIO(io); err != nil {
		return nil, err
	}

	a := &Agent{
		Agent:            options.LangChainAgent,
		Ctx:              options.Ctx,
		Builder:          options.Builder,
		Chains:           c,
		Tools:            options.Tools,
		Memory:           memory,
		CallbacksHandler: handler,
		Conversation:     options.Conversation,
		IO:               io,
	}

	return a, nil
}

func (a Agent) CreateExecutor() agents.Executor {
	executor := agents.NewExecutor(
		a.Agent,
		agents.WithMaxIterations(10),
		agents.WithMemory(*a.Memory),
		agents.WithCallbacksHandler(a.CallbacksHandler),
	)
	return *executor

}

func (a Agent) Plan(
	ctx context.Context,
	intermediateSteps []schema.AgentStep,
	inputs map[string]any,
	messages []llms.ChatMessage,
) (
	[]schema.AgentAction,
	*schema.AgentFinish,
	[]llms.ChatMessage,
	error,
) {
	fmt.Println("")
	fmt.Println("Agent Plan: inputs=", inputs)
	fmt.Println("")
	return a.Agent.Plan(ctx, intermediateSteps, inputs, nil)
}

func (a Agent) GetInputKeys() []string {
	fmt.Println("")
	fmt.Println("Agent GetInputKeys")
	fmt.Println("")
	return a.Agent.GetInputKeys()
}

func (a Agent) GetOutputKeys() []string {
	fmt.Println("")
	fmt.Println("Agent GetOutputKeys")
	fmt.Println("")
	return a.Agent.GetOutputKeys()
}

func (a Agent) GetTools() []tools.Tool {
	fmt.Println("")
	fmt.Println("Agent GetTools")
	fmt.Println("")
	return a.Agent.GetTools()
}
