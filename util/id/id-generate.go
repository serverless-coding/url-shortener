package id

import (
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

// unique id generater
// ref: bytebytego post
// Twitter snowflake,分布式id生成器,int64类型
// 1bit    41bit   5bit       5 bit   12bit
// 符号位  时间戳   数据中心id  机器id  自增序列
// 1       41       3          7      10
const (
	epochMill         = int64(1577808000000)                                  // 设置起始时间(时间戳/毫秒)：2020-01-01 00:00:00，有效期69年
	epochSecond       = int64(1577808000)                                     // 设置起始时间(时间戳/秒)：2020-01-01 00:00:00，有效期69年
	timestampBits     = int64(41)                                             // 时间戳占用位数
	datacenteridBits  = int64(3)                                              // 数据中心id所占位数
	workeridBits      = int64(7)                                              // 机器id所占位数
	sequenceBits      = int64(12)                                             // 序列所占的位数
	timestampMax      = int64(-1 ^ (-1 << timestampBits))                     // 时间戳最大值
	datacenteridMax   = int64(-1 ^ (-1 << datacenteridBits))                  // 支持的最大数据中心id数量
	workeridMax       = int64(-1 ^ (-1 << workeridBits))                      // 支持的最大机器id数量
	sequenceMax       = int64(-1 ^ (-1 << sequenceBits))                      // 支持的最大序列id数量
	workeridShift     = int64(sequenceBits)                                   // 机器id左移位数
	datacenteridShift = int64(sequenceBits + workeridBits)                    // 数据中心id左移位数
	timestampShift    = int64(sequenceBits + workeridBits + datacenteridBits) // 时间戳左移位数
)

type Snowflake struct {
	mu           sync.Mutex
	timestamp    int64
	workerid     int64
	datacenterid int64
	sequence     int64
	isSecond     bool
}

// https://github.com/GUAIK-ORG/go-snowflake
func NewSnowflake(datacenterId, workId int64, isSecond bool) (*Snowflake, error) {
	if datacenterId < 0 || datacenterId > datacenteridMax || workId < 0 || workId > workeridMax {
		return nil, errors.Errorf("datacenterId=%d workId=%d not allowed", datacenterId, workId)
	}
	return &Snowflake{
		timestamp:    0,
		workerid:     workId,
		datacenterid: datacenterId,
		sequence:     0,
		isSecond:     isSecond,
	}, nil
}

func (s *Snowflake) Id() (id int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	del := int64(1)
	epoch := epochMill
	if s.isSecond {
		del = 1000
		epoch = epochSecond
	}

	now := time.Now().UnixMilli() / del

	if s.timestamp == now {
		s.sequence = (s.sequence + 1) & sequenceMax
		if s.sequence == 0 { // 超出序列承载范围,等到下一个时间戳
			for now <= s.timestamp {
				now = time.Now().UnixMilli() / del
			}
		}
	} else {
		s.sequence = 0
	}
	s.timestamp = now
	id = (now-epoch)<<timestampShift | (s.datacenterid << datacenteridShift) | (s.workerid << workeridShift) | s.sequence
	return
}

func NewRedis() (*redis.Redis, error) {
	pu, err := url.Parse(os.Getenv("KV_URL"))
	if err != nil {
		return nil, errors.Errorf("redis dsn err,%s", pu)
	}

	if pu.User == nil {
		return nil, errors.Errorf("redis dsn err,user==nil,%s", pu)
	}
	pass, _ := pu.User.Password()
	conf := redis.RedisConf{
		Host: pu.Host,
		Pass: pass,
		Tls:  true,
		Type: "node",
	}
	r, err := redis.NewRedis(conf)
	if err != nil {
		return nil, err
	}
	return r, nil
}
