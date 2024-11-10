package conversation

type StreamingConversation struct {
	InputChanel  chan byte
	OutputChanel chan byte
}

func NewStreamingConversation() *StreamingConversation {
	ic := make(chan byte, 1024*64)
	oc := make(chan byte, 1024*64)
	return &StreamingConversation{InputChanel: ic, OutputChanel: oc}
}
