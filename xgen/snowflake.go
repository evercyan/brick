package xgen

import (
	"sync"
	"time"
)

/**
 * 1 	符号位, 保留字段, 固定为 0
 * 41  	时间戳 (当前时间-纪元时间)
 * 10 	机器 id
 * 12 	自增序列
 * +---+------------------------------------------------+----- -------+---------------+
 * | 1 | 41                                             | 10          | 12            |
 * +---+------------------------------------------------+-------------+---------------+
 * | 0 | 00000000 00000000 00000000 00000000 00000000 0 | 00000000 00 | 00000000 0000 |
 * +---+------------------------------------------------+-------------+---------------+
 */

const (
	twepoch        = int64(1577808000000)             // 开始时间截 (2020-01-01)
	workeridBits   = uint(10)                         // 机器id所占的位数
	sequenceBits   = uint(12)                         // 序列所占的位数
	workeridMax    = int64(-1 ^ (-1 << workeridBits)) // 支持的最大机器id数量
	sequenceMask   = int64(-1 ^ (-1 << sequenceBits)) // 自增序列最大
	workeridShift  = sequenceBits                     // 机器id左移位数
	timestampShift = sequenceBits + workeridBits      // 时间戳左移位数
)

type snowflake struct {
	sync.Mutex
	timestamp int64
	workerid  int64
	sequence  int64
}

var (
	sf   *snowflake
	once sync.Once
)

// NewSnowflake ...
func NewSnowflake(workerid int64) {
	if workerid < 0 || workerid > workeridMax {
		return
	}
	once.Do(func() {
		sf = &snowflake{
			timestamp: 0,
			workerid:  workerid,
			sequence:  0,
		}
	})
	return
}

// SnowflakeID ...
func SnowflakeID() int64 {
	sf.Lock()
	defer sf.Unlock()
	now := time.Now().UnixNano() / 1000000
	if sf.timestamp == now {
		sf.sequence = (sf.sequence + 1) & sequenceMask
		if sf.sequence == 0 {
			for now <= sf.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		sf.sequence = 0
	}
	sf.timestamp = now
	return (now-twepoch)<<timestampShift | (sf.workerid << workeridShift) | (sf.sequence)
}

func init() {
	sf = new(snowflake)
}
