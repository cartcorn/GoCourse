// 2. Создать тип, описывающий контакт в телефонной книге. Создать псевдоним типа телефонной
// книги (массив контактов) и реализовать для него интерфейс Sort{}

package main

import (
	"fmt"
	"sort"
)

type Phones struct {
	Name   string
	Number int
}

func (p Phones) ViewList() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Number)
}

type MySort []Phones

func (a MySort) Len() int           { return len(a) }
func (a MySort) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MySort) Less(i, j int) bool { return a[i].Number < a[j].Number }

func main() {
	adressBook := []Phones{
		{"Миша", 89458573344},
		{"Никита", 87356764400},
		{"Женя", 89652343311},
		{"Юля", 89256770087},
	}

	fmt.Println(adressBook)
	sort.Sort(MySort(adressBook))
	fmt.Println(adressBook)

}
