package videos_test

import (
	"bytes"
	"gateway/internal/videos"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const gnu_video = "https://static.fsf.org/nosvn/videos/escape-to-freedom/videos/escape-to-freedom-360p.mp4"
const filename = "escape-to-freedom-360p.mp4"

var file string

func TestMain(m *testing.M) {
	rand.Seed(time.Now().Unix())

	dir := os.TempDir()
	file = path.Join(dir, filename)
	if stat, err := os.Stat(file); err != nil {
		res, err := http.Get(gnu_video)
		if err != nil {
			log.Fatalln(err)
		}
		content, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}
		if err := os.WriteFile(file, content, 0644); err != nil {
			log.Fatalln(err)
		}
	} else if stat.IsDir() {
		log.Fatalf("%s is a directory", filename)
	}
	m.Run()
}

func TestCheckFast(t *testing.T) {
	buf := bytes.Buffer{}
	w := multipart.NewWriter(&buf)
	filer, _ := w.CreateFormFile("data", file)
	content, _ := os.ReadFile(file)
	filer.Write(content)
	w.Close()

	r := multipart.NewReader(&buf, w.Boundary())
	form, _ := r.ReadForm(int64(buf.Cap()))
	assert.Nil(t, videos.CheckMagic(form.File["data"][0]))
}

func TestLimited(t *testing.T) {
	buf := videos.NewBuffer(128)
	n, err := buf.Write(make([]byte, 64))
	assert.Equal(t, 64, n)
	assert.Nil(t, err)
	assert.Equal(t, 64, buf.Len())
	n, err = buf.Write(make([]byte, 128))
	assert.Equal(t, 128, n)
	assert.Nil(t, err)
	assert.Equal(t, 128, buf.Len())
}

func randName() string {
	return strconv.FormatUint(rand.Uint64(), 16)
}

func TestVideoSaving(t *testing.T) {
	dir := path.Join(os.TempDir(), randName())
	assert.Nil(t, videos.Init(dir))
	file, err := videos.NewLocalVideo()
	assert.Nil(t, err)
	assert.Contains(t, file, dir)

	assert.Nil(t, os.RemoveAll(dir))
}

func TestValidation(t *testing.T) {
	assert.Nil(t, videos.Init(os.TempDir()))
	assert.Nil(t, videos.ValidateVideo(file))

	invalid := path.Join(os.TempDir(), randName()+"invalid.mp4")
	assert.Nil(t, os.WriteFile(invalid, make([]byte, 1*1024*1024), 0644))
	assert.NotNil(t, videos.ValidateVideo(invalid))
	assert.Nil(t, os.Remove(invalid))
}

func TestCoverGeneration(t *testing.T) {
	assert.Nil(t, videos.Init(os.TempDir()))
	cover := path.Join(os.TempDir(), randName()+"cover.png")
	assert.Nil(t, videos.GenerateCover(file, cover))

	stat, err := os.Stat(cover)
	assert.Nil(t, err)
	assert.False(t, stat.IsDir())
	assert.Greater(t, stat.Size(), int64(10*1024))
	assert.Nil(t, os.Remove(cover))
}
