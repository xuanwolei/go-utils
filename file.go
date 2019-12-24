package goutils

import (
	"os"
	"path/filepath"
)

const (
	FilePutFlag int = os.O_RDWR | os.O_CREATE | os.O_APPEND
)

//判断文件是否存在
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//读取文件
func FileGetContnet(path string) ([]byte, error) {
	exists, err := FileExists(path)
	if exists != true {
		return nil, err
	}
	return []byte(""), nil
}

//写文件
func FilePutContent(path string, content []byte, flag int) (int, error) {
	file, err := GetOsFileInstance(path, flag)
	if err != nil {
		return 0, err
	}
	n, err := file.Write(content)
	if err != nil {
		return 0, err
	}
	file.Close()
	return n, nil
}

//获取文件对象
func GetOsFileInstance(path string, flag int) (*os.File, error) {
	var (
		err    error
		file   *os.File
		exists bool
	)
	if flag == 0 {
		flag = FilePutFlag
	}
	exists, err = FileExists(path)
	if err != nil {
		return nil, err
	}
	//不存在文件，判断
	if exists != true {
		//判断是否存在目录
		dirPath := filepath.Dir(path)
		exists, err = FileExists(path)
		if err != nil {
			return nil, err
		}
		if exists != true {
			err = os.MkdirAll(dirPath, 0755)
			if err != nil {
				return nil, err
			}
		}
		//创建文件
		file, err = os.Create(path)
		if err != nil {
			return nil, err
		}
	} else {
		//打开文件
		file, err = os.OpenFile(path, flag, 0755)
		if err != nil {
			return nil, err
		}
	}

	return file, nil
}
