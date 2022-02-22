package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern


	Приминимость:
 Когда программа должна обрабатывать разнообразные запросы несколькими способами,
 но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
 С помощью Цепочки обязанностей вы можете связать потенциальных обработчиков в одну цепь и
  при получении запроса поочерёдно спрашивать каждого из них, не хочет ли он обработать запрос.
 Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
 Цепочка обязанностей позволяет запускать обработчиков последовательно один за другим в том порядке,
  в котором они находятся в цепочке.
 Когда набор объектов, способных обработать запрос, должен задаваться динамически.
 В любой момент вы можете вмешаться в существующую цепочку и переназначить связи так,
 чтобы убрать или добавить новое звено.

Приемущества:
Уменьшает зависимость между клиентом и обработчиками.
 Реализует принцип единственной обязанности.
 Реализует принцип открытости/закрытости.
Недостатки
 Запрос может остаться никем не обработанным.
*/

type Client struct {
	Name string
}

type HandlerPattern interface {
	Execute(*Client)
	SetNext(HandlerPattern)
}

type HandlerOne struct {
	next HandlerPattern
}

func (h *HandlerOne) Execute(c *Client) {
	fmt.Println("One ", c.Name)
	h.next.Execute(c)
}

func (h *HandlerOne) SetNext(next HandlerPattern) {
	h.next = next
}

type HandlerTwo struct {
	next HandlerPattern
}

func (h *HandlerTwo) Execute(c *Client) {
	fmt.Println("Two ", c.Name)
	h.next.Execute(c)
}

func (h *HandlerTwo) SetNext(next HandlerPattern) {
	h.next = next
}

type HandlerThree struct {
	next HandlerPattern
}

func (h *HandlerThree) Execute(c *Client) {
	fmt.Println("Three ", c.Name)
}

func (h *HandlerThree) SetNext(next HandlerPattern) {
	h.next = next
}
