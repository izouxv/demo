package utils

import (
	"mime/multipart"
	"os"
	"io"
	log "github.com/cihub/seelog"
)
//
func CopyFileTo(filepath string, file multipart.File) int32 {

	//云端创建文件
    f ,err := os.Create(filepath)
	if err != nil {
		log.Errorf("工具类copy 文件失败 err === %s",err)
		return 30029
	}
	defer f.Close()

	//转移文件
	io.Copy(f,file)

	return  200
}

/*判断文件是否存在*/
func CheckFileIsExist(filepath string) bool  {
	_,err := os.Stat(filepath)
	if err == nil {
		//文件或者文件夹存在
		return true
	}
	if os.IsNotExist(err) {
		//文件不存在
		return false
	}
	return false
}
