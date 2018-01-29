package grabber

import (
	"fmt"
	"log"
	"strings"
)

var (
	kuaidailipage int  = 10
	xicidailipage int  = 10
	pcdailipage   int  = 1
	showErrors    bool = true
	spysone            = "http://spys.one/en/http-proxy-list/"
)

func setPcdailiURl(types int, curpage int) string {
	return fmt.Sprintf("http://www.pcdaili.com/index.php?m=daili&a=free&type=%d&page=%d", types, curpage)
}

func setKuaidailiURl(curpage int) string {
	return fmt.Sprintf("https://www.kuaidaili.com/free/inha/%d", curpage)
}

func setxXcidailiiUrl(curpage int) string {
	return fmt.Sprintf("http://www.xicidaili.com/nn/%d", curpage)
}

func setYundailiUrl(curpage int) string {
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
