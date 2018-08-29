package advertisement

import (
	"account-domain-rpc/common"
	"account-domain-rpc/module"
	"account-domain-rpc/storage"

	log "github.com/cihub/seelog"
	"github.com/jinzhu/gorm"
)

func HandlerNodeAdvertisement(nas []*storage.NodeAdv, devEUIstr string) (result []*storage.NodeAdv, ok bool) {
	devEUI := common.EUI64{}
	err := devEUI.UnmarshalString(devEUIstr)
	if err != nil {
		log.Errorf("devEUI错误")
		return
	}
	node := storage.Node{DevEUI: devEUI[:]}
	err = node.GetNodeByDevEUI()
	if err != nil {
		log.Errorf("设备不存在 %s", devEUIstr)
		return
	}
	qi := storage.NodeAdvQueue{Nid: node.Nid}
	if !qi.IsExistQueue() {
		log.Errorf("队列中不存在")
		return
	}
	nainfos, err := qi.GetNodeAdvInfoByNid()
	if err != nil {
		log.Errorf("获取设备的广告异常")
		return
	}
	err = storage.Transaction(module.MysqlClient(), func(tx *gorm.DB) error {
		for _, v := range nas {
			log.Info(v.AdPlace)
			for _, vv := range nainfos {
				if v.AdPlace == vv.AdPlace {
					if v.Md5 == vv.Md5 {
						qi.NodeAdvId = vv.Id
						qi.Delete(tx)
					} else {
						v.Md5 = vv.Md5
						v.Url = vv.Url
						result = append(result, v)
					}
					break
				}
			}
		}
		return nil
	})

	if len(result) == 0 {
		log.Errorf("队列中不存在")
		return
	}
	ok = true
	return
}
