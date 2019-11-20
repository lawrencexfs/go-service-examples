package plr

import (
	"battleservice/src/services/battle/usercmd"
	"battleservice/src/services/servicetype"

	"github.com/giant-tech/go-service/base/serializer"
	"github.com/giant-tech/go-service/framework/iserver"
)

// AOIUpdate AOI更新
func (s *ScenePlayer) AOIUpdate(aois []iserver.AOIInfo) {
	msg := usercmd.NewEntityAOISMsg()
	for i := 0; i < len(aois); i++ {
		if msg.Num >= 20 {
			s.AsyncCall(servicetype.ServiceTypeClient, "AOI", msg)
			msg = usercmd.NewEntityAOISMsg()
		}

		info := aois[i]

		var data []byte

		if info.IsEnter {
			num, propBytes := s.PackProps(uint32(servicetype.ServiceTypeClient))
			m := &usercmd.EnterAOI{
				EntityID:   s.GetID(),
				EntityType: s.GetType(),
				PropNum:    uint16(num),
				Properties: propBytes,
			}

			data = make([]byte, serializer.GetSizeNew(m)+1)
			data[0] = 1
			serializer.SerializeNewWithBuff(data[1:], m)

		} else {
			m := &usercmd.LeaveAOI{
				EntityID: s.GetID(),
			}

			data = make([]byte, serializer.GetSizeNew(m)+1)
			data[0] = 1
			serializer.SerializeNewWithBuff(data[1:], m)
		}

		msg.AddData(data)
	}

	if msg.Num > 0 {
		s.AsyncCall(servicetype.ServiceTypeClient, "AOI", msg)
	}
}
