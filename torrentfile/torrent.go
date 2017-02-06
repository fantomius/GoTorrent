package torrentfile

import (
	"GoTorrent/bencode"
	"crypto/sha1"
	"time"
)

// Torrent - структура из содержимого torrent-файла
type Torrent struct {
	Info         *Info
	InfoHash     [20]byte
	Announce     string
	AnnounceList [][]string
	CreationDate time.Time
	Comment      string
	CreatedBy    string
	Encoding     string
}

// LoadTorrentFromBencode - загружает torrent-структуру из bencode-словаря
func LoadTorrentFromBencode(bc bencode.BencodeType) *Torrent {
	bcDict := bc.(*bencode.Dict)
	torrent := new(Torrent)

	// Парсим info
	torrent.Info = loadInfoFromBencode(bcDict.Value["info"])
	torrent.InfoHash = sha1.Sum(bcDict.Value["info"].GetBencodeData())

	torrent.Announce = string(bcDict.ValueAsByteString("announce").Value)
	if bcDict.HasKey("comment") {
		torrent.Comment = string(bcDict.ValueAsByteString("comment").Value)
	}
	if bcDict.HasKey("created by") {
		torrent.CreatedBy = string(bcDict.ValueAsByteString("created by").Value)
	}
	if bcDict.HasKey("encoding") {
		torrent.Encoding = string(bcDict.ValueAsByteString("encoding").Value)
	}
	if bcDict.HasKey("creation date") {
		torrent.CreationDate = time.Unix(bcDict.ValueAsInteger("creation date").Value, 0)
	}
	if bcDict.HasKey("announce-list") {
		result := make([][]string, 0, 1)
		for _, item := range bcDict.ValueAsList("announce-list").Value {
			itemResult := make([]string, 0, 1)
			itemList := item.(*bencode.List)
			for _, innerItem := range itemList.Value {
				strItem := innerItem.(*bencode.ByteString)
				itemResult = append(itemResult, string(strItem.Value))
			}
			result = append(result, itemResult)
		}
		torrent.AnnounceList = result
	}

	return torrent
}
