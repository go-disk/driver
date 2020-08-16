package driver_test

import (
	"os"
	"testing"
)

func TestClient_MkDir(t *testing.T) {
	t.Parallel()

	c, assert := start(t)

	res, err := c.MkDir(ctx, dirPath)
	assert.Nil(err)

	assert.Equal(dirID, res)
}

func TestClient_RmFile(t *testing.T) {
	t.Parallel()

	c, assert := start(t)

	err := c.RmFile(ctx, filePath)
	assert.Nil(err)
}

func TestClient_RmDir(t *testing.T) {
	t.Parallel()

	c, assert := start(t)

	err := c.RmDir(ctx, dirPath)
	assert.Nil(err)
}

func TestClient_UploadFile(t *testing.T) {
	t.Parallel()

	c, assert := start(t)
	f, err := os.Open(testFile)
	assert.Nil(err)
	t.Cleanup(func() {
		assert.Nil(f.Close())
	})

	res, err := c.UploadFile(ctx, filePath, fileMeta, f)
	assert.Nil(err)

	assert.Equal(fileID, res)
}
