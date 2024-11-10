package conversation

import (
	"context"
	"fmt"

	"getidex_api/internal/utils"

	"github.com/starmvp/langchaingo/agents"
	"github.com/starmvp/langchaingo/chains"
	"github.com/starmvp/langchaingo/memory"
)

type StandardConversation struct {
	Ctx context.Context

	Buffer memory.ConversationBuffer

	utils.IO
}

func NewStandardConversation() *StandardConversation {
	// TODO: input/output channel options
	ctx := context.Background()
	sic := make(chan string, 5)
	soc := make(chan string, 5)

	io := utils.IO{
		StringInputChannel:  &sic,
		StringOutputChannel: &soc,
	}

	return &StandardConversation{
		Ctx: ctx,
		IO:  io,
	}
}

func (c *StandardConversation) Start(executor agents.Executor) error {
	go func() {
		for {
			input := <-*c.StringInputChannel
			fmt.Println("conv: input=", input)
			fmt.Printf("conv: c.Executor=%+v\n", executor)
			result, err := chains.Run(c.Ctx, &executor, input)
			if err != nil {
				fmt.Println("conv: error:", err)
			} else {
				fmt.Println("conv: result=", result)
			}
			*c.StringOutputChannel <- result
		}
	}()

	return nil
}

func (c *StandardConversation) Execute(executor agents.Executor, input string) (string, error) {
	fmt.Println("conv: input=", input)
	result, err := chains.Run(c.Ctx, &executor, input)
	if err != nil {
		fmt.Println("conv: error:", err)
		return "", err
	}

	fmt.Println("conv: result=", result)
	return result, nil
}

func (c *StandardConversation) GetStringInputChannel() *chan string {
	return c.StringInputChannel
}

func (c *StandardConversation) GetStringOutputChannel() *chan string {
	return c.StringOutputChannel
}

func (c *StandardConversation) GetByteInputChannel() *chan byte {
	return c.ByteInputChannel
}

func (c *StandardConversation) GetByteOutputChannel() *chan byte {
	return c.ByteOutputChannel
}
