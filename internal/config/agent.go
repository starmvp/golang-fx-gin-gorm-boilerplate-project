package config

import "os"

type OpenAIConfig struct {
	BaseUrl            string `yaml:"base_url"`
	ApiKey             string `yaml:"api_key"`
	ApiVersion         string `yaml:"api_version"`
	ModelName          string `yaml:"model"`
	EmbeddingModelName string `yaml:"embedding_model"`
}

type AzureConfig struct {
	Endpoint           string `yaml:"endpoint"`
	ApiKey             string `yaml:"api_key"`
	ApiVersion         string `yaml:"api_version"`
	ModelName          string `yaml:"model"`
	EmbeddingModelName string `yaml:"embedding_model"`
}

type AgentConfig struct {
	LLMType string       `yaml:"llm_type"`
	OpenAI  OpenAIConfig `yaml:"openai"`
	Azure   AzureConfig  `yaml:"azure"`
}

func LoadAgentConfig() *AgentConfig {
	agentConfig := &AgentConfig{
		LLMType: os.Getenv("LLM_TYPE"),
		OpenAI: OpenAIConfig{
			BaseUrl:            os.Getenv("OPENAI_BASE_URL"),
			ApiKey:             os.Getenv("OPENAI_API_KEY"),
			ApiVersion:         os.Getenv("OPENAI_API_VERSION"),
			ModelName:          os.Getenv("OPENAI_MODEL"),
			EmbeddingModelName: os.Getenv("OPENAI_EMBEDDING_MODEL"),
		},
		Azure: AzureConfig{
			Endpoint:           os.Getenv("AZURE_ENDPOINT"),
			ApiKey:             os.Getenv("AZURE_API_KEY"),
			ApiVersion:         os.Getenv("AZURE_API_VERSION"),
			ModelName:          os.Getenv("AZURE_MODEL"),
			EmbeddingModelName: os.Getenv("AZURE_EMBEDDING_MODEL"),
		},
	}
	return agentConfig
}
