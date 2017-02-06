package bencode

import (
	"bufio"
	"strconv"
)

// Integer - representation of integer
type Integer struct {
	bencodeBase
	Value int64
}

// Load - load integer from reader
func (integer *Integer) load(reader *bufio.Reader) error {
	bytes, err := reader.ReadBytes(bencodeTerminator)
	if err != nil {
		return err
	}

	numberString := string(bytes[:len(bytes)-1])
	number, err := strconv.ParseInt(numberString, 10, 64)
	if err != nil {
		return err
	}

	integer.BencodeData = make([]byte, 0, len(bytes)+1)
	integer.BencodeData = append(integer.BencodeData, integerStartSymbol)
	integer.BencodeData = append(integer.BencodeData, bytes...)
	integer.Value = number

	return nil
}
