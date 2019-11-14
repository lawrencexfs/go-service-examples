package main

import (
	"physx-demo/GiantPhysXGo"
	"time"
)

func main() {
	var physics = GiantPhysXGo.GxCreatePhysics("TropicalStorm.gxgame", "192.168.133.79")
	var scene = physics.CreateScene("factory01_area04_01.gxscene")

	for {
		scene.Update()
		time.Sleep(1 * time.Second)
	}

	GiantPhysXGo.GxDestroyPhysics()
}
