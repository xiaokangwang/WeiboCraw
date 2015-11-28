package main

func Crawall() {
	var theExecTimeDbS ExecTimeDbS
	theExecTimeDbS.Boot()
	conf := theExecTimeDbS.LoadConfigure()
	instuuid := StyleMkcrawinst(conf["crawas"], conf["crawua"], conf["crawnote"], conf["crawcookie"])
	crawTaskExec(instuuid, &theExecTimeDbS)
}

func Addcrawtargetx(uid) {
	var theExecTimeDbS ExecTimeDbS
	theExecTimeDbS.Boot()
	conf := theExecTimeDbS.LoadConfigure()
	Addcrawtarget(uid, dbe)
}

func Crawuid(uid string) {
	var theExecTimeDbS ExecTimeDbS
	theExecTimeDbS.Boot()
	conf := theExecTimeDbS.LoadConfigure()
	instuuid := StyleMkcrawinst(conf["crawas"], conf["crawua"], conf["crawnote"], conf["crawcookie"])
	Docraw(uid, instuuid, &theExecTimeDbS)
}
