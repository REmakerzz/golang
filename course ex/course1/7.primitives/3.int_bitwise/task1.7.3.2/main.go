package main

import (
	"fmt"
	"strconv"
)

func main() {
    num,err:= strconv.ParseInt("700",8,0)
    fmt.Println(num,err)
	fmt.Println(getFilePermissions(644))
}

func getFilePermissions(flag int) string {

    str := strconv.Itoa(flag)
    num, err := strconv.ParseInt(str,8,0)
    if err != nil {
        panic(err)
    }
    flag = int(num)
	owner := []string{"-", "-", "-"}
	group := []string{"-", "-", "-"}
	other := []string{"-", "-", "-"}

	if flag&0400 != 0 {
		owner[0] = "Read"
	}
	if flag&0200 != 0 {
		owner[1] = "Write"
	}
	if flag&0100 != 0 {
		owner[2] = "Execute"
	}

	if flag&040 != 0 {
		group[0] = "Read"
	}
	if flag&020 != 0 {
		group[1] = "Write"
	}
	if flag&010 != 0 {
		group[2] = "Execute"
	}

	if flag&04 != 0 {
		other[0] = "Read"
	}
	if flag&02 != 0 {
		other[1] = "Write"
	}
	if flag&01 != 0 {
		other[2] = "Execute"
	}

	return fmt.Sprintf("Owner:%s,%s,%s Group:%s,%s,%s Other:%s,%s,%s",
		owner[0], owner[1], owner[2],
		group[0], group[1], group[2],
		other[0], other[1], other[2])
}
