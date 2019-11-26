package internal

// 附加压迫力

import (
	"github.com/cihub/seelog"
	"github.com/giant-tech/go-service/base/linmath"
)

//附加力
type AddonForceData struct {
	leftTime uint64
	force    linmath.Vector3
}

type Force struct {
	addonForceDatas []*AddonForceData //附加力数据
	currForce       linmath.Vector3   //附加力
}

func (this *Force) ClearForce() {
	this.currForce.X = 0
	this.currForce.Y = 0
	this.addonForceDatas = this.addonForceDatas[:0]
}

func (this *Force) HasForce() bool {
	return len(this.addonForceDatas) != 0
}

func (this *Force) AddForce(force linmath.Vector3, time uint64) {
	data := &AddonForceData{force: force, leftTime: time}
	this.addonForceDatas = append(this.addonForceDatas, data)
}

func (this *Force) UpdateForce(detaTime float64) {
	this.currForce.X = 0
	this.currForce.Y = 0

	if len(this.addonForceDatas) < 1 {
		return
	}

	var tempList []*AddonForceData

	for _, data := range this.addonForceDatas {
		if data.leftTime > 0 {
			data.leftTime--
			this.currForce.Add(data.force)
			tempList = append(tempList, data)
		}
	}

	if len(tempList) != len(this.addonForceDatas) {
		this.addonForceDatas = tempList
	}

	if len(this.addonForceDatas) > 100 {
		seelog.Error("BIG ERROR,addonForceDatas overflow:", len(this.addonForceDatas))
		this.ClearForce()
	}
}
func (this *Force) GetForce() *linmath.Vector3 {
	return &this.currForce
}
