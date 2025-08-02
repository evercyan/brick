package xstock

import (
	"time"
)

// ...
var (
	SleepDuration = time.Millisecond * 1500 // 避免请求频率过快被封 IP
	Debug         = false                   // 调试模式
)
