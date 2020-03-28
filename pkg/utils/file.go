package utils

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

// GetSize 获取文件大小 .
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// GetExt 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckExist 检查是否存在 .
func CheckExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckPermission 检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir 如果不存在则新建文件夹
func IsNotExistMkDir(dir string)  error{
	// 判断目录是否存在
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		// 必须分成两步：先创建文件夹、再修改权限
		err := os.MkdirAll(dir, 0777) //0777也可以os.ModePerm
		os.Chmod(dir, 0777)
		if err != nil {
			fmt.Printf("\n创建文件夹失败 failed![%v]\n", err)
			return err
		} else {
			fmt.Printf("\n 创建文件夹成功!\n")

			return nil
		}
	}
	if os.IsExist(err) {
		return nil
	}

	return err

}

// MkDir 创建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// OpenFile 打开
func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
