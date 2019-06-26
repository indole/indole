package aesencodepacket

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// AESEncodePacket ...
type AESEncodePacket struct {
	queue chan []byte
	key   []byte
}

// Close ...
func (thisptr *AESEncodePacket) Close() error {
	close(thisptr.queue)
	return nil
}

// Read ...
func (thisptr *AESEncodePacket) Read(p []byte) (n int, err error) {
	return copy(p, <-thisptr.queue), nil
}

// Write ...
func (thisptr *AESEncodePacket) Write(p []byte) (n int, err error) {
	data, err := thisptr.encode(p)
	if err != nil {
		return 0, err
	}
	thisptr.queue <- data
	return len(p), nil
}

func (thisptr *AESEncodePacket) encode(data []byte) ([]byte, error) {
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
