package utils

import (
	"fmt"
	"strings"

	"github.com/starmvp/langchaingo/prompts"
	"github.com/starmvp/langchaingo/tools"
)

func CreateConversationalPrompt(lt []tools.Tool, prefix, instructions, suffix string) prompts.PromptTemplate {
	template := strings.Join([]string{prefix, instructions, suffix}, "\n\n")

	result := prompts.PromptTemplate{
		Template:       template,
		TemplateFormat: prompts.TemplateFormatGoTemplate,
		// InputVariables: []string{},
		InputVariables: []string{"input", "agent_scratchpad"},
		PartialVariables: map[string]any{
			"tool_names":        toolNames(lt),
			"tool_descriptions": toolDescriptions(lt),
		},
	}
	return result
}

func toolNames(lt []tools.Tool) string {
	var tn strings.Builder
	for i, tool := range lt {
		if i > 0 {
			tn.WriteString(", ")
		}
		tn.WriteString(tool.Name())
	}

	result := tn.String()
	return result
}

func toolDescriptions(lt []tools.Tool) string {
	var ts strings.Builder
	for _, tool := range lt {
		ts.WriteString(fmt.Sprintf("- %s: %s\n", tool.Name(), tool.Description()))
	}

	result := ts.String()
	return result
}
