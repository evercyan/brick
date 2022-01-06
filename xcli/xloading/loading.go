// Package xloading 终端加载渲染
// 	l := xloading.New("message success").Start()
// 	time.Sleep(2 * time.Second)
// 	l.Success()
package xloading

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/evercyan/brick/xcli"
	"github.com/evercyan/brick/xcli/xcolor"
)

// ----------------------------------------------------------------

type loading struct {
	message   string        // 显示消息
	stopChan  chan struct{} // 中止信号
	startOnce sync.Once     // 开始标识
	stopOnce  sync.Once     // 中止标识
	style     Style         // 加载样式字符
	symbol    Symbol        // 状态样式字符
	speed     int           // 加载速度
}

// New ...
func New(messages ...string) Loading {
	return &loading{
		message:  strings.Join(messages, " "),
		stopChan: make(chan struct{}),
		style:    Style1,
		symbol:   Symbol1,
		speed:    100,
	}
}

// ----------------------------------------------------------------

// Style 加载样式
func (t *loading) Style(style Style) Loading {
	t.style = style
	return t
}

// Symbol 结束字符
func (t *loading) Symbol(symbol Symbol) Loading {
	t.symbol = symbol
	return t
}

// Speed 加载速度
func (t *loading) Speed(ms int) Loading {
	if ms < 10 {
		ms = 10
	}
	if ms > 1000 {
		ms = 1000
	}
	t.speed = ms
	return t
}

// Start 开始加载
func (t *loading) Start() Loading {
	t.startOnce.Do(func() {
		go func() {
			index := 0
			ticker := time.NewTicker(time.Millisecond * time.Duration(int64(t.speed)))
			for {
				select {
				case <-ticker.C:
					fmt.Printf("\r%s %s", t.style.Elements()[index], t.message)
					index = (index + 1) % len(t.style.Elements())
				case <-t.stopChan:
					return
				}
			}
		}()
	})
	return t
}

// Success 加载成功
func (t *loading) Success(messages ...string) {
	t.stop(true, messages...)
	return
}

// Fail 加载失败
func (t *loading) Fail(messages ...string) {
	t.stop(false, messages...)
	return
}

// ----------------------------------------------------------------

// stop 结束加载
func (t *loading) stop(success bool, messages ...string) {
	t.stopOnce.Do(func() {
		close(t.stopChan)
		xcli.ClearLine()
		message := t.message
		if len(messages) > 0 {
			message = strings.Join(messages, " ")
		}
		if success {
			xcolor.New("\r"+t.symbol.Elements()[0], message).Fg(xcolor.FgGreen).Render()
		} else {
			xcolor.New("\r"+t.symbol.Elements()[1], message).Fg(xcolor.FgRed).Render()
		}
	})
}
