package db

import (
	"errors"
	"fmt"
	log "github.com/cihub/seelog"
	"sync"
	"github.com/garyburd/redigo/redis"
	"petfone-rpc/util"
)

const (
	nodeShift = 29                     // The machine code of digits to the left
	nodeMax   = -1 ^ (-1 << nodeShift) // Support the biggest machine id
)

type Node struct {
	sync.Mutex //Lock Mutex is an exclusive lock. Go language there are two state, locking and unlocked state,
	// when is locked, the operation of the other try to lock will be waiting, until the unlock.
	node int32 //The machine ID
	step int32 //The sequence
}

// IDs
func NewNode(node int32) (*Node, error) {
	switch node {
	case 1:
		return &Node{
			node: node * 1000000,
		}, nil
	case 2:
		return &Node{
			node: node * 20000000,
		}, nil
	case 3:
		return &Node{
			node: node * 20000000,
		}, nil
	default:
		return nil, errors.New("sid number must be between 1 and 3")
	}
}

// Generate creates and returns a unique snowflake ID
func (n *Node) Generate() (int32,error) {
	key := fmt.Sprintf("%s_%d", util.UidRedisKey, n.node)
	v, err := redis.Int64(Redis_Str(6379,util.Incr,key))
	if err != nil {
		log.Error(err)
		return 0,err
	}
	r := (n.node) + int32(v)
	return r,nil
}

func GenerateBigId() (int64,error) {
	return redis.Int64(Redis_Str(6379,util.Incr,util.BigIdRedisKey))
}
