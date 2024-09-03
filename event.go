package aquagram

type Event interface {
	GetMessage() *Message
	GetCallbackQuery() *CallbackQuery
	GetFrom() *User
	GetChat() *Chat
	GetEntities() []MessageEntity
}
