package handler

import (
	"Go-Grasscutter/generated/pb"
	"Go-Grasscutter/log"
	"Go-Grasscutter/server/kcp/packet/base"
	"Go-Grasscutter/server/kcp/session"
	"google.golang.org/protobuf/proto"
)

// SetPlayerBornDataReq
func init() {
	session.Router[base.SetPlayerBornDataReq] = HandlerSetPlayerBornDataReq
}

// HandlerSetPlayerBornDataReq 处理SetPlayerBornDataReq协议的函数
func HandlerSetPlayerBornDataReq(sess *session.Session, header, payload []byte) {
	req := &pb.SetPlayerBornDataReq{
		// AvatarId: 10000007,
	}
	err := proto.Unmarshal(header, req)
	if err != nil {
		log.SugaredLogger.Error(err)
		return
	}

	player := sess.Player
	player.Nickname = req.GetNickName()
	// todo AvatarId

	player.OnLogin()

	// 发送Born响应包
	sess.Send(base.NewPacketWithCode(base.SetPlayerBornDataRsp))
}
