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
	queue       chan []byte
	key         []byte
	r           *io.PipeReader
	w           *io.PipeWriter
	limit       uint64
	buffer      []byte
	bytesbuffer *bytes.Buffer
}

// Close ...
func (thisptr *AESDEC) Close() error {
	close(thisptr.queue)
	return nil
}

// Read ...
func (thisptr *AESDEC) Read(p []byte) (n int, err error) {
	n, err = thisptr.bytesbuffer.Read(p)
	if err == nil || err != io.EOF {
		return
	}
	var size uint64
	err = binary.Read(thisptr.r, binary.LittleEndian, &size)
	if err != nil {
		return
	}
	if size > thisptr.limit || size < aes.BlockSize {
		err = errors.New("size error")
		return
	}
	for i := 0; i < int(size); {
		n, err := thisptr.r.Read(thisptr.buffer[i:size])
		if err != nil {
			return 0, err
		}
		i += n
	}
	block, err := aes.NewCipher(thisptr.key)
	if err != nil {
		return 0, err
	}
	iv := thisptr.buffer[:aes.BlockSize]
	stream := cipher.NewCFBDecrypter(block, iv)
	ciphertext := thisptr.buffer[aes.BlockSize:size]
	stream.XORKeyStream(ciphertext, ciphertext)
	thisptr.bytesbuffer.Write(ciphertext)

	return thisptr.bytesbuffer.Read(p)
}

// Write ...
func (thisptr *AESDEC) Write(p []byte) (n int, err error) {
	return thisptr.w.Write(p)
}
