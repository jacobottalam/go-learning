package dbmodule

import "fmt"

type Person struct {
	Per_Name string
	}

func (p *Person) PrintDetails(){
	fmt.Println(p.Per_Name)
}
