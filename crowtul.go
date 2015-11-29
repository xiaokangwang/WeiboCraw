package main

import "fmt"

func Crawall() {
	var theExecTimeDbS ExecTimeDbS
	theExecTimeDbS.Boot()
	conf, err := theExecTimeDbS.LoadConfigure()
	if err != nil {
		fmt.Println(err)
	}
	instuuid, err := StyleMkcrawinst((*conf)["crawas"], (*conf)["crawua"], (*conf)["crawnote"], (*conf)["crawcookie"], &theExecTimeDbS)
	if err != nil {
		fmt.Println(err)
	}
	crawTaskExec(instuuid, &theExecTimeDbS)
}

func Addcrawtargetx(uid string) {
	var theExecTimeDbS ExecTimeDbS
	theExecTimeDbS.Boot()
	theExecTimeDbS.LoadConfigure()
	Addcrawtarget(uid, &theExecTimeDbS)
}

func Crawuid(uid string) {
	var theExecTimeDbS ExecTimeDbS
	theExecTimeDbS.Boot()
	conf, _ := theExecTimeDbS.LoadConfigure()
	instuuid, _ := StyleMkcrawinst((*conf)["crawas"], (*conf)["crawua"], (*conf)["crawnote"], (*conf)["crawcookie"], &theExecTimeDbS)
	Docraw(uid, instuuid, &theExecTimeDbS)
}
