package lumberjack

import (
	"os"
	"syscall"
	"time"
)

func GetFileCreateTime(fileInfo os.FileInfo) time.Time {
	// 获取文件的 BirthTime，这是 macOS 上的“创建时间”
	birthTime := fileInfo.Sys().(*syscall.Stat_t).Birthtimespec
	return time.Unix(birthTime.Sec, birthTime.Nsec)
}
