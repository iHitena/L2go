package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern


Променимость:
Когда вам нужно выполнить какую-то операцию над всеми элементами сложной структуры объектов, например, деревом.
 Посетитель позволяет применять одну и ту же операцию к объектам различных классов.
 Когда над объектами сложной структуры объектов надо выполнять некоторые не связанные между собой операции,
 но вы не хотите «засорять» классы такими операциями.
 Посетитель позволяет извлечь родственные операции из классов, составляющих структуру объектов,
  поместив их в один класс-посетитель. Если структура объектов является общей для нескольких приложений,
   то паттерн позволит в каждое приложение включить только нужные операции.
 Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии.
 Посетитель позволяет определить поведение только для этих классов, оставив его пустым для всех остальных.

	плюсы:
Упрощает добавление операций, работающих со сложными структурами объектов.
 Объединяет родственные операции в одном классе.
 Посетитель может накапливать состояние при обходе структуры элементов.
минусы:
 Паттерн не оправдан, если иерархия элементов часто меняется.
 Может привести к нарушению инкапсуляции элементов.
*/

type Visitor interface {
	VisitorForCar(*Car)
	VisitorForMotorbike(*Motorbike)
}

type Car struct {
	GasTankLeft  int
	GasTankRight int
}

func (c *Car) Accept(v Visitor) {
	v.VisitorForCar(c)
}

type Motorbike struct {
	GasTank int
}

func (m *Motorbike) Accept(v Visitor) {
	v.VisitorForMotorbike(m)
}

type Travel struct {
	distance int
}

func (t *Travel) VisitorForCar(c *Car) {

	t.distance = c.GasTankLeft*10 + c.GasTankRight*10
	fmt.Println(t.distance)
}

func (t *Travel) VisitorForMotorbike(b *Motorbike) {
	t.distance = b.GasTank * 10
	fmt.Println(t.distance)
}

type SumGas struct {
	sum int
}

func (t *SumGas) VisitorForCar(c *Car) {

	t.sum = c.GasTankLeft + c.GasTankRight
	fmt.Println(t.sum)
}

func (t *SumGas) VisitorForMotorbike(b *Motorbike) {
	t.sum = b.GasTank
	fmt.Println(t.sum)

}
