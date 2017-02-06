package bencode

import (
	"bufio"
	"math"
	"strconv"
)

// ByteString - representation of bencode ByteString
type ByteString struct {
	bencodeBase
	Value []byte
}

// Load - loads byte string from bencode reader
func (str *ByteString) load(reader *bufio.Reader) error {
	lengthBytes, err := reader.ReadBytes(stringLengthValueDelimiter)
	if err != nil {
		return err
	}

	lengthString := string(lengthBytes[:len(lengthBytes)-1])
	length, err := strconv.ParseInt(lengthString, 10, 64)
	if err != nil {
		return err
	}

	result := make([]byte, 0, length)
	rest := length
	for rest > 0 {
		capacity := int(math.Min(1024, float64(rest)))
		buffer := make([]byte, capacity)
		readed, err := reader.Read(buffer)
		if err != nil {
			return err
		}
		result = append(result, buffer[0:readed]...)
		rest = rest - int64(readed)
	}

	str.BencodeData = make([]byte, 0, len(lengthBytes)+len(result))
	str.BencodeData = append(str.BencodeData, lengthBytes...)
	str.BencodeData = append(str.BencodeData, result...)
	str.Value = result

	return nil
}
