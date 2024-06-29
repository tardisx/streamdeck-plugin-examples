package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"sync"
	"time"

	"github.com/tardisx/streamdeck-plugin"
	"github.com/tardisx/streamdeck-plugin/events"
	"github.com/tardisx/streamdeck-plugin/tools"
)

// keep track of instances we've seen, each one has a different
// colour
type clocks struct {
	contextColour map[string]string
	lock          sync.Mutex
}

// template for making a clock image in SVG
const svgClock = `
<svg width="144" height="144" xmlns="http://www.w3.org/2000/svg">
  <rect width="144" height="144" fill="black"/>
  <text x="72" y="108"
        font-family="Arial, sans-serif"
        font-size="96"
        font-weight="bold"
        fill="%s"
        text-anchor="middle"
        dominant-baseline="central"
		transform="rotate(%d, 72, 72)">
    %02d
  </text>
</svg>`

func main() {
	clocks := clocks{
		contextColour: map[string]string{},
		lock:          sync.Mutex{},
	}
	slog.Info("Starting up")
	c := streamdeck.NewWithLogger(slog.Default())

	slog.Info("Registering handlers")
	c.RegisterHandler(func(e events.ERWillAppear) {
		// clock appeared, give it a random colour
		slog.Info("appearing " + e.Context)
		clocks.lock.Lock()
		defer clocks.lock.Unlock()
		clocks.contextColour[e.Context] = randRGB()
	})
	c.RegisterHandler(func(e events.ERWillDisappear) {
		// Stop updating this clock by simply removing it from our struct.
		// Note that this is not required, and in this case it means that
		// when it gets re-instantiated it will get a new colour.
		// But it is good practice to not spend CPU on updating things that
		// are not currently being displayed.
		slog.Info("disappearing " + e.Context)
		clocks.lock.Lock()
		defer clocks.lock.Unlock()
		delete(clocks.contextColour, e.Context)
	})
	c.RegisterHandler(func(e events.ERKeyDown) {
		// button pressed, change its colour
		slog.Info("keyDown " + e.Context)
		clocks.lock.Lock()
		defer clocks.lock.Unlock()
		clocks.contextColour[e.Context] = randRGB()
		drawClock(c, e.Context, clocks.contextColour[e.Context])
	})

	slog.Info("Connecting web socket")
	err := c.Connect()
	if err != nil {
		panic(err)
	}

	// update all clocks, continuously
	go func() {
		for {
			clocks.lock.Lock()
			for context, colour := range clocks.contextColour {
				drawClock(c, context, colour)
			}
			clocks.lock.Unlock()
			time.Sleep(time.Second)
		}
	}()

	slog.Info("waiting for the end")
	c.WaitForPluginExit()
}

func drawClock(c streamdeck.Connection, context string, colour string) {
	// rotation for this minute of the hour
	rot := int(360.0 * (float64(time.Now().Minute()) / 60.0))
	// generate the SVG
	svg := fmt.Sprintf(svgClock, colour, rot, time.Now().Hour())

	// create the event
	newImage := events.NewESSetImage(
		context,
		tools.SVGToPayload(svg),
		events.EventTargetBoth,
		nil)

	// send it
	c.Send(newImage)
}

// randRGB creates a random colour
func randRGB() string {
	return fmt.Sprintf("#%02x%02x%02x", rand.Intn(256), rand.Intn(256), rand.Intn(256))
}
