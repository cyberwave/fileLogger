// Package: fileLogger
// File: util.go
// Created by: mint(mint.zhao.chiu@gmail.com)_aiwuTech
// Useage: some useful utils
// DATE: 14-8-23 17:03
package fileLogger

import (
	"os"
	"path/filepath"
	"fmt"
)

// Determine a file or a path exists in the os
func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// joinFilePath joins path & file into a single string
func joinFilePath(path, file string) string {
	return filepath.Join(path, file)
}

func createDir( path string ){
	os.Mkdir( path, 0755)
}

func createFile( fileName string )( file *os.File, err error ){
	file, err = os.OpenFile( fileName, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	return
}

// return length in bytes for regular files
func fileSize(file string) int64 {
	f, e := os.Stat(file)
	if e != nil {
		fmt.Println(e.Error())
		return 0
	}

	return f.Size()
}

// return file name without dir
func shortFileName(file string) string {
	return filepath.Base(file)
}
