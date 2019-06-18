package aesdec

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"errors"
	"io"
)

// AESDEC ...
type AESDEC struct {
	queue  chan []byte
	key    []byte
	r      *io.PipeReader
	w      *io.PipeWriter
	buf    []byte
	buffer *bytes.Buffer
}

// Close ...
func (thisptr *AESDEC) Close() error {
	thisptr.r.Close()
	thisptr.w.Close()
	close(thisptr.queue)
	return nil
}

// Read ...
func (thisptr *AESDEC) Read(p []byte) (n int, err error) {
	n, err = thisptr.buffer.Read(p)
	if err == nil || err != io.EOF {
		return
	}
	sizetext, err := thisptr.readbufdec(sizesize)
	if err != nil {
		return
	}
	size := int(binary.LittleEndian.Uint64(sizetext))
	if size > len(thisptr.buf) || size < aes.BlockSize {
		err = errors.New("size error")
		return
	}
	text, err := thisptr.readbufdec(size)
	if err != nil {
		return
	}
	thisptr.buffer.Write(text)
	return thisptr.buffer.Read(p)
}

// Write ...
func (thisptr *AESDEC) Write(p []byte) (n int, err error) {
	return thisptr.w.Write(p)
}

func (thisptr *AESDEC) dec(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(thisptr.key)
	if err != nil {
		return nil, err
	}
	iv := data[:aes.BlockSize]
	stream := cipher.NewCFBDecrypter(block, iv)
	ciphertext := data[aes.BlockSize:]
	stream.XORKeyStream(ciphertext, ciphertext)
	return ciphertext, nil
}

func (thisptr *AESDEC) readbuf(size int) error {
	for i := 0; i < size; {
		n, err := thisptr.r.Read(thisptr.buf[i:size])
		if err != nil {
			return err
		}
		i += n
	}
	return nil
}

func (thisptr *AESDEC) readbufdec(size int) ([]byte, error) {
	err := thisptr.readbuf(size)
	if err != nil {
		return nil, err
	}
	return thisptr.dec(thisptr.buf[:size])
}

const (
	sizesize = 24
)
