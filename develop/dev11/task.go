package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"./store"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.


site
Реализовать HTTP-сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP-библиотекой.

В рамках задания необходимо:
Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
Реализовать middleware для логирования запросов


Методы API:
POST /create_event
POST /update_event
POST /delete_event
GET /events_for_day
GET /events_for_week
GET /events_for_month

Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09). В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON-документ содержащий либо {"result": "..."} в случае успешного выполнения метода, либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
Реализовать все методы.
Бизнес логика НЕ должна зависеть от кода HTTP сервера.
В случае ошибки бизнес-логики сервер должен возвращать HTTP 503.
В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400.
В случае остальных ошибок сервер должен возвращать HTTP 500.
Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.

*/
//
type storeServer struct {
	store *store.StoreServer
}

// func инициализация хранилища
func newStoreSerever() *storeServer {
	store := store.NewStore()
	return &storeServer{store: store}
}

// обработчики пост запросов
func (ss *storeServer) handlerCreateEvent(w http.ResponseWriter, r *http.Request) {
	_, date, mes, err := handlerDataPost(w, r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	id := ss.store.CreateEvent(date, mes)
	renderJSON(w, id)
}

func (ss *storeServer) handlerUpdateEvent(w http.ResponseWriter, r *http.Request) {
	id, date, mes, err := handlerDataPost(w, r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	event, err := ss.store.UpdateEvent(id, date, mes)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}
	renderJSON(w, event)
}

func (ss *storeServer) handlerDeleteEvent(w http.ResponseWriter, r *http.Request) {
	id, _, _, err := handlerDataPost(w, r)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	errDelete := ss.store.DeleteEvent(id)
	if errDelete != nil {
		http.Error(w, errDelete.Error(), 503)
		return
	}

	renderJSON(w, "great")

}

// обработчики гет запросов
func (ss *storeServer) handlerEventsForDay(w http.ResponseWriter, r *http.Request) {
	// обрабатываем данные с формы
	date := handlerDataGet(r)
	// обращаемся к хранилищу для обработки
	events, err := ss.store.EventsForDay(date, 0)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	// выводим на экран
	renderJSON(w, events)
}
func (ss *storeServer) handlerEventsForWeek(w http.ResponseWriter, r *http.Request) {
	date := handlerDataGet(r)

	events, err := ss.store.EventsForDay(date, 7)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	renderJSON(w, events)
}
func (ss *storeServer) handlerEventsForMonth(w http.ResponseWriter, r *http.Request) {
	date := handlerDataGet(r)

	events, err := ss.store.EventsForDay(date, 30)
	if err != nil {
		http.Error(w, err.Error(), 503)
		return
	}

	renderJSON(w, events)
}

// обработчик данных пост запроса
func handlerDataPost(w http.ResponseWriter, r *http.Request) (int, time.Time, string, error) {
	var id int
	var date time.Time
	var mes string

	// сначала проверяем вернула ли форма хоть что-то, а далее пробуем привести в нужный фoрмат
	idString := r.FormValue("id")
	if idString != "" {
		idInt, err := strconv.Atoi(idString)
		if err != nil {
			return 0, time.Time{}, "", errors.New("400: bad int")
		}

		id = idInt
	}

	dateString := r.FormValue("date")
	if dateString != "" {
		dateString += "T00:00:00Z"
		dateTime, err := time.Parse(time.RFC3339, dateString)
		if err != nil {
			return 0, time.Time{}, "", errors.New("400: bad date")
		}

		date = dateTime
	}

	mes = r.FormValue("mes")

	return id, date, mes, nil
}

// обработчик данных гетзапроса
func handlerDataGet(r *http.Request) time.Time {
	dateF := r.FormValue("date") + "T00:00:00Z"
	date, err := time.Parse(time.RFC3339, dateF)
	if err != nil {
		fmt.Println(err)
	}
	return date
}

// вывод джейсона на страницу
func renderJSON(w http.ResponseWriter, v interface{}) {
	fmt.Printf("v: %v\n", v)

	resultJSON := struct {
		Result interface{}
	}{Result: v}

	js, err := json.Marshal(&resultJSON)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

func main() {

	mux := http.NewServeMux()
	// создания хранилищас сервера
	ss := newStoreSerever()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	//разные точки входа
	mux.HandleFunc("/create_event", ss.handlerCreateEvent)
	mux.HandleFunc("/update_event", ss.handlerUpdateEvent)
	mux.HandleFunc("/delete_event", ss.handlerDeleteEvent)
	mux.HandleFunc("/events_for_day", ss.handlerEventsForDay)
	mux.HandleFunc("/events_for_week", ss.handlerEventsForWeek)
	mux.HandleFunc("/events_for_month", ss.handlerEventsForMonth)
	// запуск сервера
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
