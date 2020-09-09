package driver

import (
	"errors"
	"fmt"
	"io"

	"github.com/go-disk/driver/pb"
)

var _ io.Reader = &grpcReader{}

type grpcReader struct {
	stream pb.Disk_UploadFileServer
	cache  []byte
}

const chunkSize = 64 * 1024

// NewGRPCReader create new grpc reader.
func NewGRPCReader(stream pb.Disk_UploadFileServer) io.Reader {
	return &grpcReader{
		stream: stream,
		cache:  []byte{},
	}
}

// Errors.
var (
	ErrMsgNotChunk = errors.New("message should be chunk")
)

// Read for implemented io.Reader.
func (r *grpcReader) Read(b []byte) (sent int, err error) {
	var res []byte

	if r.cacheIsEmpty() {
		msg, err := r.stream.Recv()
		if err != nil && !errors.Is(err, io.EOF) {
			return 0, fmt.Errorf("stream recv: %w", err)
		}

		if errors.Is(err, io.EOF) {
			return 0, io.EOF
		}

		chunk := msg.GetChunk()
		if chunk == nil {
			return 0, ErrMsgNotChunk
		}

		res = chunk.GetData()
	} else {
		res = r.pullCache()
	}

	if len(res) == 0 {
		return 0, io.EOF
	}

	sent = copy(b, res)
	if sent < len(res) {
		r.cache = res[sent:]
	}

	return sent, nil
}

func (r *grpcReader) cacheIsEmpty() bool {
	return len(r.cache) == 0
}

func (r *grpcReader) pullCache() []byte {
	res := r.cache
	r.cache = []byte{}

	return res
}
