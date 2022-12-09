package fileutil

import "os"

var (
	// DefaultDirPerm 文件权限
	DefaultDirPerm   os.FileMode = 0775
	DefaultFilePerm  os.FileMode = 0665
	OnlyReadFilePerm os.FileMode = 0444

	// DefaultFileFlags 创建文件 只写 追加
	DefaultFileFlags = os.O_CREATE | os.O_WRONLY | os.O_APPEND
	CoverFileFlags   = os.O_CREATE | os.O_WRONLY | os.O_TRUNC
	// OnlyReadFileFlags 只读
	OnlyReadFileFlags = os.O_RDONLY
)

const (
	// MimeSniffLen 嗅探长度，用于检测文件 mime 类型
	MimeSniffLen = 512
)
