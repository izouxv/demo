package permission

import (
    log "github.com/cihub/seelog"
	"cotx-http/redis"
)

const(
	 LL_ToB_WeChat          string = "AQEBAA=="
	 LL_ToB_Web             string = "AQECAA=="
	 LL_ToB_App             string = "AQEDAA=="
	 LL_ToC_WeChat          string = "AQIBAA=="
	 LL_ToC_Web             string = "AQICAA=="
	 LL_ToC_App             string = "AQIDAA=="

	 PL_ToB_WeChat          string = "AgEBAA=="
	 PL_ToB_Web             string = "AgECAA=="
	 PL_ToB_App			    string = "AgEDAA=="
	 PL_ToC_WeChat          string = "AgIBAA=="
	 PL_ToC_Web             string = "AgICAA=="
	 PL_ToC_App			    string = "AgIDAA=="

	 YA_ToB_WeChat			string = "AwEBAA=="
	 YA_ToB_Web				string = "AwECAA=="
	 YA_ToB_App				string = "AwEDAA=="
	 YA_ToC_WeChat			string = "AwIBAA=="
	 YA_ToC_Web				string = "AwICAA=="
	 YA_ToC_App				string = "AwIDAA=="
)
const(
	 LL_Explain             string = "lanlian"
	 pL_Explain             string = "penslink"
	 YA_Explain             string = "yuanan"
)

var Sources = map[string]string{
	LL_ToB_WeChat :            LL_Explain ,
	LL_ToB_Web :               LL_Explain ,
	LL_ToB_App :               LL_Explain ,
	LL_ToC_WeChat :            LL_Explain ,
	LL_ToC_Web :               LL_Explain ,
	LL_ToC_App :               LL_Explain ,

	PL_ToB_WeChat :            pL_Explain,
	PL_ToB_Web :               pL_Explain,
	PL_ToB_App :		       pL_Explain,
	PL_ToC_WeChat :            pL_Explain,
	PL_ToC_Web  :              pL_Explain,
	PL_ToC_App	:		       pL_Explain,

	YA_ToB_WeChat :			   YA_Explain,
	YA_ToB_Web	:			   YA_Explain,
	YA_ToB_App	:			   YA_Explain,
	YA_ToC_WeChat :			   YA_Explain,
	YA_ToC_Web	:			   YA_Explain,
	YA_ToC_App	:		       YA_Explain,
}

func  RedisInit() {
	log.Info("Start RedisInit")
	client := redis.RedisClient("persistence").Get()
	defer client.Close()
	for k, v := range Sources {
		reply, err := client.Do("set", redis.SourceRedisKeyPrefix + k , v)
		if err != nil || reply == 0 {
			log.Error("Set SourceRedis Failed,", err)
		}
	}
}