package videos

import (
	"common/utils"
	"crypto/md5"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/godruoyi/go-snowflake"
)

var storage string

func Init(dir string) error {
	if storage != "" {
		return fmt.Errorf("storage already initialized")
	}
	full, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(full, 0755); err != nil {
		return err
	}

	if err := ensureFfmpeg(); err != nil {
		return err
	}

	storage = full
	return nil
}

func NewLocalVideo() (string, error) {
	name := strconv.FormatUint(snowflake.ID(), 16)
	// Hashed upload directories
	sum := md5.Sum([]byte(name))
	pathPart1 := strconv.FormatUint(uint64(sum[0]), 16)
	pathPart2 := strconv.FormatUint(uint64(sum[1]), 16)
	dir := path.Join(storage, pathPart1, pathPart2)
	if err := os.MkdirAll(dir, 0755); err != nil {
		hlog.Error("error saving video files", err)
		return "", utils.ErrorFilesystem
	}
	return path.Join(dir, name), nil
}
