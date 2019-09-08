package lib

import (
	// It FFI
	. "github.com/purescript-native/go-runtime"
	"github.com/veandco/go-sdl2/sdl"
)

func init() {
	exports := Foreign("Lib.Sdl2")

	exports["quit"] = func() Any {
		sdl.Quit()
		return nil
	}

	exports["initTimer"] = sdl.INIT_TIMER
	exports["initAudio"] = sdl.INIT_AUDIO
	exports["initVideo"] = sdl.INIT_VIDEO
	exports["initJoystick"] = sdl.INIT_JOYSTICK
	exports["initHaptic"] = sdl.INIT_HAPTIC
	exports["initGamecontroller"] = sdl.INIT_GAMECONTROLLER
	exports["initEvents"] = sdl.INIT_EVENTS
	exports["initNoparachute"] = sdl.INIT_NOPARACHUTE
	exports["initEverything"] = sdl.INIT_EVERYTHING

	exports["doit"] = func() Any {
		if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
			panic(err)
		}
		defer sdl.Quit()

		window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
			800, 600, sdl.WINDOW_SHOWN)
		if err != nil {
			panic(err)
		}
		defer window.Destroy()

		surface, err := window.GetSurface()
		if err != nil {
			panic(err)
		}
		surface.FillRect(nil, 0)

		rect := sdl.Rect{0, 0, 200, 200}
		surface.FillRect(&rect, 0xffff0000)
		window.UpdateSurface()

		running := true
		for running {
			for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
				switch event.(type) {
				case *sdl.QuitEvent:
					println("Quit")
					running = false
					break
				}
			}
		}
		return nil
	}
}
