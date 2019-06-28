package streamtopacketwithaes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"indole/utils"
	"io"
)

// StreamToPacketWithAES ...
type StreamToPacketWithAES struct {
	queue  chan []byte
	reader *io.PipeReader
	writer *io.PipeWriter
	key    []byte
}

// Close ...
func (thisptr *StreamToPacketWithAES) Close() error {
	close(thisptr.queue)
	return utils.FirstError(thisptr.reader.Close(), thisptr.writer.Close())
}

// Read ...
func (thisptr *StreamToPacketWithAES) Read(p []byte) (n int, err error) {
	return copy(p, <-thisptr.queue), nil
}

// Write ...
func (thisptr *StreamToPacketWithAES) Write(p []byte) (n int, err error) {
	return thisptr.writer.Write(p)
}

func (thisptr *StreamToPacketWithAES) transcode() {
	defer func() {
		thisptr.Close()
	}()
	for {
		ciphersize := make([]byte, headersize)
		if err := thisptr.readbuf(ciphersize); err != nil {
			return
		}
		textsize, err := thisptr.decode(ciphersize)
		if err != nil {
			return
		}
		size := int(binary.LittleEndian.Uint64(textsize))
		buffer := make([]byte, size)
		if err := thisptr.readbuf(buffer); err != nil {
			return
		}
		thisptr.queue <- buffer
	}
}

func (thisptr *StreamToPacketWithAES) readbuf(p []byte) error {
	for i := 0; i < len(p); {
		n, err := thisptr.reader.Read(p[i:])
		if err != nil {
			return err
		}
		i += n
	}
	return nil
}

func (thisptr *StreamToPacketWithAES) decode(data []byte) ([]byte, error) {
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

const headersize = aes.BlockSize + 8
