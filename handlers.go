package driver

import (
	"context"
	"errors"
	"fmt"
	"io"

	"github.com/go-disk/driver/pb"
	"github.com/gofrs/uuid"
)

// UploadFile upload new file with path and meta.
func (c *FileSystemClient) UploadFile(ctx context.Context, path string, meta []byte, r io.Reader) (uuid.UUID, error) {
	fileInfo := &pb.UploadData{
		Data: &pb.UploadData_Info{
			Info: &pb.Data{
				Path: path,
				Meta: meta,
			},
		},
	}

	stream, err := c.disk.UploadFile(ctx)
	if err != nil {
		return uuid.Nil, fmt.Errorf("disk UploadFile by path: %s, err: %w", path, err)
	}

	err = stream.Send(fileInfo)
	if err != nil {
		return uuid.Nil, fmt.Errorf("send file info: %w", err)
	}

	chunk := make([]byte, chunkSize)

	for {
		n, err := r.Read(chunk)
		if err != nil && !errors.Is(err, io.EOF) {
			return uuid.Nil, fmt.Errorf("read file: %w", err)
		}

		if n == 0 {
			break
		}

		in := &pb.UploadData{
			Data: &pb.UploadData_Chunk{
				Chunk: &pb.NewChunk{
					Data: chunk[:n],
				},
			},
		}

		err = stream.Send(in)
		if err != nil {
			return uuid.Nil, fmt.Errorf("send chunk: %w", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		return uuid.Nil, fmt.Errorf("disk UploadFile: %w", err)
	}

	id, err := uuid.FromString(res.Value)
	if err != nil {
		return uuid.Nil, fmt.Errorf("parse id: %w", err)
	}

	return id, nil
}

// DeleteFile remove file by path.
func (c *FileSystemClient) DeleteFile(ctx context.Context, path string) error {
	in := &pb.DeleteFileInfo{Path: path}
	_, err := c.disk.DeleteFile(ctx, in)
	if err != nil {
		return fmt.Errorf("disk RmFile by path: %s, err: %w", path, err)
	}

	return nil
}
