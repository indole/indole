package aesdecodepacket

import (
	"crypto/aes"
	"crypto/cipher"
)

// AESDecodePacket ...
type AESDecodePacket struct {
	queue chan []byte
	key   []byte
}

// Close ...
func (thisptr *AESDecodePacket) Close() error {
	close(thisptr.queue)
	return nil
}

// Read ...
func (thisptr *AESDecodePacket) Read(p []byte) (n int, err error) {
	return copy(p, <-thisptr.queue), nil
}

// Write ...
func (thisptr *AESDecodePacket) Write(p []byte) (n int, err error) {
	data, err := thisptr.decode(p)
	if err != nil {
		return 0, err
	}
	thisptr.queue <- data
	return len(p), nil
}

func (thisptr *AESDecodePacket) decode(data []byte) ([]byte, error) {
	block, err := aes.NewCipher(thisptr.key)
	if err != nil {
		return nil, err
	}
	iv := data[:aes.BlockSize]
	stream := cipher.NewCFBDecrypter(block, iv)
	ciphertext := data[aes.BlockSize:]
	text := make([]byte, len(ciphertext))
	stream.XORKeyStream(text, ciphertext)
	return text, nil
}
