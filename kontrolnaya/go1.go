package main

import "fmt"

type PepeShneine struct {
	Speed    int
	Charisma int
	Wisdom   int
}

func NewPepeShneine(speed, charisma, wisdom int) *PepeShneine {
	return &PepeShneine{speed, charisma, wisdom}
}

func (p *PepeShneine) GetRating() int {
	return p.Speed*2 + p.Charisma*3 + p.Wisdom
}

func main() {
	p1 := NewPepeShneine(5, 7, 6)
	p2 := NewPepeShneine(8, 4, 5)

	fmt.Println("Пепе Шнейне [Скорость:", p1.Speed, ", Харизма:", p1.Charisma, ", Мудрость:", p1.Wisdom, "] | Рейтинг:", p1.GetRating())
	fmt.Println("Пепе Шнейне [Скорость:", p2.Speed, ", Харизма:", p2.Charisma, ", Мудрость:", p2.Wisdom, "] | Рейтинг:", p2.GetRating())

	if p1.GetRating() > p2.GetRating() {
		fmt.Println("Первый Пепе сильнее")
	} else {
		fmt.Println("Второй Пепе сильнее")
	}
}