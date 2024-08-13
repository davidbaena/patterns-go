package eventbus

import (
	"testing"
	"time"
)

func TestPublishSendsEventToSubscribers(t *testing.T) {
	eb := NewEventBus()
	ch := make(chan Event, 1)
	eb.Subscribe("test_event", ch)

	event := Event{Type: "test_event", Timestamp: time.Now(), Data: "test_data"}
	eb.Publish(event)

	receivedEvent := <-ch
	if receivedEvent != event {
		t.Errorf("Expected event %v, got %v", event, receivedEvent)
	}
}

func TestPublishDoesNotSendEventToUnsubscribedTypes(t *testing.T) {
	eb := NewEventBus()
	ch := make(chan Event, 1)
	eb.Subscribe("other_event", ch)

	event := Event{Type: "test_event", Timestamp: time.Now(), Data: "test_data"}
	eb.Publish(event)

	select {
	case receivedEvent := <-ch:
		t.Errorf("Did not expect to receive event %v", receivedEvent)
	default:
	}
}

func TestPublishToMultipleSubscribers(t *testing.T) {
	eb := NewEventBus()
	ch1 := make(chan Event, 1)
	ch2 := make(chan Event, 1)
	eb.Subscribe("test_event", ch1)
	eb.Subscribe("test_event", ch2)

	event := Event{Type: "test_event", Timestamp: time.Now(), Data: "test_data"}
	eb.Publish(event)

	receivedEvent1 := <-ch1
	receivedEvent2 := <-ch2
	if receivedEvent1 != event || receivedEvent2 != event {
		t.Errorf("Expected event %v, got %v and %v", event, receivedEvent1, receivedEvent2)
	}
}

func TestPublishWithNoSubscribers(t *testing.T) {
	eb := NewEventBus()

	event := Event{Type: "test_event", Timestamp: time.Now(), Data: "test_data"}
	eb.Publish(event)

	// No assertion needed, just ensure no panic or error
}
