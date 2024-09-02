package aquagram

type Updater interface {
	DispatchUpdate(update Update)
}
