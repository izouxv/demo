package controller

import (
	log "github.com/cihub/seelog"
	"petfone-http/result"
	"petfone-http/rpc"
	"petfone-http/pb"
	"petfone-http/po"
	"petfone-http/core"
	. "petfone-http/util"
	. "petfone-http/thirdPartyServer"
	"io/ioutil"
	"io"
	"os"
	"path"
	"strings"
	"unicode"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/disintegration/imaging"
	"strconv"
	"petfone-http/util"
	"bytes"
	"mime/multipart"
)

//语音识别,表单提交
func VoiceRecognition(res http.ResponseWriter, req *http.Request, params httprouter.Params)  {
	log.Info("VoiceRecognition......")
	length := req.ContentLength
	if length > 83886080 {
		log.Error("VoiceRecognition ContentLength:", length)
		result.ResCode(404, res)
		return
	}
	pidStr := params.ByName("pid")
	typesStr := req.PostFormValue("types")
	log.Info("VoiceRecognition-pidStr:", pidStr,",typesStr:",typesStr)
	if VerifyParamsStr(pidStr,typesStr) {
		result.RESC(21001, res)
		return
	}
	sso := GetContext(req)
	log.Info("VoiceRecognition-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	pid, err := StrToInt32(pidStr)
	types, err := StrToInt32(typesStr)
	if err != nil {
		result.RESC(21002, res)
		return
	}
	var text string
	switch types {
	case 1:
		text = req.PostFormValue("input")
	case 2:
		//todo 从HTTP中获取文件信息并保存
		err := req.ParseMultipartForm(32 << 20)
		if err != nil {
			log.Debug("VoiceRecognition-getHttpFile-err", err)
			result.RESC(21002, res)
			return
		}
		file, _, err := req.FormFile("input")
		if err != nil {
			log.Info("VoiceRecognition-fileErr:",err)
			result.RESC(21002, res)
			return
		}
		defer file.Close()
		bs, err := ioutil.ReadAll(file)
		if err != nil {
			log.Info("VoiceRecognition-dataErr:",err)
			result.RESC(21002, res)
			return
		}
		log.Info("VoiceRecognition-FileSize:", len(bs))
		xfResult := VoiceToText(bs)
		if xfResult.Code != "0" {
			log.Error("讯飞语音听写调用失败",xfResult)
			result.RESC(21002, res)
			return
		}
		text = xfResult.Data
	default:
		result.RESC(21001, res)
		return
	}
	if VerifyParamsStr(text) {
		log.Info(text)
		result.RESC(21002, res)
		return
	}
	petChat := &pb.PetChatRequest{Source: sso.Source, Pid: pid, Uid: sso.Uid,Input:strings.TrimSpace(strings.ToLower(text))}
	if unicode.Is(unicode.Scripts["Han"], []rune(text)[0]) {
		petChat.Language = pb.Language_Cn
	} else {
		petChat.Language = pb.Language_En
	}
	//todo 调用rpc
	petChatRe := rpc.PetChatRpc(petChat, "GetPetChatByPid")
	if petChatRe.Code != 10000 {
		result.RESC(petChatRe.Code, res)
		return
	}
	var chatMsgs []*po.ChatMsg
	for _,v := range petChatRe.ChatMsgs {
		chatMsg := &po.ChatMsg{Types:v.Types}
		switch v.Types {
		case 1:
			chatMsg.Output = v.Output1
		case 3:
			chatMsg.Output = v.Output2
		case 4:
			chatMsg.Output = v.Output2[0]
		}
		chatMsgs = append(chatMsgs,chatMsg)
	}
	result.REST(chatMsgs, res)
}
func VoiceRecognition11(res http.ResponseWriter, req *http.Request, params httprouter.Params)  {
	log.Info("VoiceRecognition11......")
	length := req.ContentLength
	if length > 83886080 {
		log.Error("VoiceRecognition11 ContentLength:", length)
		result.ResCode(404, res)
		return
	}
	pidStr := params.ByName("pid")
	typesStr := req.PostFormValue("types")
	languageStr := req.PostFormValue("language")
	log.Info("VoiceRecognition11-pidStr:", pidStr,",typesStr:",typesStr,",languageStr:",languageStr)
	if VerifyParamsStr(pidStr,typesStr,languageStr) {
		result.RESC(21001, res)
		return
	}
	sso := GetContext(req)
	log.Info("VoiceRecognition11-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	pid, err := StrToInt32(pidStr)
	types, err := StrToInt32(typesStr)
	if err != nil {
		result.RESC(21001, res)
		return
	}
	var text string
	switch types {
	case 1:
		text = req.PostFormValue("input")
	case 2:
		//todo 从HTTP中获取文件信息并保存
		err := req.ParseMultipartForm(32 << 20)
		if err != nil {
			log.Debug("VoiceRecognition11-getHttpFile-err", err)
			result.RESC(21002, res)
			return
		}
		file, _, err := req.FormFile("input")
		if err != nil {
			log.Info("VoiceRecognition11-fileErr:",err)
			result.RESC(21002, res)
			return
		}
		defer file.Close()
		bs, err := ioutil.ReadAll(file)
		if err != nil {
			log.Info("VoiceRecognition11-dataErr:",err)
			result.RESC(21002, res)
			return
		}
		log.Info("VoiceRecognition11-FileSize:", len(bs))
		xfResult := VoiceToText(bs)
		if xfResult.Code != "0" {
			log.Error("讯飞语音听写调用失败",xfResult)
			result.RESC(21002, res)
			return
		}
		text = xfResult.Data
	default:
		result.RESC(21001, res)
		return
	}
	if VerifyParamsStr(text) {
		result.RESC(21002, res)
		return
	}
	petChat := &pb.PetChatRequest{Source: sso.Source, Pid: pid, Uid: sso.Uid,
		Input:strings.TrimSpace(strings.ToLower(text))}
	switch pb.Language_value[languageStr] {
	case 1:
		petChat.Language = pb.Language_Cn
	case 2:
		petChat.Language = pb.Language_En
	default:
		result.RESC(21001, res)
		return
	}
	//todo 调用rpc
	petChatRe := rpc.PetChatRpc(petChat, "GetPetChatByPid")
	if petChatRe.Code != 10000 {
		result.RESC(petChatRe.Code, res)
		return
	}
	var chatMsgs []*po.ChatMsg
	for _,v := range petChatRe.ChatMsgs {
		chatMsg := &po.ChatMsg{Types:v.Types}
		switch v.Types {
		case 1:
			chatMsg.Output = v.Output1
		case 3:
			chatMsg.Output = v.Output2
		case 4:
			chatMsg.Output = v.Output2[0]
		}
		chatMsgs = append(chatMsgs,chatMsg)
	}
	result.REST(chatMsgs, res)
}
//func VoiceRecognition12(res http.ResponseWriter, req *http.Request, params httprouter.Params)  {
//	log.Info("VoiceRecognition12......")
//	length := req.ContentLength
//	if length > 83886080 {
//		log.Error("authInterceptor ContentLength:", length)
//		result.ResCode(404, res)
//		return
//	}
//	pidStr := params.ByName("pid")
//	//typesStr := req.PostFormValue("types")
//	//languageStr := req.PostFormValue("language")
//	typesStr := req.FormValue("types")
//	languageStr := req.FormValue("language")
//	log.Info("VoiceRecognition-pidStr:", pidStr,",typesStr:",typesStr,",languageStr:",languageStr)
//	if util.VerifyParamsStr(pidStr,typesStr,languageStr) {
//		result.RESC(21001, res)
//		return
//	}
//	sso := GetContext(req)
//	log.Info("VoiceRecognition-sso:", sso)
//	if sso == nil {
//		result.RESC(10001, res)
//		return
//	}
//	pid, err := util.StrToInt32(pidStr)
//	types, err := util.StrToInt32(typesStr)
//	if err != nil {
//		result.RESC(21001, res)
//		return
//	}
//	text := req.FormValue("input")
//	log.Debug("text: ",text)
//	var chatMsgs []*po.ChatMsg
//	log.Debug("### = ",util.ChatKeyMap[text])
//	if types == 1  && languageStr == "Cn" && (util.ChatKeyMap[text] > 7 || util.ChatKeyMap[text] < 1 ){
//		// 调用AIUI RPC
//		aiuiTextRequest := &pb.GetTextSemanticsRequest{Input:strings.TrimSpace(strings.ToLower(text)),Types:pb.SemanticsType_Text}
//		reply,err := rpc.GetTextSemantics(aiuiTextRequest)
//		log.Debugf("reply(%#v) ",reply)
//		if err != nil || reply == nil{
//			result.RESC(21001, res)
//			return
//		}
//		aiuiReply := make([]*po.ChatMsg,0)
//		log.Debug("aiuiReply = ",aiuiReply)
//		chatMgs := &po.ChatMsg{Types:types,Output:reply.Semantics}
//		aiuiReply = append(aiuiReply, chatMgs)
//		log.Debug("aiuiReply1 = ",aiuiReply)
//		result.REST(aiuiReply, res)
//		return
//	}else if types == 1 {
//		log.Debugf("用户language(%#v)！",languageStr)
//	} else if types == 2 {
//		//todo 从HTTP中获取文件信息并保存
//		err := req.ParseMultipartForm(32 << 20)
//		if err != nil {
//			log.Debug("VoiceRecognition-getHttpFile-err", err)
//			result.RESC(21002, res)
//			return
//		}
//		file, _, err := req.FormFile("input")
//		if err != nil {
//			log.Info("VoiceRecognition-fileErr:",err)
//			result.RESC(21002, res)
//			return
//		}
//		defer file.Close()
//		bs, err := ioutil.ReadAll(file)
//		if err != nil {
//			log.Info("VoiceRecognition-dataErr:",err)
//			result.RESC(21002, res)
//			return
//		}
//		log.Info("VoiceRecognition-FileSize:", len(bs))
//		xfResult := VoiceToText(bs)
//		if xfResult.Code != "0" {
//			log.Error("讯飞语音听写调用失败",xfResult)
//			result.RESC(21002, res)
//			return
//		}
//		text = xfResult.Data
//	}else {
//		result.RESC(21001, res)
//		return
//	}
//	if util.VerifyParamsStr(text) {
//		result.RESC(21002, res)
//		return
//	}
//	petChat := &pb.PetChatRequest{Source: sso.Source, Pid: pid, Uid: sso.Uid,
//		Input:strings.TrimSpace(strings.ToLower(text))}
//	switch pb.Language_value[languageStr] {
//	case 1:
//		petChat.Language = pb.Language_Cn
//	case 2:
//		petChat.Language = pb.Language_En
//	default:
//		result.RESC(21001, res)
//		return
//	}
//	//todo 调用rpc
//	petChatRe := rpc.PetChatRpc(petChat, "GetPetChatByPid")
//	if petChatRe.Code != 10000 {
//		result.RESC(petChatRe.Code, res)
//		return
//	}
//	for _,v := range petChatRe.ChatMsgs {
//		chatMsg := &po.ChatMsg{Types:v.Types}
//		switch v.Types {
//		case 1:
//			chatMsg.Output = v.Output1
//		case 3:
//			chatMsg.Output = v.Output2
//		case 4:
//			chatMsg.Output = v.Output2[0]
//		}
//		chatMsgs = append(chatMsgs,chatMsg)
//	}
//	result.REST(chatMsgs, res)
//}

//上传文件
func BackUpFiles(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("BackUpFiles......")
	length := req.ContentLength
	if length > 83886080 {
		log.Error("BackUpFiles ContentLength:", length)
		result.ResCode(404, res)
		return
	}
	useStr := params.ByName("use")
	numberStr := req.PostFormValue("number")
	idStr := req.PostFormValue("id")
	log.Info("BackUpFiles-use:", useStr, ",number:", numberStr, "id:", idStr)
	if VerifyParamsStr(useStr, numberStr) {
		result.RESC(21001, res)
		return
	}
	sso := GetContext(req)
	log.Info("BackUpFiles-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	err := req.ParseMultipartForm(5 << 20)
	if err != nil {
		log.Debug("BackUpFiles-err:", err)
		result.RESC(21002, res)
		return
	}
	_,fileHeader, err := req.FormFile("file")
	if err != nil {
		if fileHeader == nil {
			log.Info("file为空:",err)
			result.RESC(21001, res)
			return
		}
		log.Info("从HTTP中获取字节流失败:",err)
		result.RESC(21002, res)
		return
	}
	log.Info("BackUpFiles-Filename:", fileHeader.Filename)
	file,err := fileHeader.Open()
	if err != nil {
		log.Debug("BackUpFiles fileHeader.Open err:", err)
		result.RESC(21002, res)
		return
	}
	defer file.Close()
	buff := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(buff)
	fileWriter, err := bodyWriter.CreateFormFile("file", fileHeader.Filename)
	if err != nil {
		log.Error("SendFileHttp CreateFormFile err:",err)
		result.RESC(21002, res)
		return
	}
	copySize,err := io.Copy(fileWriter, file)
	if err != nil || copySize == 0 {
		log.Error("SendFileHttp Copy err:",err,copySize)
		result.RESC(21002, res)
		return
	}
	bodyWriter.Close()
	fid, err := SendFileHttp(core.ConstStr.FileServer,bodyWriter.FormDataContentType(),buff)
	if err != nil || fid == "" {
		log.Error("BackUpFiles err:", err)
		result.RESC(10001, res)
		return
	}
	switch useStr {
	case "account":
		account := &pb.AccountRequest{Avatar: fid}
		account.Uid = sso.Uid
		account.Source = sso.Source
		account.Token = sso.Token
		accountR := rpc.AccountRpc(account, "UpdateAccountInfo")
		if accountR.Code != 10000 {
			log.Info("UpdateUserInfo-AccountRpc Error:", accountR.Code)
			result.RESC(accountR.Code, res)
			return
		}
		result.REST(accountR, res)
	case "train":
		if VerifyParamsStr(idStr) {
			result.RESC(21001, res)
			return
		}
		id, err := StrToInt32(idStr)
		if err != nil {
			result.RESC(21001, res)
			return
		}
		number, err := StrToInt32(numberStr)
		if err != nil {
			result.RESC(21001, res)
			return
		}
		trainReq := &pb.PetTrainRequest{Source: sso.Source, Uid: sso.Uid, Pid: id,Id:number,Voice:fid,}
		log.Info("BackUpFiles-trainReq:", trainReq)
		//todo 调用rpc
		deviceRe := rpc.TrainRpc(trainReq, "UpdatePetTrainByPid")
		if deviceRe.Code != 10000 {
			result.RESC(deviceRe.Code, res)
			return
		}
		result.RESC(deviceRe.Code, res)
	case "pet":
		if VerifyParamsStr(idStr) {
			result.RESC(21001, res)
			return
		}
		id, err := StrToInt32(idStr)
		if err != nil {
			result.RESC(21001, res)
			return
		}
		petinfo := &pb.PetInfoRequest{Source: sso.Source, Pid: id, Uid: sso.Uid, Avatar: fid}
		//todo 调用rpc
		deviceRe := rpc.PetinfoRpc(petinfo, "UpdatePetInfoByPid")
		if deviceRe.Code != 10000 {
			result.RESC(deviceRe.Code, res)
			return
		}
		result.REST(deviceRe, res)
	default:
		result.ResCode(404, res)
	}
}

//下载文件
func GetImages(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("GetImages......")
	fidStr := params.ByName("fid")
	wStr := req.FormValue("w")
	hStr := req.FormValue("h")
	log.Info("GetImages fidStr:", fidStr,",view:", wStr,hStr)
	if VerifyParamsStr(fidStr) {
		result.RESC(21001, res)
		return
	}
	sso := GetContext(req)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	var w,h int
	var err error
	flag := VerifyParamsStr(wStr,hStr)
	if !flag {
		w,err = strconv.Atoi(wStr)
		h,err = strconv.Atoi(hStr)
		if err != nil {
			log.Info("GetImages view err:", err)
			result.RESC(21002, res)
			return
		}
	}
	clientRes,code := GetFileHttp(core.ConstStr.FileServer+fidStr)
	defer clientRes.Body.Close()
	if code != 10000 {
		log.Error("GetImages GetFileHttp code:",code)
		result.RESC(code, res)
		return
	}
	fileName := strings.Split(clientRes.Header.Get("content-disposition"),"\"")[1]
	format,err := imaging.FormatFromFilename(fileName)
	log.Info("GetImages imaging format:",format)
	if err == imaging.ErrUnsupportedFormat {
		log.Info("GetImages FormatFromFilename err:",err)
		result.RESC(21002, res)
		return
	}
	res.Header().Add("content-disposition","attachment; filename="+fileName)
	res.Header().Add("content-type","application/octet-stream")
	res.Header().Add("content-type","application/octet-stream;charset=UTF-8")
	if flag {
		io.Copy(res,clientRes.Body)
		return
	}
	log.Info("view:",w,h)
	imageFile, err := imaging.Decode(clientRes.Body)
	if err != nil {
		log.Error("GetImages open image err:", err)
		result.RESC(10001, res)
		return
	}
	dst := imaging.Resize(imageFile,w,h,imaging.Lanczos)
	imaging.Encode(res,dst,format)
}

func SaveFile(req *http.Request, types string) (int32, string) {
	//从HTTP中获取文件信息
	err := req.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Info("获取HTTP信息失败", err)
		return 21002, ""
	}
	log.Info("req-Type:", req.Header.Get("Content-Type"))
	file, fileHeader, err := req.FormFile("file")
	if err != nil {
		if file == nil {
			log.Info("file isEmpty")
			return 21006, ""
		}
		log.Info("从HTTP中获取字节流失败")
		return 21002, ""
	}
	log.Info("BackUpFiles-Filename:", fileHeader.Filename) //,"fileSize:",fileHeader.Size
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Info("文件读取失败")
		return 10001, ""
	}
	md5data := data
	md5 := CalMd5(md5data)
	log.Info("md5:", md5)
	//判断path是否存在，不存在则创建
	fileSuffix := path.Ext(fileHeader.Filename)          //获取文件后缀
	filenameWithSuffix := path.Base(fileHeader.Filename) //名称
	log.Info("file name: ", fileSuffix, filenameWithSuffix)
	pathFile := "E://" + types + "/" + md5 + fileSuffix
	isExists := ExistsPath(pathFile)
	address := "http://download.penslink.com:8082" + pathFile
	if isExists {
		log.Info("file is exist")
		return 10000, address
	}
	err = ioutil.WriteFile(pathFile, data, 0666)
	if err != nil {
		log.Info("文件写入失败")
		return 10001, ""
	}
	err = file.Close()
	if err != nil {
		file.Close()
		log.Info("结束文件失败")
		return 10001, ""
	}
	return 10000, address
}

var sema = make(chan struct{}, 20)

func dirents(file io.Reader) int {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadAll(file)
	if err != nil {
		log.Info(os.Stderr, "du:", err)
		return 0
	}
	return len(entries)
}

//更新宠端设备训练录音
func UploadTrainRecording(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	log.Info("UploadTrainRecording...")
	length := req.ContentLength
	if length > 83886080 {
		log.Error("authInterceptor ContentLength:", length)
		result.ResCode(404, res)
		return
	}
	sso := GetContext(req)
	log.Info("UploadTrainRecording-sso:", sso)
	if sso == nil {
		result.RESC(10001, res)
		return
	}
	////todo 从HTTP中获取文件信息并保存
	err := req.ParseMultipartForm(5 << 20)
	if err != nil {
		log.Debug("getHttpFile-err:", err)
		result.RESC(21002, res)
		return
	}
	_, fileHeader, err := req.FormFile("file")
	if err != nil {
		if fileHeader == nil {
			log.Info("file为空:",err)
			result.RESC(21001, res)
			return
		}
		log.Info("从HTTP中获取字节流失败:",err)
		result.RESC(21002, res)
		return
	}
	log.Info("UploadTrainRecording-Filename:", fileHeader.Filename)
	//todo 上传文件
	file,err := fileHeader.Open()
	if err != nil {
		log.Debug("UploadTrainRecording fileHeader.Open err:", err)
		result.RESC(21002, res)
		return
	}
	defer file.Close()
	buff := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(buff)
	fileWriter, err := bodyWriter.CreateFormFile("file", fileHeader.Filename)
	if err != nil {
		log.Error("SendFileHttp CreateFormFile err:",err)
		result.RESC(21002, res)
		return
	}
	copySize,err := io.Copy(fileWriter, file)
	if err != nil || copySize == 0 {
		log.Error("SendFileHttp Copy err:",err,copySize)
		result.RESC(21002, res)
		return
	}
	bodyWriter.Close()
	fid, err := SendFileHttp(core.ConstStr.FileServer,fileHeader.Filename,buff)
	if err != nil || fid == "" {
		log.Info("UpdateUserInfo-UploadFile-err:", err)
		result.RESC(10001, res)
		return
	}
	did, err := util.StrToInt32(req.PostFormValue("did"))
	if err != nil {
		result.RESC(21001, res)
		return
	}
	number, err := util.StrToInt32(req.PostFormValue("number"))
	if err != nil {
		result.RESC(21001, res)
		return
	}
	trainReq := &pb.DeviceTrainRequest{Source: sso.Source, Uid: sso.Uid, Did: did,Id:number,Voice:fid,}
	log.Info("UploadTrainRecording-DeviceTrainRequest:", trainReq)
	//todo 调用rpc
	deviceRe := rpc.UpdateDeviceTrainByDid(trainReq)
	if deviceRe.Code != 10000 {
		result.RESC(deviceRe.Code, res)
		return
	}
	result.RESC(deviceRe.Code, res)
}


