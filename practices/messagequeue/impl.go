package messagequeue

type MessageQueue struct {
	messages []string
}

func NewMessageQueue() *MessageQueue {
	return &MessageQueue{}
}

func (mq *MessageQueue) Send(message string) {
	mq.messages = append(mq.messages, message)
}

func (mq *MessageQueue) Receive() string {
	if len(mq.messages) == 0 {
		return ""
	}
	msg := mq.messages[0]
	mq.messages = mq.messages[1:]
	return msg
}
