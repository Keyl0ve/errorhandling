package a

import "fmt"

func ok() {
	{
		_, err := addNum("test")
		if err != nil { //want "Separated notation"
			fmt.Printf("%v", err)
		}
	}
	{
		_, err := addNum("test")
		if err != nil { //want "Separated notation"
			fmt.Printf("%v", err)
		}
	}
}

func addNum(num int) (int, error) {
	if num == 0 {
		return 0, fmt.Errorf("err:")
	}
	return num, nil
}
