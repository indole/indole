package packettostreamwithaes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/binary"
	"indole/utils"
	"io"
)

// PacketToStreamWithAES ...
type PacketToStreamWithAES struct {
	reader *io.PipeReader
	writer *io.PipeWriter
	key    []byte
}

// Close ...
func (thisptr *PacketToStreamWithAES) Close() error {
	return utils.FirstError(thisptr.reader.Close(), thisptr.writer.Close())
}

// Read ...
func (thisptr *PacketToStreamWithAES) Read(p []byte) (n int, err error) {
	return thisptr.reader.Read(p)
}

// Write ...
func (thisptr *PacketToStreamWithAES) Write(p []byte) (n int, err error) {
	size := make([]byte, 8)
	binary.LittleEndian.PutUint64(size, uint64(len(p)))
	ciphersize, err := thisptr.encode(size)
	if err != nil {
		return 0, err
	}
	_, err = thisptr.writer.Write(ciphersize)
	if err != nil {
		return 0, err
	}
	return thisptr.writer.Write(p)
}

func (thisptr *PacketToStreamWithAES) encode(data []byte) ([]byte, error) {
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
