package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern


	Пораждающий паттерн - гибкое создание объектов без внесения лишних зависимостей
Позволяет использовать один и тот же код, для получения разных представлений объектов (у объектов есть что-то общее)
Используется когда необходимо упростить создания сложного объекта (позволяет разбить его создания на части)

 Когда вы хотите избавиться от «телескопического конструктора».

 Допустим, у вас есть один конструктор с десятью опциональными параметрами.
  Его неудобно вызывать, поэтому вы создали ещё десять конструкторов с меньшим количеством параметров.
   Всё, что они делают — это переадресуют вызов к базовому конструктору, подавая какие-то значения
   по умолчанию в параметры, которые пропущены в них самих.

   Паттерн Строитель позволяет собирать объекты пошагово, вызывая только те шаги, которые вам нужны. А значит, больше не нужно пытаться «запихнуть» в конструктор все возможные опции продукта.

 Когда ваш код должен создавать разные представления какого-то объекта. Например,
 деревянные и железобетонные дома.
 Строитель можно применить, если создание нескольких представлений объекта состоит из одинаковых этапов,
  которые отличаются в деталях.
Интерфейс строителей определит все возможные этапы конструирования. Каждому представлению будет
соответствовать собственный класс-строитель. А порядок этапов строительства будет задавать класс-директор.

Когда вам нужно собирать сложные составные объекты, например, деревья Компоновщика.

 Строитель конструирует объекты пошагово, а не за один проход. Более того, шаги
  строительства можно выполнять рекурсивно. А без этого не построить древовидную структуру,
   вроде Компоновщика.
Заметьте, что Строитель не позволяет посторонним объектам иметь доступ к конструируемому объекту,
 пока тот не будет полностью готов. Это предохраняет клиентский код от получения незаконченных
 «битых» объектов.

Плюсы:
Позволяет пошагово создавать сложные объекты.
Позволяет использовать один и тот же код для реализации различных объектов
Минусы:
усложняет код программы введением новых структур
клиент привязан к классам строителей

*/

type builder interface {
	pouringFoundation()
	putWalls()
	buildingRoof()
	addDecor()
}

type firstBuilder struct {
	Foundation string
	Walls      string
	Roof       string
	Decor      string
}

func NewFirstBuilder() *firstBuilder {
	return &firstBuilder{}
}

func (fb *firstBuilder) pouringFoundation() {
	println("First: found")
}

func (fb *firstBuilder) putWalls() {
	println("First: walls", fb.Walls)
}

func (fb *firstBuilder) buildingRoof() {
	println("First: roof")
}

func (fb *firstBuilder) addDecor() {
	println("First: decor")
}

type SecondBuilder struct {
	Foundation string
	Walls      string
	Roof       string
	Decor      string
}

func NewSecondBuild() *SecondBuilder {
	return &SecondBuilder{}
}

func (sb *SecondBuilder) pouringFoundation() {
	println("Second: found")
}

func (sb *SecondBuilder) putWalls() {
	println("Second:   walls")
}

func (sb *SecondBuilder) buildingRoof() {
	println("Second: roof")
}

func (sb *SecondBuilder) addDecor() {
	println("Second: decor")
}

type director struct {
	builder
}

func NewDirector(b builder) *director {
	return &director{b}
}

func (d *director) SetBulder(b builder) {
	d.builder = b
}

func (d *director) BuildHouse() {
	d.pouringFoundation()
	d.putWalls()
	d.buildingRoof()
	d.addDecor()
}
