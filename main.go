// Copyright 2017 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build example
// +build example

package main

import (
	"bytes"
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	rplatformer "github.com/hajimehoshi/ebiten/v2/examples/resources/images/platformer"
)

const (
	// Settings
	screenWidth  = 960
	screenHeight = 540
)

var (
	leftSprite      *ebiten.Image
	rightSprite     *ebiten.Image
	idleSprite      *ebiten.Image
	backgroundImage *ebiten.Image
)

func init() {
	data, err := os.ReadFile("img/blue_plane.png")
	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		panic(err)
	}
	idleSprite = ebiten.NewImageFromImage(img)

	img, _, err = image.Decode(bytes.NewReader(rplatformer.Background_png))
	if err != nil {
		panic(err)
	}
	backgroundImage = ebiten.NewImageFromImage(img)
}

const (
	unit    = 16
	groundY = 380
)

type char struct {
	x  int
	y  int
	vx int
	vy int
}

func (c *char) update() {
	c.x += c.vx
	c.y += c.vy
	if c.y > groundY*unit {
		c.y = groundY * unit
	}
	if c.vx > 0 {
		c.vx -= 4
	} else if c.vx < 0 {
		c.vx += 4
	}
	if c.vy > 0 {
		c.vy -= 4
	} else if c.vy < 0 {
		c.vy += 4
	}
}

func (c *char) draw(screen *ebiten.Image) {
	s := idleSprite
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(float64(c.x)/unit, float64(c.y)/unit)
	screen.DrawImage(s, op)
}

type Game struct {
	gopher *char
}

func (g *Game) Update() error {
	if g.gopher == nil {
		g.gopher = &char{x: 50 * unit, y: groundY * unit}
	}

	// Controls
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.gopher.vx = -4 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.gopher.vx = 4 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.gopher.vy = -4 * unit
	} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.gopher.vy = 4 * unit
	}
	g.gopher.update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draws Background Image
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.5, 0.5)
	screen.DrawImage(backgroundImage, op)

	// Draws the Gopher
	g.gopher.draw(screen)

	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Platformer (Ebiten Demo)")
	if err := ebiten.RunGame(&Game{}); err != nil {
		panic(err)
	}
}
