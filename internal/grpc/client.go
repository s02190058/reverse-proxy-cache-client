package grpc

import (
	thumbnailpb "github.com/s02190058/reverse-proxy-cache/gen/go/thumbnail/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewClient creates a gRPC client.
func NewClient(host, port string) (thumbnailpb.ThumbnailServiceClient, func() error, error) {
	conn, err := grpc.Dial(
		addr(host, port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, nil, err
	}

	return thumbnailpb.NewThumbnailServiceClient(conn), conn.Close, nil
}

// addr returns a host:port gRPC address.
func addr(host, port string) string {
	return host + ":" + port
}
