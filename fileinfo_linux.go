package lumberjack

import (
	"os"
	"syscall"
	"time"
)

func GetFileCreateTime(fileInfo os.FileInfo) time.Time {
	statT := fileInfo.Sys().(*syscall.Stat_t)
	sec, nsec := statT.Ctim.Unix()
	return time.Unix(sec, nsec)
}
