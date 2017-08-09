package main

import (
	"flag"
	"parking_lot/parking_lot_pkg"
)

var file_name = flag.String("input_file_name", "", "input_file_name")

func main() {
	flag.Parse()
	parking_lot_pkg.Start_fn(*file_name)
	return
}
