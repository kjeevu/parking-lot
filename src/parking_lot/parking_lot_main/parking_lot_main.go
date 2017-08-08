package main

import (
	"flag"
	"parkinglot/parking_lot_pkg"
)

var file_name = flag.String("file_name", "", "file_name")

func main() {
	flag.Parse()
	parking_lot_pkg.Start_fn(*file_name)
	return
}
