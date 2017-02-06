package bencode

import "bufio"

// Dict - representation of dictionary
type Dict struct {
	bencodeBase
	Value map[string]BencodeType
}

// HasKey - check existance of key
func (d *Dict) HasKey(key string) bool {
	_, ok := d.Value[key]
	return ok
}

// ValueAsList - get value by key and force cast
func (d *Dict) ValueAsList(key string) *List {
	return d.Value[key].(*List)
}

// ValueAsDict - get value by key and force cast
func (d *Dict) ValueAsDict(key string) *Dict {
	return d.Value[key].(*Dict)
}

// ValueAsInteger - get value by key and force cast
func (d *Dict) ValueAsInteger(key string) *Integer {
	return d.Value[key].(*Integer)
}

// ValueAsByteString - get value by key and force cast
func (d *Dict) ValueAsByteString(key string) *ByteString {
	return d.Value[key].(*ByteString)
}

func (d *Dict) load(reader *bufio.Reader) error {
	bencodeData := make([]byte, 0, 1)
	bencodeData = append(bencodeData, dictStartSymbol)
	data := make(map[string]BencodeType)

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

		// read key
		key := new(ByteString)
		err = key.load(reader)
		if err != nil {
			return err
		}
		keyString := string(key.Value)

		// read valye
		value, err := ReadBencode(reader)
		if err != nil {
			return err
		}

		bencodeData = append(bencodeData, key.BencodeData...)
		bencodeData = append(bencodeData, value.GetBencodeData()...)
		data[keyString] = value
	}

	d.BencodeData = bencodeData
	d.Value = data

	return nil
}
