package torrentfile

import "GoTorrent/bencode"

// Info - info section in torrent file
type Info struct {
	PieceLength int64
	PieceHashes [][]byte
	IsPrivate   bool
	RootDir     string
	Files       []File
}

func loadInfoFromBencode(bc bencode.BencodeType) *Info {
	bcDict := bc.(*bencode.Dict)
	info := new(Info)

	info.PieceLength = bcDict.ValueAsInteger("piece length").Value

	piecesVal := bcDict.ValueAsByteString("pieces")
	if len(piecesVal.Value)%20 != 0 {
		panic("wrong pieces length")
	}
	piecesHashes := make([][]byte, 0, 1)
	for i := 0; i < len(piecesVal.Value)/20; i++ {
		piecesHashes = append(piecesHashes, piecesVal.Value[i*20:(i+1)*20])
	}
	info.PieceHashes = piecesHashes

	if bcDict.HasKey("private") {
		info.IsPrivate = (bcDict.ValueAsInteger("private").Value == 1)
	}

	if bcDict.HasKey("files") {
		// Режим нескольких файлов
		info.RootDir = string(bcDict.ValueAsByteString("name").Value)
		for _, item := range bcDict.ValueAsList("files").Value {
			file := loadFileFromBencodeDict(item)
			info.Files = append(info.Files, *file)
		}
	} else {
		// Режим одного файла
		info.RootDir = ""
		file := loadFileFromBencodeDict(bc)
		info.Files = append(info.Files, *file)
	}

	return info
}
