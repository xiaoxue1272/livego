package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func GetCacheRoomFile(function func(file *os.File)) {
	dir, err := os.Getwd()
	if err == nil {
		var cacheFile *os.File
		filePathBuilder := strings.Builder{}
		filePathBuilder.WriteString(dir)
		filePathBuilder.WriteByte(os.PathSeparator)
		filePathBuilder.WriteString("room.cache")
		filePath := filePathBuilder.String()
		cacheFile, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_SYNC, 0755)
		if err != nil {
			log.Info("Open File[", filePath, "] error:", err)
		}
		function(cacheFile)
		_ = cacheFile.Close()
	}
	return
}
