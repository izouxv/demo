package common

var MaxDepth int32 = 5
var DomainID64NewNodeID int64 = 1
var UserID64NewNodeID int64 = 2

/*RoleMap存储角色信息的map*/
var RoleMap = map[int32]string{
	Root:             "root",
	SuperAdmin:       "超级管理员",
	DomianAdmin:      "域管理员",
	ApplicationAdmin: "应用管理员",
	NodeAdmin:        "设备管理员",
	ReadOnly:         "普通用户",
	Operator:         "运营人员",
}

const (
	Root int32 = iota + 1
	SuperAdmin
	DomianAdmin
	ApplicationAdmin
	NodeAdmin
	ReadOnly
	Operator
)
