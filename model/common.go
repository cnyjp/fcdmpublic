package model

import (
	"errors"
	"os"
	"strings"
)

/**
  Some common func.
*/

/**
  trans the bytes to props like map[string]string
*/
func GetProps(ab []byte) map[string]string {
	props := map[string]string{}
	if len(ab) == 0 {
		return props
	}
	all := string(ab)
	all = strings.ReplaceAll(all, "\r\n", "\n")
	lines := strings.Split(all, "\n")
	for _, v := range lines {
		if len(v) == 0 {
			continue
		}
		if strings.Index("#", v) == 0 {
			continue
		}
		splits := strings.SplitN(v, "=", 2)
		if len(splits) != 2 {
			continue
		}
		if strings.HasPrefix(splits[1], "\"") && strings.HasSuffix(splits[1], "\"") {
			props[splits[0]] = strings.Trim(splits[1], "\"")
		} else {
			props[splits[0]] = splits[1]
		}
	}
	return props
}

/**
  if the path exist, no matter the path is a file or dir, will return true.
*/
func PathExists(path string) (bool, error) {

	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**
  if the file exist, only file.
*/
func FileExists(path string) (bool, error) {
	if len(path) == 0 {
		return false, errors.New("path is empty")
	}
	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			return false, errors.New("path is a directory, not a file")
		}
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**
  if the dir exist, only dir.
*/
func DirExists(path string) (bool, error) {
	if len(path) == 0 {
		return false, errors.New("path is empty")
	}
	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			return true, nil
		}
		return false, errors.New("path is not a directory")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
