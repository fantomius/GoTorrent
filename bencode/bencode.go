package bencode

import "bufio"

// Special symbols
const integerStartSymbol = 105
const stringLengthValueDelimiter = 58
const listStartSymbol = 108
const dictStartSymbol = 100
const bencodeTerminator = 101

// BencodeType - base interface for all bencode elements
type BencodeType interface {
	GetBencodeData() []byte
}

// bencodeBase - base struct for all bencode elements
type bencodeBase struct {
	BencodeData []byte
}

func (b *bencodeBase) GetBencodeData() []byte {
	return b.BencodeData
}

// ReadBencode - read bencode from reader
func ReadBencode(reader *bufio.Reader) (BencodeType, error) {
	symbol, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	switch symbol {
	case integerStartSymbol:
		result := new(Integer)
		result.load(reader)
		return result, nil
	case dictStartSymbol:
		result := new(Dict)
		result.load(reader)
		return result, nil
	case listStartSymbol:
		result := new(List)
		result.load(reader)
		return result, nil
	default:
		err = reader.UnreadByte()
		if err != nil {
			return nil, err
		}
		result := new(ByteString)
		result.load(reader)
		return result, nil
	}
}
