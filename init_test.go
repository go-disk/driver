package driver_test

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"os"
	"testing"

	"github.com/go-disk/driver"
	"github.com/go-disk/driver/pb"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

var (
	ctx = context.Background()

	fileMeta = []byte{1, 1, 1, 1, 1, 1}
	bufFile  = readFile()
	fileID   = id()
)

func start(t *testing.T) (*driver.FileSystemClient, assertion) {
	r := assertion(require.New(t))

	srv := grpc.NewServer()
	pb.RegisterDiskServer(srv, serverMock{r})
	ln, err := net.Listen("tcp", "")
	r.Nil(err)
	go func() { r.Nil(srv.Serve(ln)) }()
	c, err := grpc.Dial(ln.Addr().String(), grpc.WithInsecure())
	r.Nil(err)

	t.Cleanup(func() {
		srv.Stop()
		r.Nil(c.Close())
	})

	return driver.New(c), r
}

var _ pb.DiskServer = &serverMock{}

type assertion interface {
	Nil(obj interface{}, msg ...interface{})
	NotNil(obj interface{}, msg ...interface{})
	Equal(expected, actual interface{}, msgAndArgs ...interface{})
}

type serverMock struct {
	assert assertion
}

func (t serverMock) UploadFile(stream pb.Disk_UploadFileServer) error {
	msg, err := stream.Recv()
	t.assert.Nil(err)

	info := msg.GetInfo()
	t.assert.NotNil(info)

	t.assert.Equal(fileMeta, info.Meta)

	res, err := ioutil.ReadAll(driver.NewGRPCReader(stream))
	t.assert.Nil(err)

	t.assert.Equal(bufFile, res)

	return stream.SendAndClose(&pb.UUID{Value: fileID.String()})
}

func (t serverMock) DeleteFile(ctx context.Context, file *pb.UUID) (*empty.Empty, error) {
	t.assert.Equal(fileID.String(), file.Value)

	return &empty.Empty{}, nil
}

func id() uuid.UUID {
	return uuid.Must(uuid.NewV4())
}

const (
	testFile = `testdata/file.txt`
)

func readFile() []byte {
	f, err := os.Open(testFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	return buf
}
