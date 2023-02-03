package videos

import "bytes"

type LimitedBuffer struct {
	buffer bytes.Buffer
	limit  int
}

func NewBuffer(limit int) *LimitedBuffer {
	return &LimitedBuffer{
		buffer: *bytes.NewBuffer(make([]byte, 0, limit)),
		limit:  limit,
	}
}

func (buf *LimitedBuffer) Write(p []byte) (n int, err error) {
	if free := buf.limit - buf.buffer.Len(); free > len(p) {
		buf.buffer.Write(p)
	} else {
		buf.buffer.Write(p[0:free])
	}
	return len(p), nil
}

func (buf *LimitedBuffer) String() string {
	return buf.buffer.String()
}

func (buf *LimitedBuffer) Len() int {
	return buf.buffer.Len()
}
