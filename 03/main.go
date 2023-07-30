package main

import "fmt"

type Cat interface {
	Miaow() string
}

type EvilCat interface {
	Attack()
}

type Clothes struct {
	Size int
}

type Tuanzi struct {
	Color  string
	Weight int
	Old    int
	// Cls    Clothes
	Clothes
}

func (tz *Tuanzi) Miaow() string {
	if tz.Old <= 5 {
		tz.Old++
		return "haihaihai"
	} else {
		return "hihihi"
	}
}

func (tz *Tuanzi) Attack() {
	print(tz.Clothes.Size)
}

type Human struct {
	Shoes
	Name   string
	Weight int
}

func (hu *Human) Miaow() string {
	return "*#%^@%!%&"
}

type Shoes struct {
	Colour string
}

// 任意叫
func AnyMiaow(c Cat) string {
	return c.Miaow()
}

func AddWeight(h *Human) {
	(*h).Weight++
}

func main() {
	tz := new(Tuanzi)
	cat := Cat(tz)
	ecat := EvilCat(tz)
	print(cat, ecat)

	ren := Human{
		Shoes: Shoes{
			Colour: "red",
		},
		Name:   "ren",
		Weight: 100,
	}

	tz2 := Tuanzi{
		Color:  "yellow",
		Weight: 5,
		Old:    1,
		Clothes: Clothes{
			Size: 12345678,
		},
	}

	AnyMiaow(&ren)
	AnyMiaow(&tz2)

	fmt.Println("-----------")
	fmt.Println(ren.Weight)
	AddWeight(&ren)
	fmt.Println(ren.Weight)
}
