package driver

import (
	"github.com/go-disk/driver/pb"
	"google.golang.org/grpc"
)

// FileSystemClient for communication with the database Disk.
type FileSystemClient struct {
	disk pb.DiskClient
}

// New create new instance FileSystemClient.
func New(c grpc.ClientConnInterface) *FileSystemClient {
	return &FileSystemClient{
		disk: pb.NewDiskClient(c),
	}
}
