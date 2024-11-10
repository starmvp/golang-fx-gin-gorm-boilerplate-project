package conversation

import (
	"github.com/starmvp/langchaingo/agents"
)

type Conversation interface {
	Start(agents.Executor) error
	Execute(agents.Executor, string) (string, error)

	GetStringInputChannel() *chan string
	GetStringOutputChannel() *chan string
	GetByteInputChannel() *chan byte
	GetByteOutputChannel() *chan byte
}
