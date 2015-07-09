package minibus
import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMiniBus(t *testing.T) {
	bus := NewMiniBus()
	assert.NotNil(t, bus)

	bus.Sub("/topic", func(interface{}){})
	assert.Equal(t, 1, len(bus.GetSub("/topic")))

	bus.Sub("/topic", func(interface{}){})
	assert.Equal(t, 2, len(bus.GetSub("/topic")))

	bus.Sub("/topic/second", func(interface{}){})
	assert.Equal(t, 1, len(bus.GetSub("/topic/second")))

	called := false
	bus.Sub("/topic/called", func(msg interface{}){
		called = true
	})
	bus.Pub("/topic/called", "")
	time.Sleep(1 * time.Second)

	assert.Equal(t, 1, len(bus.GetSub("/topic/called")))
	assert.True(t, called)

	bus.ClearTopic("/topic/called")
	assert.Equal(t, 0, len(bus.GetSub("/topic/called")))

	assert.Equal(t, 2, len(bus.GetSub("/topic")))

	bus.ClearSubs()
	assert.Equal(t, 0, len(bus.GetSub("/topic")))
}