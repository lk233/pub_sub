package Component

import "fmt"

type Subscriber struct {
	Id int64
}

func NewSubscriber(Id int64) *Subscriber {
	return &Subscriber{Id: Id}
}

type Subscribe interface {
	SAction(msg string)
}

func (sub *Subscriber)SAction(msg string) {
	// 回调函数,处理消息
	fmt.Printf("sub ID: %d, msg: %s\n", sub.Id, msg)
}