package lumberjack

func GetFileCreateTime(fileInfo os.FileInfo) time.Time {
	// 获取文件的 BirthTime，这是 macOS 上的“创建时间”
	birthTime := info.Sys().(*syscall.Stat_t).Birthtimespec
	return time.Unix(0, int64(birthTime.Sec*1e9+birthTime.Nsec))
}
