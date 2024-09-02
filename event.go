package aquagram

type Event interface {
	GetText() string
	GetCallbackQuery() *CallbackQuery
	GetFrom() *User
	GetChat() *Chat
	GetEntities() []MessageEntity
}
