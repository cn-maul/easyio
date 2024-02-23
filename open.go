package easyio

import (
	"log"
	"os"
)

func Open(path string, mode string) *os.File {
	switch mode {
	case "r": //read-only
		return openR(path)
	case "c": //create-write
		return openC(path)
	case "wa": //write-only-append
		return openWA(path)
	case "wc": //write-only-cover
		return openWC(path)
	default:
		return nil
	}
}

// 打开文件，只读
func openR(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		log.Println(err)
	}
	return file
}

// 创建并打开文件，只写
func openC(path string) *os.File {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Println(err)
	}
	return file
}

// 打开文件，追加写
func openWA(path string) *os.File {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Println(err)
	}
	return file
}

// 打开文件，覆盖写
func openWC(path string) *os.File {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0755)
	if err != nil {
		log.Println(err)
	}
	return file
}
