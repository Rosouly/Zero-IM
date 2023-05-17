package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"math/rand"
	"time"
)

var (
	c           = make(chan string)
	nd          *snowflake.Node
	NodId       int64
	uniquePodId string
)

func init() {
	rand.Seed(time.Now().UnixNano())
	NodId = int64(rand.Intn(128))
	nd, _ = snowflake.NewNode(NodId)
	uniquePodId = nd.Generate().String()
	go loop()
}

func loop() {
	for {
		c <- nd.Generate().String()
	}
}

// GetID 获取ID
// 参数: podId 当前机器标识 这里使用 pod ip
func GetID() string {
	md5 := Md5(uniquePodId + <-c)
	//fmt.Println("md5:", md5)
	return md5
}

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}
