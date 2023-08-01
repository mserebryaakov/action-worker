package queue

import "action-worker/internal/action"

type QueueClient interface {
	PullMessage() (*action.Action, error)
}
