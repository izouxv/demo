package air

import (
	"crypto/aes"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"

	"account-domain-rpc/common"
	"account-domain-rpc/link/standard"
	"account-domain-rpc/module"
	"account-domain-rpc/storage"
	log "github.com/cihub/seelog"
)

//入网请求
type PHYPayload struct {
	MHDR    MHDR    `json:"mhdr"`
	Payload Payload `json:"data"`
	MIC     MIC     `json:"mic"`
}

// MarshalBinary phy to []byte
func (phy PHYPayload) MarshalBinary() ([]byte, error) {
	if phy.Payload == nil {
		return []byte{}, errors.New("plink: Data should not be nil")
	}

	var out []byte
	var b []byte
	var err error

	if b, err = phy.MHDR.MarshalBinary(); err != nil {
		return []byte{}, err
	}
	out = append(out, b...)

	if b, err = phy.Payload.MarshalBinary(); err != nil {
		return []byte{}, err
	}
	out = append(out, b...)
	out = append(out, phy.MIC[0:len(phy.MIC)]...)
	return out, nil
}

// UnmarshalBinary []byte to Phy
func (phy *PHYPayload) UnmarshalBinary(data []byte) error {
	if len(data) < 3 {
		return errors.New("plink: at least 3 bytes needed to decode PHYPayload")
	}
	// MHDR
	if err := phy.MHDR.UnmarshalBinary(data[0:2]); err != nil {
		return err
	}
	// MACPayload
	switch byte(phy.MHDR.MType) {
	case JoinRequest:
		phy.Payload = &JoinRequestPayload{}
	default:
		return errors.New("plink: no type to decode PHYPayload,only JoinRequest")
	}
	if err := phy.Payload.UnmarshalBinary(data[2 : len(data)-4]); err != nil {
		return err
	}
	// MIC
	for i := 0; i < 4; i++ {
		phy.MIC[i] = data[len(data)-4+i]
	}
	return nil
}

func (phy PHYPayload) HandlerDataPayload(node storage.Node) ([]byte, error) {
	log.Info("===============JoinRequest==================", phy)
	equal := func(b1, b2 []byte) bool {
		if b1 == nil && b2 == nil {
			return true
		}
		if b1 == nil {
			return false
		}
		for i, v := range b1 {
			if b2[i] != v {
				return false
			}
		}
		return true
	}
	getAppNonce := func() ([3]byte, error) {
		var b [3]byte
		if _, err := rand.Read(b[:]); err != nil {
			return b, err
		}
		return b, nil
	}
	getPlinkNwkSKey := func(appkey common.AES128Key, netID standard.NetID, appNonce [3]byte, devNonce [2]byte) (common.AES128Key, error) {
		log.Info("-------appkey--------", appkey)
		log.Info("-------netID--------", netID)
		log.Info("-------appNonce--------", appNonce)
		log.Info("-------devNonce--------", devNonce)
		padding := [8]byte{8, 8, 8, 8, 8, 8, 8, 8}
		var nwkSkey common.AES128Key
		src := make([]byte, 0)
		src = append(src, appNonce[:]...)
		src = append(src, netID[:]...)
		src = append(src, devNonce[:]...)
		src = append(src, padding[:]...)
		if len(src) != len(nwkSkey) {
			return nwkSkey, errors.New("src len not 16 byte")
		}
		block, err := aes.NewCipher(appkey[:])
		if err != nil {
			return nwkSkey, err
		}
		if block.BlockSize() != len(src) {
			return nwkSkey, fmt.Errorf("block-size of %d bytes is expected", len(src))
		}
		block.Encrypt(nwkSkey[:], src)
		return nwkSkey, nil
	}
	var err error
	var netID standard.NetID
	netID = [3]byte{0x01, 0x02, 0x03}
	payload := phy.Payload.(*JoinRequestPayload)
	app := &storage.Application{Aid: node.ApplicationID}
	if err := app.GetApplicationByAID(); err != nil {
		log.Errorf("ApplicationId: ", node.ApplicationID, "get application error: %s", err)
		return nil, err
	}
	if !equal(app.AppEUI, payload.AppEUI[:]) {
		log.Errorf("join-request DevEUI exists, but with a different AppEUI,dev_eui:%x ",
			payload.DevEUI, "expected_app_eui:%x ", app.AppEUI, "request_app_eui:%x ", payload.AppEUI)
		return nil, errors.New("DevEUI exists, but with a different AppEUI")
	}
	var AppKey common.AES128Key
	AppKey.Scan(node.AppKey)
	ok, err := phy.ValidateMIC(AppKey)
	if err != nil {
		log.Errorf("dev_eui:", node.DevEUI, "app_eui:", payload.AppEUI, "join-request validate mic error: %s", err)
		return nil, err
	}
	if !ok {
		log.Errorf("dev_eui:", node.DevEUI, "app_eui:", app.AppEUI, "mic:", phy.MIC, "join-request invalid mic")
		return nil, errors.New("invalid MIC")
	}

	if equal(node.UsedDevNonces, payload.DevNonce[:]) {
		log.Errorf("join-request DevNonce has already been use,dev_eui:%s, app_eui:%s, dev_nonce:%s", node.DevEUI, payload.AppEUI, payload.DevNonce)
		return nil, errors.New("DevNonce has already been used")
	} else {
		node.UsedDevNonces = payload.DevNonce[:]
	}

	appNonce, err := getAppNonce()
	if err != nil {
		log.Errorf("get AppNone error: %s", err)
		return nil, err
	}
	nwkSKey, err := getPlinkNwkSKey(AppKey, netID, appNonce, payload.DevNonce)
	if err != nil {
		log.Errorf("get NwkSkey error: %s", err)
		return nil, err
	}
	phyPayload := &PHYPayload{
		MHDR: MHDR{MType: MType(JoinAccept), MLen: 0x10},
		Payload: &JoinAcceptPayload{
			NetID:    netID,
			DevEUI:   payload.DevEUI,
			AppNonce: appNonce,
			Pading:   []byte{2, 2},
		},
	}
	if err := phyPayload.EncryptJoinAcceptPayload(AppKey); err != nil {
		log.Errorf("encryptAesEcbJoinAcceptPayload error: %s", err)
		return nil, err
	}
	if err = phyPayload.SetMIC(AppKey); err != nil {
		log.Error("err cmac.CalculateJoinAcceptMIC,", err)
		return nil, err
	}
	nodeUp := &storage.Node{Nid: node.Nid, NwkSKey: nwkSKey[:], UsedDevNonces: payload.DevNonce[:]}
	if err := nodeUp.UpdateNodeJoinRequest(); err != nil {
		log.Error("update node error")
		return nil, err
	}
	phyData, err := phyPayload.MarshalBinary()
	if err != nil {
		fmt.Println("MarshalBinary phyPayload error")
		return nil, err
	}
	return phyData, nil
}

type BindPayload struct {
	Type   [1]byte
	Len    [1]byte
	UserID [4]byte
	DevEUI common.EUI64
	Pading [4]byte
	MIC    [4]byte
}

func (bind BindPayload) MarshalBinary() (data []byte, err error) {
	fmt.Println("-------------------", bind)
	fmt.Println(bind.Type)
	out := make([]byte, 0, 22)
	out = append(out, bind.Type[0])
	out = append(out, bind.Len[0])
	for _, v := range bind.UserID[0:4] {
		out = append(out, v)
	}
	devEUI, err := bind.DevEUI.MarshalBinary()
	if err != nil {
		return out, err
	}
	out = append(out, devEUI...)
	out = append(out, bind.Pading[:]...)
	out = append(out, bind.MIC[:]...)
	return out, nil
}
func (bind *BindPayload) UnmarshalBinary(data []byte) error {
	if len(data) != 22 {
		return errors.New("data is not 22 byte")
	}
	bind.Type[0] = data[0]
	bind.Len[0] = data[1]
	for i, v := range data[2:6] {
		bind.UserID[i] = v
	}
	for i, v := range data[6:14] {
		bind.DevEUI[i] = v
	}
	for i, v := range data[14:18] {
		bind.Pading[i] = v
	}
	for i, v := range data[18:22] {
		bind.MIC[i] = v
	}
	return nil
}
func (bind BindPayload) HandlerDataPayload(node storage.Node) ([]byte, error) {
	log.Info("===============bind==================", bind)
	in, err := bind.MarshalBinary()
	out := make([]byte, 0)
	if err != nil {
		return in, err
	}
	NwsKey := common.AES128Key{}
	NwsKey.Scan(node.NwkSKey)
	if b, err := ValidateMIC(NwsKey, in); err != nil || !b {
		return in, errors.New("validata MIC err")
	}
	din, err := DecryptDataPayload(NwsKey, in[2:18])
	if err != nil {
		log.Error("decrypt data error", err)
		return out, err
	}
	log.Info("DecryptDataPayload:", din)
	copy(in[2:18], din[:])
	if err := bind.UnmarshalBinary(in); err != nil {
		log.Error("UnmarshalBinary ：", err)
		return out, err
	}
	var result []byte = []byte{Fail, SystemError}
	UserID := binary.BigEndian.Uint32(bind.UserID[:])
	log.Info("UserId is :", UserID)
	log.Info("NodeId is ;", node.Nid)
	//b := data.Bind{Nid:node.Nid}
	//if binds,err := b.GetByNid();err == nil {
	//	if len(binds) == 0 {
	//		b.Uid = int32(UserID)
	//		b.Share = 0
	//		if err := b.Add();err != nil {
	//			result = []byte{Fail,SystemError}
	//		}else{
	//			result = []byte{Successful, FailCode}
	//		}
	//	}
	//	for _,v := range binds{
	//		if v.Share == 1 {
	//			b.Uid = int32(UserID)
	//			b.Share = 2
	//			if err := b.Add();err != nil {
	//				result = []byte{Fail,SystemError}
	//			}else{
	//				result = []byte{Successful, FailCode}
	//			}
	//			break
	//		}else {
	//			result = []byte{Fail, AlreadyBind}
	//		}
	//	}
	//}else{
	//	result = []byte{Fail,SystemError}
	//}
	out = append(out, bind.Type[0])
	out = append(out, LEN)
	out = append(out, result...)
	mic, err := CalculateMIC(NwsKey, out)
	if err != nil {
		return out, err
	}
	out = append(out, mic...)
	return out, nil
}

//设备固件信息上报/返回
type ReportPayload struct {
	Type [1]byte
	Len  [1]byte
	//主控板硬件版本Master board hardware version
	MBHV [2]byte
	//主控板软件版本 Master board software version
	MBSV [2]byte
	//主控板生产日期
	MPTime [2]byte
	//通信模块硬件版本 ommunication board hardware version
	CBHV [2]byte
	//通信模块软件版本 ommunication board software version
	CBSV [2]byte
	//通信模块生产日期
	CPTime [2]byte
	Pading [4]byte
	MIC    [4]byte
}

func (report ReportPayload) MarshalBinary() (data []byte, err error) {
	out := make([]byte, 0)
	out = append(out, report.Type[0])
	out = append(out, report.Len[0])
	out = append(out, report.MBHV[:]...)
	out = append(out, report.MBSV[:]...)
	out = append(out, report.MPTime[:]...)
	out = append(out, report.CBHV[:]...)
	out = append(out, report.CBSV[:]...)
	out = append(out, report.CPTime[:]...)
	out = append(out, report.Pading[:]...)
	out = append(out, report.MIC[:]...)
	return out, nil
}
func (report *ReportPayload) UnmarshalBinary(data []byte) error {
	if len(data) != 22 {
		return errors.New("report data must 16 byte")
	}
	report.Type[0] = data[0]
	report.Len[0] = data[1]
	report.MBHV[0] = data[2]
	report.MBHV[1] = data[3]
	report.MBSV[0] = data[4]
	report.MBSV[1] = data[5]
	report.MPTime[0] = data[6]
	report.MPTime[1] = data[7]
	report.CBHV[0] = data[8]
	report.CBHV[1] = data[9]
	report.CBSV[0] = data[10]
	report.CBSV[1] = data[11]
	report.CPTime[0] = data[12]
	report.CPTime[1] = data[13]
	for i, v := range data[14:18] {
		report.Pading[i] = v
	}
	for i, v := range data[18:22] {
		report.MIC[i] = v
	}

	return nil
}
func (report ReportPayload) HandlerDataPayload(node storage.Node) ([]byte, error) {
	log.Info("================ReportPayload:============", report)
	go func() {
		var out []byte
		in, err := report.MarshalBinary()
		if err != nil {
			log.Error(err)
		}
		NwsKey := common.AES128Key{}
		NwsKey.Scan(node.NwkSKey)
		if b, err := ValidateMIC(NwsKey, in); err != nil || !b {
			log.Error("validata MIC err")
		}
		if out, err = DecryptDataPayload(NwsKey, in[2:len(in)-4]); err != nil {
			log.Error(err)
		}
		fmt.Println(out)
	}()
	fmt.Println("-----------------ok-------------------")
	return nil, nil
}

//固件升级指令接收确认
type UOAPayload struct {
	Type [1]byte
	Len  [1]byte
	ACK  [1]byte
	MIC  [4]byte
}

func (uoa UOAPayload) MarshalBinary() (data []byte, err error) {
	out := make([]byte, 0, 7)
	out = append(out, uoa.Type[:]...)
	out = append(out, uoa.Len[:]...)
	out = append(out, uoa.ACK[:]...)
	out = append(out, uoa.MIC[:]...)
	return out, nil
}
func (uoa *UOAPayload) UnmarshalBinary(data []byte) error {
	if len(data) != 7 {
		return errors.New("report data must 16 byte")
	}
	uoa.Type[0] = data[0]
	uoa.Len[0] = data[1]
	uoa.ACK[0] = data[2]
	for i, v := range data[3:7] {
		uoa.MIC[i] = v
	}
	return nil
}
func (uoa UOAPayload) HandlerDataPayload(node storage.Node) ([]byte, error) {
	log.Info("===========upgrade Result ack payload :=========", uoa)
	go func() {
		in, err := uoa.MarshalBinary()
		if err != nil {
			log.Error(err)
		}
		NwsKey := common.AES128Key{}
		NwsKey.Scan(node.NwkSKey)
		if b, err := ValidateMIC(NwsKey, in); err != nil || !b {
			log.Error("validata MIC err")
		}
		//反馈页面
		fmt.Println(in)
	}()
	fmt.Println("-----------------ok-------------------")
	return nil, nil
}

//通信模块固件升级结果反馈
type CMUAPayload struct {
	Type     [1]byte
	Len      [1]byte
	Result   [1]byte
	FailCode [1]byte
	MIC      [4]byte
}

func (cmua CMUAPayload) MarshalBinary() (data []byte, err error) {
	out := make([]byte, 0, 8)
	out = append(out, cmua.Type[:]...)
	out = append(out, cmua.Len[:]...)
	out = append(out, cmua.Result[:]...)
	out = append(out, cmua.FailCode[:]...)
	out = append(out, cmua.MIC[:]...)
	return out, nil
}
func (cmua *CMUAPayload) UnmarshalBinary(data []byte) error {
	if len(data) != 8 {
		return errors.New("report data must 16 byte")
	}
	cmua.Type[0] = data[0]
	cmua.Len[0] = data[1]
	cmua.Result[0] = data[2]
	cmua.FailCode[0] = data[2]
	for i, v := range data[3:7] {
		cmua.MIC[i] = v
	}
	return nil
}
func (cmua CMUAPayload) HandlerDataPayload(node storage.Node) ([]byte, error) {
	log.Info("===========CMUAPayload payload :============", cmua)
	go func() {
		in, err := cmua.MarshalBinary()
		if err != nil {
			log.Error(err)
		}
		NwsKey := common.AES128Key{}
		NwsKey.Scan(node.NwkSKey)
		if b, err := ValidateMIC(NwsKey, in); err != nil || !b {
			log.Error("validata MIC err")
		}
		//todo 反馈页面
		fmt.Println(in)
	}()
	fmt.Println("-----------------ok-------------------")
	return nil, nil
}

//主控板固件升级结果反馈
type BMUAPayload struct {
	Type     [1]byte
	Len      [1]byte
	Result   [1]byte
	FailCode [1]byte
	MIC      [4]byte
}

func (bmua BMUAPayload) MarshalBinary() (data []byte, err error) {
	out := make([]byte, 0, 8)
	out = append(out, bmua.Type[:]...)
	out = append(out, bmua.Len[:]...)
	out = append(out, bmua.Result[:]...)
	out = append(out, bmua.FailCode[:]...)
	out = append(out, bmua.MIC[:]...)
	return out, nil
}
func (bmua *BMUAPayload) UnmarshalBinary(data []byte) error {
	if len(data) != 8 {
		return errors.New("report data must 16 byte")
	}
	bmua.Type[0] = data[0]
	bmua.Len[0] = data[1]
	bmua.Result[0] = data[2]
	bmua.FailCode[0] = data[2]
	for i, v := range data[3:7] {
		bmua.MIC[i] = v
	}
	return nil
}
func (bmua BMUAPayload) HandlerDataPayload(node storage.Node) ([]byte, error) {
	log.Info("===========BMUAPayload payload :============", bmua)
	go func() {
		in, err := bmua.MarshalBinary()
		if err != nil {
			log.Error(err)
		}
		NwsKey := common.AES128Key{}
		NwsKey.Scan(node.NwkSKey)
		if b, err := ValidateMIC(NwsKey, in); err != nil || !b {
			log.Error("validata MIC err")
		}
		//todo 反馈页面
		fmt.Println(in)

	}()
	fmt.Println("-----------------ok-------------------")
	return nil, nil
}

//功能指令响应
type OrderRespPayload struct {
	Type [1]byte
	Len  [1]byte
	Data [15]byte
	CRC  [1]byte
}

func (os OrderRespPayload) MarshalBinary() (data []byte, err error) {
	out := make([]byte, 0, 8)
	out = append(out, os.Type[:]...)
	out = append(out, os.Len[:]...)
	out = append(out, os.Data[:]...)
	out = append(out, os.CRC[0])
	return out, nil
}
func (os *OrderRespPayload) UnmarshalBinary(data []byte) error {
	if len(data) != 18 {
		return errors.New("report data must 18 byte")
	}
	os.Type[0] = data[0]
	os.Len[0] = data[1]
	for i, v := range data[2:17] {
		os.Data[i] = v
	}
	os.CRC[0] = data[17]
	return nil
}
func (os OrderRespPayload) HandlerDataPayload(node storage.Node) ([]byte, error) {
	log.Info("===========OrderRespPayload payload :============", os)
	go func() {
		in, err := os.MarshalBinary()
		if err != nil {
			log.Error(err)
			return
		}
		cic := CalculateCRC(in[0 : len(in)-1])
		if cic != in[len(in)-1] {
			log.Error("cic is different", cic, "in:", in[len(in)-1])
		} else {
			order := OrderResponse{}
			if err := order.UnMarshalBinary(in[0 : len(in)-1]); err != nil {
				log.Error(err)
				return
			}
			client := module.RedisClient(module.Persistence).Get()
			defer client.Close()
			value, err := order.MarshalJson()
			if err != nil {
				log.Error(err)
				return
			}
			if _, err := client.Do("hset", module.AirInfoRedisKey, node.Nid, value); err != nil {
				log.Error(err)
				return
			}
			log.Info("--------OrderRespPayload: ---------", in)
		}
	}()
	fmt.Println("-----------------ok-------------------")
	return nil, nil
}
