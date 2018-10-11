package main

import (
	"flag"
	"log"
	"time"
	"github.com/liaoxiaorong/wx/wx"


)

var addr = flag.String("addr", "0.0.0.0:7001", "listen addr")

func main() {
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := wx.Init()
	if err != nil {
		log.Fatal(err.Error())
	}
	go BoottimeTimingSettlement()
	go wx.Listening()

	log.Fatal(wx.WebServe(*addr))
}

//定时结算Boottime表数据
func BoottimeTimingSettlement() {
	us, err := wx.GetContacts()
	if err != nil {

	}
	var user string
	for _, u := range us {
		if u.NickName == "小冰" {
			user = u.UserName
		}
	}
	for {
		now := time.Now()
		// 计算下一个零点
		next := now.Add(time.Second*2)
		next = time.Date(next.Year(), next.Month(), next.Day(), next.Hour(), next.Minute(), next.Second(), 0, next.Location())

		t := time.NewTimer(next.Sub(now))
		<-t.C

		//以下为定时执行的操作
		Chat(user)
	}
}
func Chat(user string){
	//err := wx.SendMsg(user,"大冰，真的好无聊哦!")
	//if err!=nil{
	//	log.Printf("err: %s", err.Error())
	//}
}