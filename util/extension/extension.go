package extension

/**
 * @Time       : 2020/8/3 10:34 下午
 * @Author     : xulp
 * @Description: extension.go
 */

import (
	"path/filepath"
	"strings"
)

func GetFileExtName(filename string) string {
	return strings.ToLower(filepath.Ext(filename))
}
