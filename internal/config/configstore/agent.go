package configstore

import (
	"log"
)

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

func NewAgentConfig(store *ConfigStore) *AgentConfig {
	agentConfig := &AgentConfig{}
	err := agentConfig.Load(store)
	if err != nil {
		log.Fatalf("FATAL: failed to load agent config. err: %+v", err)
	}
	return agentConfig
}

func (ac AgentConfig) SectionName() string {
	return "agent"
}

func (ac *AgentConfig) Load(store *ConfigStore) error {
	result, err := LoadSection[AgentConfig](store, ac.SectionName())
	if err != nil {
		return err
	}
	*ac = *result
	store.SetSection(ac.SectionName(), ac)
	return nil
}
