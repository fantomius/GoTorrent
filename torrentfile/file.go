package torrentfile

import "GoTorrent/bencode"

// File - single file in torrent file
type File struct {
	Path   []string
	Length int64
	Md5sum []byte
}

func loadFileFromBencodeDict(bt bencode.BencodeType) *File {
	bcDict := bt.(*bencode.Dict)
	file := new(File)

	file.Length = bcDict.ValueAsInteger("length").Value

	if bcDict.HasKey("md5sum") {
		file.Md5sum = bcDict.ValueAsByteString("md5sum").Value
	}

	if bcDict.HasKey("name") {
		file.Path = append(file.Path, string(bcDict.ValueAsByteString("name").Value))
	} else if bcDict.HasKey("path") {
		for _, item := range bcDict.ValueAsList("path").Value {
			file.Path = append(file.Path, string(item.(*bencode.ByteString).Value))
		}
	} else {
		panic("No name or path key in file")
	}

	return file
}
