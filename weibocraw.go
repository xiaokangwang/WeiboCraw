package main

import "os"

func main() {

	dirarg := os.Args

	switch dirarg[1] {
	case "dba":
		dbafc(dirarg[2:])
	}

}
