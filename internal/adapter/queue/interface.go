package queue

type QueueClient interface {
	PullMessages() error
}
