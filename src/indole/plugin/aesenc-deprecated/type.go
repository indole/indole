package aesenc

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"io"
)

// AESENC ...
type AESENC struct {
	queue  chan []byte
	key    []byte
	buffer *bytes.Buffer
}

// Close ...
func (thisptr *AESENC) Close() error {
	close(thisptr.queue)
	return nil
}

// Read ...
func (thisptr *AESENC) Read(p []byte) (n int, err error) {
	n, err = thisptr.buffer.Read(p)
	if err == nil || err != io.EOF {
		return
	}

	data := <-thisptr.queue
	ciphertext, err := thisptr.enc(data)
	if err != nil {
		return 0, err
	}
	size := make([]byte, 8)
	binary.LittleEndian.PutUint64(size, uint64(len(ciphertext)))
	ciphersize, err := thisptr.enc(size)
	if err != nil {
		return 0, err
	}
	thisptr.buffer.Write(ciphersize)
	thisptr.buffer.Write(ciphertext)
	return thisptr.buffer.Read(p)
}

// Write ...
func (thisptr *AESENC) Write(p []byte) (n int, err error) {
	n = len(p)
	data := make([]byte, n)
	copy(data, p)
	thisptr.queue <- data
	return
}

func (thisptr *AESENC) enc(data []byte) ([]byte, error) {
	n := len(data)
	block, err := aes.NewCipher(thisptr.key)
	if err != nil {
		return nil, err
	}
	size := aes.BlockSize + n
	ciphertext := make([]byte, size)
	iv := ciphertext[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data[0:n])
	return ciphertext, nil
}
