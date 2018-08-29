package id32

import (
	"account-domain-rpc/module"
	"errors"
	"fmt"
	log "github.com/cihub/seelog"
	"sync"
)

const (
	nodeShift = 29                     // The machine code of digits to the left
	nodeMax   = -1 ^ (-1 << nodeShift) // Support the biggest machine id32
)

type Node struct {
	sync.Mutex       //Lock Mutex is an exclusive lock. Go language there are two state, locking and unlocked state,
	node       int32 //The machine ID
	step       int32 //The sequence
}

// IDs
func NewNode(node int32) (*Node, error) {
	if node < 0 || node > nodeMax {
		return nil, errors.New("Node number must be between 000 and 011")
	}
	return &Node{
		node: node,
	}, nil
}

// Generate creates and returns a unique snowflake ID
func (n *Node) Generate() int32 {
	client := module.RedisClient("persistence").Get()
	key := fmt.Sprintf("%s_%d", module.UidRedisKey, n.node)
	idd, err := client.Do("incr", key)
	if err != nil {
		log.Error("redis incr error ", err)
		return 0
	} else {
		fmt.Println(err)
	}
	v, ok := idd.(int64)
	if !ok {
		fmt.Println(ok)
		return 0
	}
	r := (n.node)<<nodeShift | int32(v)
	defer client.Close()
	return r
}
