package controller

import (
	"net/http"
	"github.com/gorilla/websocket"
	"fmt"
	log "github.com/cihub/seelog"
	"cotx-http/po"
	"cotx-http/rpcClient"
	"cotx-http/pb"
	"github.com/golibs/uuid"
	"encoding/json"
	"cotx-http/redis"
	"time"
)

var WebsocketChan po.WebscoketChan
var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func GetWebsocketChan() po.WebscoketChan {
	return WebsocketChan
}


func WebsocketHandle(w http.ResponseWriter,r  *http.Request) {
   log.Info("Websocket:start handle webscoket")
   //获取客户端的连接
   conn,err := Upgrader.Upgrade(w,r,nil)
	if err != nil {
		fmt.Println("Websocket:获取客户端连接错误")
		return 
	}
	//将conn与唯一表示做对应
	var id uuid.UUID  = uuid.Rand()
	websocketId := id.Hex()
	go GetPushMessage(websocketId)
	go PushHandle(conn,websocketId)
	go PullHandle(conn,websocketId)

}

func PullHandle(conn *websocket.Conn,websocketid string)  {
	for   {
		_,d,err := conn.ReadMessage()
		data := po.Websocket{}
		json.Unmarshal(d,&data)
		log.Info("data===",data)
		if err != nil {
			log.Error("Webscoket:读取客户端信息失败:  ",err)
			break
		}
		switch data.Code {
		case 100:
			log.Info("Websocket:获取心跳：",data.Result)
			/*进行数据解析和数据写入操作*/
			heart,ok := data.Result.(map[string]interface{})
			token,_ := heart["token"].(string)
			source,_ := heart["source"].(string)
			if !ok {
				log.Error("Webscoket:心跳类型判断错误")
				continue
			}
			/*根据token 获取session 用户缓存的信息*/
			reply := rpcClient.GetUserInfo(&pb.SsoRequest{Source:source, SessionName:token})
			reply.SessionName = token
			log.Info("Check Token Reply:  ",reply)
			if reply.ErrorCode != 10000 {
				log.Error("Token:", token, " Check  Failed, ", reply.ErrorCode)
				return
			}
			webscoketinfo := &po.WebScoketInfo{
				WebsocketId:websocketid,
				UserId:reply.Uid,
				Token:token,
				Conn:conn,
			}
			// json 化
			data ,err:= json.Marshal(webscoketinfo)
			if err != nil {
				log.Error("Json:json失败",err)
				continue
			}
			/*数据写入redis*/
			err = redis.WriteToRedis(websocketid,data)
			if err != nil {
				log.Error("Redis:数据写入redis失败",err)
				continue
			}
			break
		default:
			log.Error("Webscoket:code未定义",data.Code)
			conn.Close()
			log.Info("websocket/接收数据的协程生命周期结束，将退出该协程...")
			return
		}
	}
	redis.DeleteFromRedis(websocketid)
	log.Info("websocket/读进程结束...")
}

func PushHandle(conn *websocket.Conn,websocketid string)  {
	log.Info("Webscoket:向客户端推送信息的进程",conn)
	for  messagemap := range  GetWebsocketChan().Websocketchan {
		log.Info("messagemap==:",messagemap)
		if messagemap[websocketid].Conn == nil{
			 log.Info("websocket/客户端推送信息的进程生命周期结束，将退出.....")
			 return
		}
         message :=  messagemap[websocketid]
         var pushwebsocket po.Websocket
         var result    po.WebscoketResult
		if len(message.Message)== 0 {
			continue
		}
            m := message.Message[0]
			pushwebsocket.Code = 200
			pushwebsocket.Msg = "系统通知"
			result.Code =  m.Code
			result.State = m.State
			result.Describe =  m.Describe
			result.GatewayId = m.GayewayId
			result.Value = m.Value
			pushwebsocket.Result = result
			data ,err := json.Marshal(pushwebsocket)
			if err != nil {
				log.Error("WEBSOCKET/json解析失败")
				continue
			}
			err = conn.WriteMessage(1,data)
			if err != nil {
				log.Error("websocket/推送数据失败")
				continue
			}
	}
}

func GetPushMessage(websocketid string)  {
     /*从redis中读取信息*/
	for  {
		time.Sleep(time.Second * 30)
		websocketchan := GetWebsocketChan()
		websocketmessgage := po.WebSocketMessgage{}
		data,err := redis.ReadFromRedis(websocketid)
		if err != nil {
			websocketchan.Websocketmap[websocketid] = websocketmessgage
			websocketchan.Websocketchan <- websocketchan.Websocketmap
			log.Info("Redis/没有对应的设备接入，将退出协程.....")
			return
		}
		if data == nil {
			websocketchan.Websocketmap[websocketid] = websocketmessgage
			websocketchan.Websocketchan <- websocketchan.Websocketmap
			log.Info("Redis/没有对应的设备接入，将退出协程.....")
			return
		}
		value,ok := data.([]byte)
		if !ok {
			log.Infof("Redis:data类型判断错误")
			continue
		}
		websocketinfo :=  po.WebScoketInfo{}
		json.Unmarshal(value,&websocketinfo)
		//rpc 调用
		req := &pb.ReqShadow{UserId:websocketinfo.UserId}
		res,err := rpcClient.GetPushMessage(req)
		if err != nil {
			log.Error("Shadow:调用获取pushmessage错误")
			continue
		}
		if len(res.Message) != 0 {
			websocketmessgage.Conn = websocketinfo.Conn
			websocketmessgage.Message = res.Message
			websocketchan.Websocketmap[websocketid] = websocketmessgage
			websocketchan.Websocketchan <-  websocketchan.Websocketmap
		}
	}
}