package main

/*
#cgo CXXFLAGS: -std=c++11
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L ./bin -lGiantPhysXRelease_x64
#include <stdint.h>
#include <stdlib.h>
#include <string.h>
#include "GiantPhysX/GxAPI.h"
*/
import "C"

func main() {
	//var physics = GiantPhysXGo.GxCreatePhysics("TropicalStorm.gxgame", "192.168.133.79")
	//var str1 = "TropicalStorm.gxgame"
	//var str2 = "192.168.133.79"

	//cstr1, cstr2 := C.CString(str1), C.CString(str2)
	//var physics = GiantPhysXGo.GxCreatePhysics(cstr1, cstr2)

	//defer C.free(unsafe.Pointer(cstr1)) // must call
	//defer C.free(unsafe.Pointer(cstr2))

	//var scene = physics.CreateScene("factory01_area04_01.gxscene")

	/*for {
		scene.Update()
		time.Sleep(1 * time.Second)
	}
	*/
	//GiantPhysXGo.GxDestroyPhysics()

	//C.GxCreatePhysics(cstr1, cstr2)
}
