package main

import (
	"embed"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/mpetavy/common"
	"golang.org/x/image/colornames"
	"time"
)

//go:embed go.mod
var resources embed.FS

func init() {
	common.Init("", "", "", "", "test", "", "", "", &resources, nil, nil, run, 0)
}

func run() error {
	common.DebugFunc()

	monitors := pixelgl.Monitors()

	x, y := monitors[0].Position()
	w, h := monitors[0].Size()

	dim := pixel.R(x, y, w, h)

	cfg := pixelgl.WindowConfig{
		Title: "Pixel Rocks!",
		//Bounds: pixel.R(0, 0, 1024, 768),
		Bounds: dim,
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if common.Error(err) {
		return err
	}

	win.SetMonitor(monitors[0])
	win.Clear(colornames.Skyblue)

	time.AfterFunc(time.Second*3, func() {
		win.SetMonitor(nil)
		win.SetBounds(dim)
	})

	p, err := LoadPicture("ghost.png")
	if common.Error(err) {
		return err
	}

	s := pixel.NewSprite(p, p.Bounds())
	s.Draw(win, pixel.IM)

	for !win.Closed() {
		win.Update()
	}

	return nil
}

func main() {
	pixelgl.Run(func() {
		common.Run(nil)
	})
}
