package event

import (
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Manager struct {
	log      *logrus.Logger
	events   map[uuid.UUID]*Event
	category map[string][]uuid.UUID

	mu sync.RWMutex
}

func NewManager(log *logrus.Logger) *Manager {
	return &Manager{
		log:      log,
		events:   map[uuid.UUID]*Event{},
		category: map[string][]uuid.UUID{},
	}
}

func (m *Manager) NewEvent(category, name string, startAt time.Time) (e *Event) {
	e = NewEvent(m.log, name, category, startAt)
	m.log.Error("NewEvent", e.Category, e.Name, startAt)

	m.mu.Lock()
	m.events[e.ID] = e
	m.category[e.Category] = append(m.category[e.Category], e.ID)
	m.mu.Unlock()

	ch := time.After(startAt.Sub(time.Now()))
	go func() {
		<-ch
		m.log.Error("Start event", e.Category, e.Name)
		err := e.Start()
		if err != nil {
			m.log.Error("e.Start", err)
		}
	}()

	return
}

func (m *Manager) Categories() (categories []string) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	categories = make([]string, 0, len(m.category))
	for c, _ := range m.category {
		categories = append(categories, c)
	}
	return
}

func (m *Manager) EventsByCategory(category string) (events []*Event) {
	m.mu.RLock()
	for _, eventID := range m.category[category] {
		events = append(events, m.events[eventID])
	}
	m.mu.RUnlock()

	sort.Sort(byStartAt(events))
	return
}

type byStartAt []*Event

func (e byStartAt) Len() int           { return len(e) }
func (e byStartAt) Less(i, j int) bool { return e[i].StartAt.Before(e[j].StartAt) }
func (e byStartAt) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

func (m *Manager) SettleEvent(eventID uuid.UUID, winSide WinSide) (err error) {
	m.mu.RLock()
	e, ok := m.events[eventID]
	m.mu.RUnlock()

	if !ok {
		err = ErrNotFound
		return
	}

	err = e.Settle(winSide)
	return
}

func (m *Manager) DeleteEvent(eventID uuid.UUID) {
	m.mu.Lock()
	delete(m.events, eventID)
	m.mu.Unlock()
}
