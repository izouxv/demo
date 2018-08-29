package common

import "errors"

const (
	Air = 1 + iota
	Radacat
	Collar
	Gateway
	Probe
)

var NodeType = map[int32]string{
	Air:     "空气净化器",
	Collar:  "项圈",
	Gateway: "网关",
}

var IsLogin int32 = 1

var (
	DefaultLongitude = 116.487694
	DefaultLatitude  = 39.99632
)

var NodeID64NewNodeID int64 = 3
var NodeID64NewApplicationID int64 = 4

var (
	ErrAlreadyExists = errors.New("object already exists")
	ErrDoesNotExist  = errors.New("object does not exist")
	ErrCanNotDelete  = errors.New("object should not delete ")
)

var AdPlace = map[int32]string{
	1: "空气净化器广告位1",
	2: "空气净化器广告位2",
	3: "空气净化器广告位3",
}

var NodeAdPlace = map[int32]map[int32]string{
	Air: AdPlace,
}

var DownloadServer []string = []string{
	0: "http://iot.penslink.com:8082",
}

func GetHashDownloadServer(hash int64) string {
	index := hash % int64(len(DownloadServer))
	return DownloadServer[index]
}

// 最大深度
var MaxDepth int32 = 6
