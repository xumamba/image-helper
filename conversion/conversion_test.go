package conversion

/**
 * @Time       : 2020/8/3 10:53 下午
 * @Author     : xulp
 * @Description: conversion_test.go
 */

import (
	"testing"
)

func TestConversion_FormatConversion(t *testing.T) {
	// Need to enter the absolute path of the image file
	conversionObj := InitConversion("$GOPATH/src/image-helper/conversion/test.jpg",
									"$GOPATH/src/image-helper/conversion/test.png")
	conversionObj.FormatConversion()
}
