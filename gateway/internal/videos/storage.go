package videos

import (
	"common/snowy"
	"common/utils"
	"crypto/md5"
	"fmt"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

var storage string
var base string

func Init(dir string, baseUrl string) error {
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

	if u, err := url.Parse(baseUrl); err != nil {
		return fmt.Errorf("expecting a valid url: %v", err)
	} else {
		baseUrl = u.String()
	}

	storage = full
	base = baseUrl
	return nil
}

func BaseUrl() string {
	return base
}

func Dir() string {
	return storage
}

func Storage(subpath string) string {
	return path.Join(storage, subpath)
}

func NewLocalVideo() (string, error) {
	name := strconv.FormatUint(uint64(snowy.ID()), 16)
	// Hashed upload directories
	sum := md5.Sum([]byte(name))
	pathPart1 := strconv.FormatUint(uint64(sum[0]), 16)
	pathPart2 := strconv.FormatUint(uint64(sum[1]), 16)
	if err := os.MkdirAll(path.Join(storage, pathPart1, pathPart2), 0755); err != nil {
		hlog.Error("error saving video files", err)
		return "", utils.ErrorFilesystem
	}
	return path.Join(pathPart1, pathPart2, name), nil
}
