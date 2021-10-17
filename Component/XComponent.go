package Component

import (
	"fmt"
	"sync"
)

type XComponent struct {
	rw sync.RWMutex
	TMap map[string][]*Subscriber
	wg sync.WaitGroup
}

func NewXComponent() *XComponent {
	return &XComponent{TMap: make(map[string][]*Subscriber)}
}

// 发布函数
func (xc *XComponent) Publish(topic string, msg string) {
	xc.rw.RLock()
	subs, exist := xc.TMap[topic]
	if exist {
		len := len(subs)
		xc.wg.Add(len)
		for i := 0; i < len; i++ {
			go func(i int) {
				subs[i].SAction(msg)
				xc.wg.Done()
			}(i)
		}
	} else {
		fmt.Println("no such topic: %s", topic)
	}
	xc.wg.Wait()
	xc.rw.RUnlock()
}

// 订阅函数
func (xc *XComponent) Subscribe(topic string, sub *Subscriber) {
	xc.rw.RLock()
	_, exist := xc.TMap[topic]
	xc.rw.RUnlock()

	xc.rw.Lock()
	if !exist {
		subs := []*Subscriber{sub}
		xc.TMap[topic] = subs
	} else {
		xc.TMap[topic] = append(xc.TMap[topic], sub)
	}
	xc.rw.Unlock()
}