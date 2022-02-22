package store

import (
	"errors"
	"reflect"
	"sync"
	"time"
)

type EvetCalendar struct {
	Date time.Time
	Mes  string
}

// структура хранилища
type StoreServer struct {
	m     sync.Mutex
	store map[int]EvetCalendar
}

func NewStore() *StoreServer {
	ss := &StoreServer{}
	ss.m = sync.Mutex{}
	ss.store = make(map[int]EvetCalendar)
	return ss
}

func (ss *StoreServer) CreateEvent(date time.Time, mes string) int {

	event := EvetCalendar{date, mes}

	ss.m.Lock()
	defer ss.m.Unlock()

	id := len(ss.store)

	for {
		// добавляем новый элемент если карта была пустой
		if reflect.DeepEqual(ss.store[id], EvetCalendar{}) {
			ss.store[id] = event
			return id
		}
		id++
	}
}
func (ss *StoreServer) UpdateEvent(id int, date time.Time, mes string) (EvetCalendar, error) {
	ss.m.Lock()
	defer ss.m.Unlock()
	// вернем ошибку если элемента нет

	if reflect.DeepEqual(ss.store[id], EvetCalendar{}) {
		return EvetCalendar{}, errors.New("503: miss element")
	}

	event := EvetCalendar{date, mes}

	ss.store[id] = event

	return ss.store[id], nil

}
func (ss *StoreServer) DeleteEvent(id int) error {
	ss.m.Lock()
	defer ss.m.Unlock()
	// вернем ошибку если элемента нет
	if reflect.DeepEqual(ss.store[id], EvetCalendar{}) {
		return errors.New("503: miss element")
	}

	delete(ss.store, id)
	return nil
}

func (ss *StoreServer) EventsForDay(date time.Time, days int) ([]EvetCalendar, error) {
	var result []EvetCalendar
	for _, event := range ss.store {
		if event.Date.Sub(date) >= time.Duration(days*time.Now().Day()) {
			result = append(result, event)
		}
	}
	// вернем ошибку если элементы не были надены
	if len(result) == 0 {
		return []EvetCalendar{}, errors.New("503: no events")
	}

	return result, nil
}
