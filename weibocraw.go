package main

import "os"

func main() {

	dirarg := os.Args

	switch dirarg[1] {
	case "dba":
		Dbafc(dirarg[2:])
	case "craw":
		switch dirarg[2] {
		case "commit":
			Crawall()
		case "tadd":
			Addcrawtargetx(dirarg[3])
		case "test_uid":
			Crawuid(dirarg[3])
		case "stat":
			Crawstat()
		}
	}

}
