package chain

import (
	"context"
	"fmt"
	"log"

	"boilerplate/internal/config"

	"github.com/starmvp/langchaingo/chains"
	"github.com/starmvp/langchaingo/llms"
	"github.com/starmvp/langchaingo/llms/openai"
	"github.com/starmvp/langchaingo/prompts"
)

type ChainBuilder struct {
	LLMClient llms.Model
}

func NewChainBuilder(c *config.AgentConfig) (*ChainBuilder, error) {

	fmt.Printf("NewChainBuilder: c: %v\n", c)

	var client llms.Model
	var err error

	llmType := c.LLMType
	fmt.Println("LLM Type: ", llmType)
	switch llmType {
	case "openai":
		apiKey := c.OpenAI.ApiKey
		if apiKey == "" {
			log.Fatalf("missing OpenAI API key")
		}
		modelName := c.OpenAI.ModelName
		fmt.Println("Model Name: ", modelName)
		if modelName == "" {
			log.Fatalf("missing OpenAI model name")
		}
		embeddingModelName := c.OpenAI.EmbeddingModelName
		fmt.Println("Embedding Model Name: ", embeddingModelName)
		if embeddingModelName == "" {
			log.Fatalf("missing OpenAI embedding model name")
		}
		client, err = openai.New(
			openai.WithAPIType(openai.APITypeOpenAI),
			openai.WithToken(apiKey),
			openai.WithModel(modelName),
			openai.WithEmbeddingModel(embeddingModelName),
		)
		if err != nil {
			log.Fatalf("failed to create OpenAI client: %v", err)
			return nil, err
		}
		// case "azure":
		// 	endpoint := c.AzureEndpoint
		// 	if endpoint == "" {
		// 		log.Fatalf("missing Azure endpoint")
		// 	}
		// 	apiKey := c.AzureApiKey
		// 	if apiKey == "" {
		// 		log.Fatalf("missing Azure API key")
		// 	}
		// 	client, err = azure.New(endpoint, apiKey)
		// 	if err != nil {
		// 		log.Fatalf("failed to create OpenAI client: %v", err)
		// 		return nil, err
		// 	}
	}
	return &ChainBuilder{LLMClient: client}, nil
}

func (b *ChainBuilder) BuildBasicChain(pt prompts.PromptTemplate, opts ...chains.ChainCallOption) *chains.LLMChain {
	llmChain := chains.NewLLMChain(b.LLMClient, pt, opts...)
	return llmChain
}

func (b *ChainBuilder) BuildSequentialChain(chainNodes []chains.Chain, inputKeys []string, outputKeys []string, opts ...chains.SequentialChainOption) (*chains.SequentialChain, error) {
	c, err := chains.NewSequentialChain(chainNodes, inputKeys, outputKeys, opts...)
	if err != nil {
		log.Fatalf("failed to create chain: %v", err)
	}
	return c, err
}

func (b *ChainBuilder) RunChain(ctx context.Context, chain chains.Chain, input map[string]interface{}) (map[string]interface{}, error) {
	output, err := chains.Call(ctx, chain, input)
	if err != nil {
		return nil, err
	}
	return output, nil
}
