package air

import (
	"account-domain-rpc/api/ns"
	"account-domain-rpc/common"
	"account-domain-rpc/rpc/client"
	"account-domain-rpc/storage"

	"context"
	"fmt"
)

type downLink struct {
	nsRpc  ns.NetworkServerClient
	devEUI string
	pararm byte
}

type DataDownPayLoad struct {
	Type [1]byte
	Len  [1]byte
	data [1]byte
	CRC  [1]byte
}

func (dp DataDownPayLoad) GetCRCdata() []byte {
	out := make([]byte, 0, 3)
	out = append(out, dp.Type[0])
	out = append(out, dp.Len[0])
	out = append(out, dp.data[0])
	return out
}
func (dp DataDownPayLoad) MarshalBinary() []byte {
	out := make([]byte, 0, 3)
	out = append(out, dp.Type[0])
	out = append(out, dp.Len[0])
	out = append(out, dp.data[0])
	out = append(out, dp.CRC[0])
	return out
}
func SendPackageToNS(Type byte, d *downLink) error {
	dp := DataDownPayLoad{}
	dp.Type[0] = Type
	dp.Len[0] = LEN
	dp.data[0] = d.pararm
	dp.CRC[0] = getCRC(dp.GetCRCdata())
	fmt.Println("downLink is .....", d)
	fmt.Println("downLink is .....", d.nsRpc)
	nsrpc := d.nsRpc
	resp, err := nsrpc.TransparentCmd(context.Background(), &ns.CmdRequest{PhyPayload: dp.MarshalBinary(), DevEUI: d.devEUI})
	fmt.Println("resp is ......", resp)
	fmt.Println("resp error .....", err)
	return err
}

const LEN byte = 1

func NewDownLink(devEUI string, pararm byte) *downLink {
	network := client.NsRpcClient()
	if network == nil {
		return nil
	}
	return &downLink{
		nsRpc:  network,
		devEUI: devEUI,
		pararm: pararm,
	}
}

//下行空气净化器指令
func (d *downLink) DownOrderAir(Type byte) error {
	return SendPackageToNS(Type, d)
}

type UpgradePayload struct {
	Type [1]byte
	Len  [1]byte
	Data [1]byte
	MIC  [4]byte
}

func (d *downLink) GetUpgrade(node *storage.Node) error {
	up := &UpgradePayload{}
	up.Type[0] = 0xF7
	up.Len[0] = 0x01
	up.Data[0] = 0x00
	micBytes := make([]byte, 0)
	micBytes = append(micBytes, up.Type[0])
	micBytes = append(micBytes, up.Len[0])
	micBytes = append(micBytes, up.Data[0])
	var appKey common.AES128Key
	appKey.Scan(node.NwkSKey)
	mic, err := CalculateMIC(appKey, micBytes)
	if err != nil {
		return err
	}
	copy(up.MIC[:], mic)
	return nil
}

//计算CRC
func getCRC(data []byte) byte {
	var cic int32 = 0
	for _, v := range data {
		cic += int32(v)
	}
	return byte(cic & 0x00000000000000ff)
}
