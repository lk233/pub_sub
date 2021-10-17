package main

import (
	"testProject/Component"
)

func main()  {
	xc := Component.NewXComponent()
	// 两个订阅者订阅topic名称为123的主题
	sub1 := Component.NewSubscriber(1)
	sub2 := Component.NewSubscriber(2)
	xc.Subscribe("123", sub1)
	xc.Subscribe("123", sub2)
	// 向topic名为123的主题发布信息
	xc.Publish("123", "1234")

	// 一个订阅者订阅topic名称为222的主题
	sub3 := Component.NewSubscriber(3)
	xc.Subscribe("222", sub3)
	// 向topic名为222的主题发布信息
	xc.Publish("222", "233")

}
