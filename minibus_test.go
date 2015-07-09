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
	bus.Sub("/test/call", func(msg interface{}){
		called = true
	})
	bus.Pub("/test/call", "")
	time.Sleep(1 * time.Second)

	assert.True(t, called)
}