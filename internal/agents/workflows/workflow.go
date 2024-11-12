package workflows

import (
	"errors"
	"fmt"

	"boilerplate/internal/agents/tools"
	"boilerplate/internal/utils"

	Ltools "github.com/starmvp/langchaingo/tools"
)

var (
	ErrNilMemory = errors.New("nil memory")
)

type Workflow interface {
	Ltools.Tool

	GetTools() []Ltools.Tool
	InPrompt() bool
	ToolsInPrompt() []Ltools.Tool
}

// WorkflowAgent is an tool set that can be used as a Tool also
//
//	As base agent, the Tool interface is not implemented, like an abstract class
type WorkflowAgent struct {
	tools.Tool

	Tools        []Ltools.Tool
	PromptInputs map[string]any
	Prompt       string
}

func NewWorkflowAgent(opts ...Option) *WorkflowAgent {
	options := Options{}
	for _, opt := range opts {
		opt(&options)
	}

	// tools.Options of the workflow, not for its Tools
	wato := GetToolOptions(options)
	// embedded tools.Tool in WorkflowAgent
	wat := tools.NewTool(wato...)

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
		Tool:   *wat,
		Tools:  options.Tools,
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

func (w WorkflowAgent) ToolsInPrompt() []Ltools.Tool {
	// default: no tools in prompt
	return []Ltools.Tool{}
}
