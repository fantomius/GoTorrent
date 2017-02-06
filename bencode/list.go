package bencode

import "bufio"

// List - representation of bencode list
type List struct {
	bencodeBase
	Value []BencodeType
}

// Length - number of elements
func (l *List) Length() int {
	return len(l.Value)
}

// ElementAsList - get child and force cast
func (l *List) ElementAsList(index int) *List {
	element := l.Value[index]
	return element.(*List)
}

// ElementAsDict - get child and force cast
func (l *List) ElementAsDict(index int) *Dict {
	element := l.Value[index]
	return element.(*Dict)
}

// ElementAsInteger - get child and force cast
func (l *List) ElementAsInteger(index int) *Integer {
	element := l.Value[index]
	return element.(*Integer)
}

// ElementAsByteString - get child and force cast
func (l *List) ElementAsByteString(index int) *ByteString {
	element := l.Value[index]
	return element.(*ByteString)
}

// Load - load list from bencode reader
func (l *List) load(reader *bufio.Reader) error {
	bencodeData := make([]byte, 0, 1)
	bencodeData = append(bencodeData, listStartSymbol)
	data := make([]BencodeType, 0, 1)

	for {
		bytes, err := reader.Peek(1)
		if err != nil {
			return err
		}

		if bytes[0] == bencodeTerminator {
			reader.ReadByte()
			bencodeData = append(bencodeData, bencodeTerminator)
			break
		}

		element, err := ReadBencode(reader)
		if err != nil {
			return err
		}
		data = append(data, element)
		bencodeData = append(bencodeData, element.GetBencodeData()...)
	}

	l.BencodeData = bencodeData
	l.Value = data

	return nil
}
