package main

//import "C"
import "time"
import "physx-demo/GiantPhysXGo"

func main() {
	//var physics = GiantPhysXGo.GxCreatePhysics("TropicalStorm.gxgame", "192.168.133.79")
	var str1 = "TropicalStorm.gxgame"
	var str2 = "192.168.133.79"

	//cstr1, cstr2 := C.CString(str1), C.CString(str2)
	var physics = GiantPhysXGo.GxCreatePhysics(str1, str2)

	//defer C.free(unsafe.Pointer(cstr1)) // must call
	//defer C.free(unsafe.Pointer(cstr2))

	var scene = physics.CreateScene("factory01_area04_01.gxscene")

	for {
		scene.Update()
		time.Sleep(1 * time.Second)
	}

	GiantPhysXGo.GxDestroyPhysics()

	//C.GxCreatePhysics(cstr1, cstr2)
}
