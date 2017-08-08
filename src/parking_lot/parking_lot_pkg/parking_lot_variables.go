package parking_lot_pkg

type slot struct {
	vechicle_no string
	color       string
}

var (
	parking_slot_info     = make(map[int]slot)
	park_number_slot_info = make(map[string]int)
	n                     int
	slot_remaining        []int
)
