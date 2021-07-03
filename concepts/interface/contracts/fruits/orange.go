package fruits

import "fmt"

type Orange struct {
}

func (o Orange) Printname() {
	fmt.Println("Orange")
}
