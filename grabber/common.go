package grabber

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var (
	kuaidailipage = 10
	xicidailipage = 10
	yundailipage  = 4
	pcdailipage   = 1
	showErrors    = true
)

//Charset 字符集
type Charset string

const (
	//UTF8 UTF8
	UTF8 = Charset("UTF-8")
	//GB18030 GBK
	GB18030 = Charset("GB18030")
	//HZGB2312 GB2312
	HZGB2312 = Charset("HZ-GB2312")
)

func setPcdailiURL(types int, curpage int) string {
	return fmt.Sprintf("http://www.pcdaili.com/index.php?m=daili&a=free&type=%d&page=%d", types, curpage)
}

func setKuaidailiURL(curpage int) string {
	return fmt.Sprintf("https://www.kuaidaili.com/free/inha/%d", curpage)
}

func setxXcidailiURL(curpage int) string {
	return fmt.Sprintf("http://www.xicidaili.com/nn/%d", curpage)
}

func setYundailiURL(curpage int) string {
	return fmt.Sprintf("http://www.yun-daili.com/free.asp?page=%d", curpage)
}

func printErrors(err ...interface{}) {
	if err != nil && showErrors == true {
		log.Println(err)
	}
}

func remove(slice []string, index int) []string {
	return append(slice[0:index], slice[index+1:]...)
}

func removeSpaceElem(strArr []string) []string {
	for i := 0; i < len(strArr); i++ {
		if strings.TrimSpace(strArr[i]) == "" {
			strArr = remove(strArr, i)
		}
	}
	return strArr
}

//ConvertByte2String 转换编码
func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case HZGB2312:
		var decodeBytes, _ = simplifiedchinese.HZGB2312.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}

	return str
}

//ConvertEncoder 转换编码
func ConvertEncoder(str string, charset string) string {
	var s string
	switch charset {
	case "GBK":
		data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(str)), simplifiedchinese.GBK.NewEncoder()))
		return string(data)
	case "GB2312":
		data, _ := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(str)), simplifiedchinese.HZGB2312.NewEncoder()))
		return string(data)
	case "UTF8":
		return str
	default:
		s = str
	}
	return s
}
