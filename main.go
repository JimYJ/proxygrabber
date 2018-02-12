package main

import (
	"github.com/JimYJ/proxygrabber/checkproxy"
	"github.com/JimYJ/proxygrabber/grabber"
	"log"
)

var (
	fineProxy []map[string]string
)

func main() {
	ch := make(chan *[]map[string]string, 4)
	go grabber.GetKuaidaili(ch)
	go grabber.GetPcdaili(ch)
	go grabber.GetXicidaili(ch)
	go grabber.GetYundaili(ch)
	mapList := make(map[string]string)
	for i := 0; i < 4; i++ {
		temp := checkproxy.Check(<-ch)
		for j := 0; j < len(temp); j++ {
			value, ok := mapList[temp[j]["ip"]]
			if ok && value == mapList[temp[j]["port"]] {
				continue
			} else {
				mapList[temp[j]["ip"]] = mapList[temp[j]["port"]]
				fineProxy = append(fineProxy, temp[j])
			}
		}
	}
	log.Println(fineProxy)
}
