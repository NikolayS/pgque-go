// pgque-go -- Go client for PgQue
// Copyright 2026 Nikolay Samokhvalov. Apache-2.0 license.

package pgque

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Client is the main PgQue client backed by a pgx connection pool.
type Client struct {
	pool *pgxpool.Pool
}

// Connect creates a new Client connected to the given DSN.
func Connect(ctx context.Context, dsn string) (*Client, error) {
	panic("not implemented")
}

// Close releases the connection pool.
func (c *Client) Close() {
	panic("not implemented")
}

// Pool returns the underlying pgxpool for direct SQL access.
func (c *Client) Pool() *pgxpool.Pool {
	panic("not implemented")
}

// Send publishes an event to the named queue and returns the event ID.
func (c *Client) Send(ctx context.Context, queue string, ev Event) (int64, error) {
	panic("not implemented")
}

// Receive fetches up to maxMessages from the next batch for the consumer.
func (c *Client) Receive(ctx context.Context, queue, consumer string, maxMessages int) ([]Message, error) {
	panic("not implemented")
}

// Ack acknowledges (finishes) a batch, advancing the consumer position.
func (c *Client) Ack(ctx context.Context, batchID int64) error {
	panic("not implemented")
}

// NewConsumer creates a Consumer for the given queue and consumer name.
func (c *Client) NewConsumer(queue, name string, opts ...Option) *Consumer {
	panic("not implemented")
}
