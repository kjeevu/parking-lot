package parking_lot_pkg

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Process_fn(input_cmd string) (r_err error) {
	input_cmd = strings.TrimSuffix(input_cmd, "\n")
	cmd_arr := strings.Split(input_cmd, " ")

	switch cmd_arr[0] {
	case "create_parking_lot":

		n, r_err = strconv.Atoi(cmd_arr[1])

		if r_err != nil {
			fmt.Println(r_err)
			r_err = errors.New("Unable to Convert")
			return
		}

		slot_remaining = make([]int, n)
		for i := 0; i < n; i++ {
			slot_remaining[i] = i + 1
		}
		fmt.Println("Created a parking lot with", n, " slots")
		fmt.Println(len(slot_remaining))
	case "park":

		if len(slot_remaining) > 0 {
			slot_no := slot_remaining[0]
			slot_remaining = slot_remaining[:0+copy(slot_remaining[0:], slot_remaining[0+1:])]
			temp := slot{vechicle_no: cmd_arr[1], color: cmd_arr[2]}
			parking_slot_info[slot_no] = temp
			park_number_slot_info[cmd_arr[1]] = slot_no
			fmt.Println("Allocated slot number: ", slot_no)

		} else {
			fmt.Println("Sorry, parking lot is full")
		}
	case "leave":
		var slot_no int
		slot_no, r_err = strconv.Atoi(cmd_arr[1])
		if r_err != nil {
			fmt.Println(r_err)
			r_err = errors.New("Unable to Convert")
			return
		}
		temp_slot := parking_slot_info[slot_no]
		delete(park_number_slot_info, temp_slot.vechicle_no)
		delete(parking_slot_info, slot_no)

		fmt.Println("Slot no ", slot_no, "  free")
		slot_remaining = append(slot_remaining, slot_no)
	case "status":

		fmt.Println("Slot No   Registration No   Colour")
		for slot_no, slot_info := range parking_slot_info {
			fmt.Println(slot_no, "     ", slot_info.vechicle_no, "      ", slot_info.color)
		}
	case "slot_number_for_registration_number":
		l_slot_number, ok := park_number_slot_info[cmd_arr[1]]

		if !ok {
			fmt.Println("Not found")
		} else {
			fmt.Println(l_slot_number)
		}
	case "registration_numbers_for_cars_with_colour":
		l_flag := false
		for _, l_slot_info := range parking_slot_info {
			if strings.EqualFold(cmd_arr[1], l_slot_info.color) {
				l_flag = true
				fmt.Print(l_slot_info.vechicle_no, "  ")
			}
		}
		if !l_flag {
			fmt.Println("Not Found")
		} else {
			fmt.Println("")
		}
	case "slot_numbers_for_cars_with_colour":
		l_flag := false
		for l_slot_number, l_slot_info := range parking_slot_info {

			if strings.EqualFold(l_slot_info.color, cmd_arr[1]) {
				l_flag = true
				fmt.Print(l_slot_number, "  ")
			}
		}
		if !l_flag {
			fmt.Println("Not Found")
		} else {
			fmt.Println("")
		}
	}

	return
}
