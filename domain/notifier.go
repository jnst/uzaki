package domain

type Notifier interface {
	Notify(success bool, title, text string)
}
