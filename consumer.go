// pgque-go -- Go client for PgQue
// Copyright 2026 Nikolay Samokhvalov. Apache-2.0 license.

package pgque

import (
	"context"
	"time"
)

// HandlerFunc processes a single message. Return nil to indicate success.
type HandlerFunc func(ctx context.Context, msg Message) error

// Consumer polls a queue and dispatches messages to registered handlers.
type Consumer struct {
	client       *Client
	queue        string
	name         string
	pollInterval time.Duration
	handlers     map[string]HandlerFunc
}

// Handle registers a handler for the given event type.
func (c *Consumer) Handle(eventType string, fn HandlerFunc) {
	panic("not implemented")
}

// Start begins the poll loop, blocking until ctx is cancelled.
func (c *Consumer) Start(ctx context.Context) error {
	panic("not implemented")
}
