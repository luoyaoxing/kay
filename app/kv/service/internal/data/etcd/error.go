package etcd

import "fmt"

var (
	ERR_ETCD_EMPTY = EtcdError{10001, "操作etcd键值不能为空"}
	ERR_ETCD_PUT   = EtcdError{10002, "存储etcd失败"}
	ERR_ETCD_CLOSE = EtcdError{10003, "etcd客户端已经被删除"}
	ERR_ETCD_GET   = EtcdError{10004, "读取etcd失败"}
)

type EtcdError struct {
	ErrCode int
	ErrMsg  string
}

func (err EtcdError) Error() string {
	return fmt.Sprintf("[errCode]:%d [errMsg]:%s", err.ErrCode, err.ErrMsg)
}
