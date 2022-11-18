package lumberjack

import (
	"os"
	"syscall"
	"time"
)

func GetFileCreateTime(fileInfo os.FileInfo) time.Time {
	wFileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)
	tNanSeconds := wFileSys.CreationTime.Nanoseconds()
	uSec := tNanSeconds / 1e6
	return time.UnixMicro(uSec)
}
