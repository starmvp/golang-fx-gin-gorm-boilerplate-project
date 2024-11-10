package workflows

import (
	"context"
	_ "embed"
)

// not a real workflow, use as common tools wrapper
type CommonWorkflowAgent struct {
	WorkflowAgent
}

func NewCommonWorkflowAgent(opts ...Option) *CommonWorkflowAgent {
	name := "common-workflow"
	description := "Common tools"

	opts = append(opts,
		WithName(name),
		WithDescription(description),
	)

	return &CommonWorkflowAgent{
		WorkflowAgent: *NewWorkflowAgent(opts...),
	}
}

func (a CommonWorkflowAgent) InPrompt() bool {
	return false
}

func (a CommonWorkflowAgent) ToolsInPrompt() bool {
	return true
}

//go:embed prompts/common_workflow_prompt.txt
var commonWorkflowPrompt string

func (t CommonWorkflowAgent) Call(ctx context.Context, input string) (string, error) {
	// TODO: prompt
	return commonWorkflowPrompt, nil
}
