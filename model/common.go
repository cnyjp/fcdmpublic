package model

import (
	bytes2 "bytes"
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

var defaultSysArgs SysArgs

func GetSysArgs() SysArgs {
	if nil == defaultSysArgs.Argmap {
		defaultSysArgs.Argmap = SysArgMap()
		defaultSysArgs.PeName = os.Args[0]
		if len(os.Args) > 1 {
			if strings.Index(os.Args[1], "-") == 0 {
				defaultSysArgs.Command = ""
			} else {
				defaultSysArgs.Command = os.Args[1]
			}
		} else {
			defaultSysArgs.Command = ""
		}
	}
	return defaultSysArgs
}

// 对命令行字符串进行转义处理
func EscapeArgString(value string) string {
	bytes := bytes2.Buffer{}
	for _, c := range value {
		if c == '"' {
			bytes.WriteString("\\")
		}
		bytes.WriteByte(byte(c))
	}
	return bytes.String()
}

/*
解析系统Args，组成为map
系统args采用统一的格式：
program {-argname} {argvalue} {-argname2} {argvalue2}
即：参数名使用-作为前缀，其后的为其参数数值；对于参数数值中有空格或-开头的，使用双引号包含起来；对于双引号中的双引号使用\进行转义；
同样，\字符本身使用\\进行转义

-help作为一个特殊的命令主体，-help后面的参数将最为命令行-c指令的帮助进行输出。如果没有对应的命令，则直接输出命令行自身的帮助信息。
*/
func SysArgMap() map[string]string {

	var hmap = make(map[string]string)
	var currentname string
	for index, value := range os.Args {
		if index == 0 { //第一个参数是程序名，不进行解析
			continue
		}
		if strings.HasPrefix(value, "-") {
			currentname = strings.TrimPrefix(value, "-")
			hmap[currentname] = "" //给一个空串的默认值，否则无法承接无参数命令
			continue
		}
		hmap[currentname] = value
	}
	return hmap
}
