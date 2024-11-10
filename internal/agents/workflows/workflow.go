package workflows

import (
	"errors"
	"fmt"
	"getidex_api/internal/agents/tools"
	"getidex_api/internal/utils"

	"github.com/starmvp/langchaingo/schema"
	Ltools "github.com/starmvp/langchaingo/tools"
)

var (
	ErrNilMemory = errors.New("nil memory")
)

type Workflow interface {
	Ltools.Tool

	GetTools() []Ltools.Tool
	InPrompt() bool
	ToolsInPrompt() bool
}

// WorkflowAgent is an tool set that can be used as a Tool also
//
//	As base agent, the Tool interface is not implemented, like an abstract class
type WorkflowAgent struct {
	tools.Tool

	Tools        []Ltools.Tool
	Memory       *schema.Memory
	PromptInputs map[string]any
	Prompt       string
}

func NewWorkflowAgent(opts ...Option) *WorkflowAgent {
	options := Options{}
	for _, opt := range opts {
		opt(&options)
	}

	if options.Memory == nil {
		panic(ErrNilMemory)
	}

	// tools.Options of the workflow, not for its Tools
	to := []tools.Option{
		tools.WithName(options.ToolName),
		tools.WithDescription(options.ToolDescription),
	}
	for _, h := range options.Options.CallbacksHandler {
		to = append(to, tools.WithCallbacksHandler(h))
	}
	to = append(to,
		tools.WithChain(options.Chain),
		tools.WithBuilder(options.Builder),
		tools.WithIO(options.Options.IO),
	)

	t := tools.NewTool(to...)

	tmpl := utils.CreateConversationalPrompt(
		options.Tools,
		options.PromptPrefix,
		options.PromptFormatInstructions,
		options.PromptSuffix,
	)
	prompt, err := tmpl.FormatPrompt(options.PromptInputs)
	if err != nil {
		fmt.Println("NewQueryEventWorkflowAgent: error:", err)
	}

	wa := WorkflowAgent{
		Tool:   *t,
		Tools:  options.Tools,
		Memory: options.Memory,
		Prompt: prompt.String(),
	}

	// fmt.Println("NewWorkflowAgent:", wa.Name(), ": Prompt=", wa.Prompt)

	return &wa
}

func (w WorkflowAgent) GetTools() []Ltools.Tool {
	return w.Tools
}

func (w WorkflowAgent) InPrompt() bool {
	// default: show in prompt
	return true
}

func (w WorkflowAgent) ToolsInPrompt() bool {
	// default: no tools in prompt
	return false
}
