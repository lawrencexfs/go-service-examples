package plr

import (
	"github.com/giant-tech/go-service/framework/iserver"
)

// AOIUpdate AOI更新
func (s *ScenePlayer) AOIUpdate(aois []iserver.AOIInfo) {
	// msg := msgdef.NewEntityAOISMsg()
	// for i := 0; i < len(e.aoies); i++ {

	// 	if msg.Num >= 20 {
	// 		e.PostToClient(msg)
	// 		msg = msgdef.NewEntityAOISMsg()
	// 	}

	// 	info := e.aoies[i]

	// 	ip := info.Entity.(iAOIPacker)

	// 	var data []byte

	// 	if info.IsEnter {
	// 		num, propBytes := ip.GetAOIProp()
	// 		m := &msgdef.EnterAOI{
	// 			EntityID:   ip.GetID(),
	// 			EntityType: ip.GetType(),
	// 			State:      ip.GetStatePack(),
	// 			PropNum:    uint16(num),
	// 			Properties: propBytes,
	// 			//BaseProps:  ip.GetBaseProps(),
	// 		}

	// 		data = make([]byte, m.Size()+1)
	// 		data[0] = 1
	// 		m.MarshalTo(data[1:])

	// 	} else {
	// 		m := &msgdef.LeaveAOI{
	// 			EntityID: ip.GetID(),
	// 		}

	// 		data = make([]byte, m.Size()+1)
	// 		data[0] = 0
	// 		m.MarshalTo(data[1:])
	// 	}

	// 	msg.AddData(data)
	// }

	// e.PostToClient(msg)
}
