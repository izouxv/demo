package routers

import (
	"github.com/julienschmidt/httprouter"
	"petfone-http/controller"
	"net/http"
)

func settingRouter(router *httprouter.Router) {
	//常见问题
	router.GET(uriFaq, authInterceptor(controller.FaqCommonInfo))
	router.GET(uriChatFaq, authInterceptor(controller.PetChatFaqCommonInfo))
	//意见反馈
	router.POST(uriFeedback, authInterceptor(controller.SetFeedback))
	//获取新版本
	router.GET(uriVersion, authInterceptor(controller.GetVersion))
	router.GET(uriVersion1, authInterceptor(controller.GetVersion1))
	//获取广告
	router.GET(uriAdvertisement, authInterceptor(controller.GetAdver))
	//app设置
	router.PUT(uriSet, authInterceptor(controller.UpdatePetfoneFence))
	router.GET(uriSet, authInterceptor(controller.GetPetfoneFence))
}

func dataRouter(router *httprouter.Router) {
	//文件
	router.POST(uriFiles, authInterceptor(controller.BackUpFiles))
	router.GET(uriImages, authInterceptor(controller.GetImages))
	//router.POST(uriFiles1, authInterceptor(controller.UploadTrainRecording))			//修改宠端设备录音文件上传
	//运动数据
	router.POST(uriPetData, authInterceptor(controller.SetExerciseData))
	router.POST(uriDataPet, authInterceptor(controller.SetExerciseData1))
	router.GET(uriPetData, authInterceptor(controller.GetExerciseData))
	router.GET(uriDataPet, authInterceptor(controller.GetExerciseData1))
}

func deviceAndPetRouter(router *httprouter.Router) {
	//设备
	router.GET(uriDeviceDid, authInterceptor(controller.VerificationDevices))
	router.POST(uriDeviceDid, authInterceptor(controller.SetDevices))
	router.DELETE(uriDeviceDid, authInterceptor(controller.DeleteDevice))
	router.PUT(uriDeviceDid, authInterceptor(controller.UpdateDeviceInfo))
	router.GET(uriDevices, authInterceptor(controller.GetDevices))
	router.PATCH(uriStatistics, authInterceptor(controller.UploadTwinsAgent))
	//宠物资料
	router.POST(uriPets, authInterceptor(controller.SetPetInfo))
	router.PUT(uriPetPid, authInterceptor(controller.UpdatePetInfo))
	router.DELETE(uriPetPid, authInterceptor(controller.DeletePetInfo))
	router.GET(uriPetPid, authInterceptor(controller.GetPetInfo))
	router.GET(uriPets, authInterceptor(controller.GetPetInfos))
	//router.GET(uriPetPid1, authInterceptor(controller.GetPetInfo1_1))					//获取宠物详情和宠端设备录音
	//router.GET(uriPets1, authInterceptor(controller.GetPetInfos1_1))					//获取批量宠物详情和宠端设备录音
	//宠聊
	router.POST(uriChat, authInterceptor(controller.VoiceRecognition))
	router.POST(uriChat11, authInterceptor(controller.VoiceRecognition11))
	//router.POST(uriChat12, authInterceptor(controller.VoiceRecognition12))				// 接入讯飞AIUI服务

	//设备与宠物关联
	router.POST(uriPetpiddev, authInterceptor(controller.RelevanceDevicePet))
	router.DELETE(uriPetpiddev, authInterceptor(controller.UnRelevanceDevicePet))
	//宠物训练
	router.PUT(uriPetstrainid, authInterceptor(controller.UpdatePetTrainByPid))
	//router.PUT(uriPetstrainid1, authInterceptor(controller.UpdateDeviceTrainByPid))		//更新宠端设备训练项信息
	router.GET(uriPetstrain, authInterceptor(controller.GetPetTrainByPid))
	router.PATCH(uriPetstrainid, authInterceptor(controller.CounterPetTrainByPid))
	//共享
	router.POST(uriShare, authInterceptor(controller.SetShare))
	router.DELETE(uriShare, authInterceptor(controller.DeleteShare))
	router.GET(uriShare, authInterceptor(controller.GetShare))
	//宠物品种
	router.GET(uriFileBreeds, authInterceptor(controller.GetBreedInfos))
	router.GET(uriFileBreeds11, authInterceptor(controller.GetBreedInfos1_1))
}

func otherRouter(router *httprouter.Router) {
	router.ServeFiles("/swagger/*filepath", http.Dir("./swagger"))
	router.POST("/test", interceptor(controller.Test))
}


