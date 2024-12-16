package main

import "fmt"

type TVer interface {
	switchOFF() bool
	switchON() bool
	GetStatus() bool
	GetModel() string
}

type SamsungTV struct {
	status bool
	model  string
}
type LgTV struct {
	status bool
	model  string
}

func (tv *SamsungTV) switchOFF() bool {
	tv.status = false
	return true
}

func (tv *SamsungTV) switchON() bool {
	tv.status = true
	return true
}

func (tv *SamsungTV) GetStatus() bool {
	return tv.status
}

func (tv *SamsungTV) GetModel() string {
	return tv.model
}

func (tv *SamsungTV) SamsungHub() string {
	return "Samsung Hub"
}

func (tv *LgTV) switchOFF() bool {
	tv.status = false
	return true
}

func (tv *LgTV) switchON() bool {
	tv.status = true
	return true
}

func (tv *LgTV) GetStatus() bool {
	return tv.status
}

func (tv *LgTV) GetModel() string {
	return tv.model
}

func (tv *LgTV) LGHub() string {
	return "LG Hub"
}

func main() {
	samsungTV := &SamsungTV{
		status: true,
		model:  "Samsung XL-100500",
	}

	lgTV := &LgTV{
		status: true,
		model:  "LG Ultra 100500 HD",
	}

	fmt.Println(samsungTV.GetStatus())

	fmt.Println(samsungTV.GetModel())
	fmt.Println(samsungTV.switchOFF())
	fmt.Println(samsungTV.GetStatus())
	fmt.Println(samsungTV.switchON())
	fmt.Println(samsungTV.GetStatus())
	fmt.Println(samsungTV.SamsungHub())

	fmt.Println(lgTV.GetStatus())
	fmt.Println(lgTV.GetModel())
	fmt.Println(lgTV.switchOFF())
	fmt.Println(lgTV.GetStatus())
	fmt.Println(lgTV.switchON())
	fmt.Println(lgTV.GetStatus())
	fmt.Println(lgTV.LGHub())
}
