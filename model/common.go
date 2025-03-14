package model

import (
	"errors"
	"os"
	"strings"
)

/*
  Some common func.
*/

/*
GetProps
trans the bytes to props as map[string]string
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

/*
PathExists
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

var ErrPathIsFile error = errors.New("path is file")
var ErrPathIsDir error = errors.New("path is dir")
var ErrPathIsEmpty = errors.New("path is empty")

/*
FileExists
if the file exist, only file.
*/
func FileExists(path string) (bool, error) {
	if len(path) == 0 {
		return false, ErrPathIsEmpty
	}
	info, err := os.Stat(path)
	if err == nil {
		if info.IsDir() {
			return false, ErrPathIsDir
		}
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/*
DirExists
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
		return false, ErrPathIsFile
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// RemoveSliceElementByIndex
// remove an element from slice by its index
func RemoveSliceElementByIndex[E any](slice []E, index int) []E {
	if index >= 0 && index < len(slice) {
		return append(slice[:index], slice[index+1:]...)
	} else {
		return slice
	}
}

// RemoveOneSliceElementByElement
// Remove an element from slice
func RemoveOneSliceElementByElement[E comparable](slice []E, ele E) []E {
	for i, e := range slice {
		if e == ele {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func RemoveAllSliceElementsByElements[E comparable](slice []E, eles ...E) []E {
	temp := []E{}
	for _, e := range slice {
		needadd := true
		for _, ele := range eles {
			if e == ele {
				needadd = false
				break
			}
		}
		if needadd {
			temp = append(temp, e)
		}
	}
	return temp

}
