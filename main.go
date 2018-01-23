package main

import (
	"log"
	"proxygrabber/checkproxy"
	"proxygrabber/grabber"
)

func main() {
	ch := make(chan *[]map[string]string, 10)
	go grabber.GetPcdaili(ch)
	fineProxy := checkproxy.Check(<-ch)
	log.Println(fineProxy)
}