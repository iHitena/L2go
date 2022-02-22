package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern

struct pattern

	Фасад применяется, когда большой и сложной системе необходимо простой интерфейс, возможно с урезанным функционалом. (может упростить использования сложной системы)
	когда объекты сложно системы необходимо разделить на более простые
	Плюс: изолированность слиента от компонентов сложной системы
Минус: фасад может быть связан с большим количеством классов
Когда вам нужно представить простой или урезанный интерфейс к сложной подсистеме.

 Часто подсистемы усложняются по мере развития программы. Применение большинства паттернов
  приводит к появлению меньших классов, но в бóльшем количестве. Такую подсистему проще
   повторно использовать, настраивая её каждый раз под конкретные нужды, но вместе с тем, применять
   подсистему без настройки становится труднее. Фасад предлагает определённый вид системы по умолчанию,
   устраивающий большинство клиентов.
 Когда вы хотите разложить подсистему на отдельные слои.
 Используйте фасады для определения точек входа на каждый уровень подсистемы. Если подсистемы зависят
  друг от друга, то зависимость можно упростить, разрешив подсистемам обмениваться информацией
  только через фасады.
Например, возьмём ту же сложную систему видеоконвертации. Вы хотите разбить её на слои работы
с аудио и видео. Для каждой из этих частей можно попытаться создать фасад и заставить классы
аудио и видео обработки общаться друг с другом через эти фасады, а не напрямую.

*/
type person struct {
}

func (p person) authorizationPerson(personID string) {
	fmt.Println("Authorization person, personID: ", personID)

}

func (p person) deliveryStatusChanges() {
	fmt.Println("Delivery ok")
}

type product struct {
}

func (p product) productBooking(productID string) {
	fmt.Println("Product Booking. Product ID: ", productID)
}

type order struct {
}

func (o order) paymentConfirmation() {
	fmt.Println("Payment conficmation")
}

func (o order) delivery() {
	fmt.Println("Delivery")
}

type Facade struct {
	*person
}

func (f Facade) MakeOrder(personID string, productID string) {

	f.person = &person{}
	f.authorizationPerson(personID)

	product := product{}
	product.productBooking(productID)

	order := order{}
	order.paymentConfirmation()
	order.delivery()

	f.deliveryStatusChanges()
}
