package models

import (
	"sync"
)

type MailEvent struct {
	*MailContent
	IntentType string `json:"intent_type"`
}

type EventEmitter struct {
	listeners sync.Map
}

// core function
func (emitter *EventEmitter) On(event string, listener func(MailEvent, error)) {
	// Load the current list of listeners for the given event
	listeners, _ := emitter.listeners.LoadOrStore(event, []func(MailEvent, error){})

	// Append the new listener to the list of listeners
	emitter.listeners.Store(event, append(listeners.([]func(MailEvent, error)), listener))
}

func (emitter *EventEmitter) Emit(event string, data MailEvent, err error) {
	if listeners, ok := emitter.listeners.Load(event); ok {

		for _, listener := range listeners.([]func(MailEvent, error)) {

			// use goroutine to avoid blocking
			go func(l func(MailEvent, error), d MailEvent, e error) {
				l(d, e)
			}(listener, data, err)

		}
	}
}
