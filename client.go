package driver

import (
	"github.com/go-disk/driver/pb"
	"google.golang.org/grpc"
)

// Client for communication with the database Disk.
type Client struct {
	disk pb.DiskClient
}

// New create new instance Client.
func New(c grpc.ClientConnInterface) *Client {
	return &Client{
		disk: pb.NewDiskClient(c),
	}
}
