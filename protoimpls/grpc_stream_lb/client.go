package grpc_stream_lb

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/tigrannajaryan/exp-otelproto/core"
	"github.com/tigrannajaryan/exp-otelproto/encodings/traceprotobuf"
)

// Client can connect to a server and send a batch of spans.
type Client struct {
	client                  traceprotobuf.StreamTracerClient
	stream                  traceprotobuf.StreamTracer_ExportClient
	lastStreamOpen          time.Time
	ReopenAfterEveryRequest bool
}

// How often to reopen the stream to help LB's rebalance traffic.
var streamReopenPeriod = 30 * time.Second

func (c *Client) Connect(server string) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c.client = traceprotobuf.NewStreamTracerClient(conn)

	// Establish stream to server.
	return c.openStream()
}

func (c *Client) openStream() error {
	var err error
	c.stream, err = c.client.Export(context.Background())
	if err != nil {
		return err
	}
	c.lastStreamOpen = time.Now()
	return nil
}

func (c *Client) Export(batch core.ExportRequest) {
	// Send the batch via stream.
	c.stream.Send(batch.(*traceprotobuf.ExportRequest))

	// Wait for response from server. This is full synchronous operation,
	// we do not send batches concurrently.
	_, err := c.stream.Recv()

	if err != nil {
		log.Fatal("Error from server when expecting batch response")
	}

	if c.ReopenAfterEveryRequest || time.Since(c.lastStreamOpen) > streamReopenPeriod {
		// Close and reopen the stream.
		c.lastStreamOpen = time.Now()
		err = c.stream.CloseSend()
		if err != nil {
			log.Fatal("Error closing stream")
		}
		if err = c.openStream(); err != nil {
			log.Fatal("Error opening stream")
		}
	}
}
