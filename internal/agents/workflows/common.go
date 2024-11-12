package workflows

import (
	"context"
	_ "embed"

	"boilerplate/internal/agents/tools"

	Ltools "github.com/starmvp/langchaingo/tools"
)

// not a real workflow, use as common tools wrapper
type CommonWorkflowAgent struct {
	WorkflowAgent
}

func NewCommonWorkflowAgent(opts ...Option) *CommonWorkflowAgent {
	name := "common-workflow"
	description := "Common tools"

	opts = append(opts,
		WithToolOptions(
			tools.WithName(name),
			tools.WithDescription(description),
		),
	)

	return &CommonWorkflowAgent{
		WorkflowAgent: *NewWorkflowAgent(opts...),
	}
}

func (a CommonWorkflowAgent) InPrompt() bool {
	return false
}

func (a CommonWorkflowAgent) ToolsInPrompt() []Ltools.Tool {
	return a.Tools
}

//go:embed prompts/common_workflow_prompt.txt
var commonWorkflowPrompt string

func (t CommonWorkflowAgent) Call(ctx context.Context, input string) (string, error) {
	// TODO: prompt
	return commonWorkflowPrompt, nil
}
