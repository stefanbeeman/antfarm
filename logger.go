package antfarm

import (
	"fmt"
	"os"
	"strconv"
)

func (this World) log(entry string) {
	written := strconv.Itoa(this.Now) + ": "
	written = written + entry
	f, openErr := os.OpenFile("antfarm.txt", os.O_RDWR|os.O_APPEND, 0660)
	if openErr != nil {
		fmt.Println(openErr)
	}
	_, writeErr := f.WriteString(written + "\n")
	if writeErr != nil {
		fmt.Println(writeErr)
	}
	f.Close()
	fmt.Println(written)
}
