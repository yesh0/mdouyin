package videos

import (
	"common/utils"
	"mime/multipart"

	"github.com/h2non/filetype"
)

// Quickly checks if the file is not a video
func CheckMagic(file *multipart.FileHeader) error {
	reader, err := file.Open()
	if err != nil {
		return utils.ErrorInternalError.Wrap(err)
	}

	size := file.Size
	if size > 1024*8 {
		size = 1024 * 8
	}
	buf := make([]byte, size)
	_, err = reader.Read(buf)
	if err != nil {
		return utils.ErrorInternalError.Wrap(err)
	}
	if err := reader.Close(); err != nil {
		return utils.ErrorInternalError.Wrap(err)
	}

	if !filetype.IsVideo(buf) {
		return utils.ErrorIncorrectFileType.With("expecting a video")
	}
	return nil
}
