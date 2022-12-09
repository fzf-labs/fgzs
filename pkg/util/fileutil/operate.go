package fileutil

import (
	"archive/zip"
	"bufio"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

// ************************************************************
//	dir
// ************************************************************

// Mkdir 创建文件夹
func Mkdir(dirPath string) error {
	return os.MkdirAll(dirPath, DefaultDirPerm)
}

// ************************************************************
// files
// ************************************************************

// OpenFile 打开文件，但会自动创建目录。
func OpenFile(filepath string, flag int, perm os.FileMode) (*os.File, error) {
	fileDir := path.Dir(filepath)

	// if err := os.Mkdir(dir, 0775); err != nil {
	if err := os.MkdirAll(fileDir, DefaultDirPerm); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(filepath, flag, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// QuickOpenFile 快速打开文件，目录不存在则会自动创建目录。
func QuickOpenFile(filepath string) (*os.File, error) {
	return OpenFile(filepath, DefaultFileFlags, DefaultFilePerm)
}

// OpenReadFile 只读方式打开文件
func OpenReadFile(filepath string) (*os.File, error) {
	return os.OpenFile(filepath, OnlyReadFileFlags, OnlyReadFilePerm)
}

// ReadFileLineToSli  按行读取文件
func ReadFileLineToSli(dir string) ([]string, error) {
	file, err := os.OpenFile(dir, os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	res := make([]string, 0)
	for {
		line, _, err := buf.ReadLine()
		context := string(line)
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		res = append(res, context)
	}
	return res, nil
}

// ReadFileByUrlToByte 读取url中的文件,并转为[]byte格式
func ReadFileByUrlToByte(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// ************************************************************
//	write, copy files
// ************************************************************

// WriteContentCover 数据写入，不存在则创建
func WriteContentCover(filePath, content string) error {
	dstFile, err := OpenFile(filePath, CoverFileFlags, DefaultFilePerm)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = dstFile.WriteString(content)
	if err != nil {
		return err
	}
	return err
}

// WriteContentAppend 数据写入，不存在则创建
func WriteContentAppend(filePath, content string) error {
	dstFile, err := OpenFile(filePath, DefaultFileFlags, DefaultFilePerm)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	_, err = dstFile.WriteString(content)
	if err != nil {
		return err
	}
	return err
}

// CopyFile 复制文件
func CopyFile(srcPath string, dstPath string) error {
	srcFile, err := os.OpenFile(srcPath, os.O_RDONLY, 0)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// create and open file
	dstFile, err := QuickOpenFile(dstPath)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

// ************************************************************
//	remove
// ************************************************************

// Remove removes the named file or (empty) directory.
func Remove(fPath string) error {
	if PathExists(fPath) {
		return os.Remove(fPath)
	}
	return nil
}

// ************************************************************
//	other operates
// ************************************************************

// Unzip a zip archive
// from https://blog.csdn.net/wangshubo1989/article/details/71743374
func Unzip(archive, targetDir string) (err error) {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err = os.MkdirAll(targetDir, DefaultDirPerm); err != nil {
		return
	}

	for _, file := range reader.File {
		fullPath := filepath.Join(targetDir, file.Name)
		if file.FileInfo().IsDir() {
			err = os.MkdirAll(fullPath, file.Mode())
			if err != nil {
				return err
			}
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}

		targetFile, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			_ = fileReader.Close()
			return err
		}

		_, err = io.Copy(targetFile, fileReader)

		// close all
		_ = fileReader.Close()
		targetFile.Close()

		if err != nil {
			return err
		}
	}

	return
}

// DownloadFile 会将url下载到本地文件，它会在下载时写入，而不是将整个文件加载到内存中。
func DownloadFile(url, filepath string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func FilePrefix(filename string) string {
	filenameall := path.Base(filename)
	return filenameall[0 : len(filenameall)-len(path.Ext(filename))]
}
