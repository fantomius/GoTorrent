package main

import (
	"GoTorrent/bencode"
	"GoTorrent/torrentfile"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func printUsage() {

}

func loadTorrentFile(torrentFile string) (result *torrentfile.Torrent) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("A problem occured while loading torrent file ", r)
			result = nil
		}
	}()

	file, err := os.Open(torrentFile)
	if err != nil {
		fmt.Println("A problem occured while reading torrent file ", err)
		return nil
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	bencodeData, err := bencode.ReadBencode(reader)
	if err != nil {
		fmt.Println("A problem occured while decoding torrent file ", err)
		return nil
	}

	result = torrentfile.LoadTorrentFromBencode(bencodeData)
	return result
}

func fillStringToLength(s string, length int) string {
	requiredSeparators := length - len(s)
	if requiredSeparators <= 0 {
		return s
	}

	return s + strings.Repeat(" ", requiredSeparators)
}

func torrentInfo(torrentFile string) {
	torrent := loadTorrentFile(torrentFile)
	if torrent == nil {
		fmt.Println("Failed to load to torrent in info command")
		return
	}

	fmt.Println("Info for '", torrentFile, "' torrent file")
	fmt.Println(strings.Repeat("#", 60))
	fmt.Println(strings.Repeat("#", 26), " INFO ", strings.Repeat("#", 26))
	fmt.Println(fillStringToLength("Piece Length: ", 20), "  ", torrent.Info.PieceLength)
	fmt.Println(fillStringToLength("Piece Hashes: ", 20), "  ", "[ ", len(torrent.Info.PieceHashes), " elements ]")
	fmt.Println(fillStringToLength("Is Private: ", 20), "  ", torrent.Info.IsPrivate)

	fmt.Println()
	fmt.Println(strings.Repeat("#", 60))
	fmt.Println(strings.Repeat("#", 25), " TORRENT ", strings.Repeat("#", 24))

	fmt.Println(fillStringToLength("Info Hash: ", 20), "  ", string(torrent.InfoHash[:]))
	fmt.Println(fillStringToLength("Announce: ", 20), "  ", torrent.Announce)
	for ind1, innerList := range torrent.AnnounceList {
		for ind2, item := range innerList {
			fmt.Println(fillStringToLength(fmt.Sprint("Announce List[", ind1, "][", ind2, "]"), 20), "  ", item)
		}
	}
	fmt.Println(fillStringToLength("Creation Date: ", 20), "  ", torrent.CreationDate.Format(time.UnixDate))
	fmt.Println(fillStringToLength("Comment: ", 20), "  ", torrent.Comment)
	fmt.Println(fillStringToLength("Comment: ", 20), "  ", torrent.Comment)
	fmt.Println(fillStringToLength("Created By: ", 20), "  ", torrent.CreatedBy)
	fmt.Println(fillStringToLength("Encoding: ", 20), "  ", torrent.Encoding)

	fmt.Println()
	fmt.Println(strings.Repeat("#", 60))
	fmt.Println(strings.Repeat("#", 26), " FILES ", strings.Repeat("#", 25))

	fmt.Println("Root Dir: ", torrent.Info.RootDir)
	for _, file := range torrent.Info.Files {
		fmt.Println(strings.Join(file.Path, "/"), "(md5:", file.Md5sum, "size:", file.Length, "Bytes)")
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid arguments count")
		printUsage()
		return
	}

	command := os.Args[1]
	torrentFile := os.Args[2]

	switch strings.ToLower(command) {
	case "info":
		torrentInfo(torrentFile)
	default:
		fmt.Println("Invalid command")
		printUsage()
	}
}
