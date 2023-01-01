package main

/*
#include "window/application.h"
#include "window/window.h"
*/
import "C"

func main() {
	name := C.CString("School21")
	w := C.Window_Create(100, 100, 300, 200, name)
}
