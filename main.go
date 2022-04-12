package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil" //ebiten工具集
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
	gridSize     = 10
	xNumInScreen = screenWidth / gridSize
	yNumInScreen = screenHeight / gridSize
)

const (
	dirNone = iota
	dirLeft
	dirRight
	dirDown
	dirUp
)

type Position struct {
	X int
	Y int
}

type Game struct {
	moveDirection int
	snakeBody     []Position
	timer         int
	moveTime      int
}

func (g *Game) reset() {
	g.snakeBody[0].X = xNumInScreen / 2
	g.snakeBody[0].Y = yNumInScreen / 2
	g.snakeBody[1].X = g.snakeBody[0].X - 1
	g.snakeBody[1].Y = g.snakeBody[0].Y
	g.snakeBody[2].X = g.snakeBody[0].X + 1
	g.snakeBody[2].Y = g.snakeBody[0].Y
	g.snakeBody[3].X = g.snakeBody[0].X
	g.snakeBody[3].Y = g.snakeBody[0].Y - 1
}

func (g *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) || inpututil.IsKeyJustPressed(ebiten.KeyA) {
		if g.moveDirection != dirRight {
			g.moveDirection = dirLeft
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) || inpututil.IsKeyJustPressed(ebiten.KeyD) {
		if g.moveDirection != dirLeft {
			g.moveDirection = dirRight
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) || inpututil.IsKeyJustPressed(ebiten.KeyS) {
		if g.moveDirection != dirUp {
			g.moveDirection = dirDown
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) || inpututil.IsKeyJustPressed(ebiten.KeyW) {
		if g.moveDirection != dirDown {
			g.moveDirection = dirUp
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.reset()
	}
	if g.needsToMoveSnake() {
		switch g.moveDirection {
		case dirLeft:
			g.snakeBody[0].X--
			g.snakeBody[1].X = g.snakeBody[0].X - 1
			g.snakeBody[2].X = g.snakeBody[0].X + 1
			g.snakeBody[3].X = g.snakeBody[0].X
		case dirRight:
			g.snakeBody[0].X++
			g.snakeBody[1].X = g.snakeBody[0].X - 1
			g.snakeBody[2].X = g.snakeBody[0].X + 1
			g.snakeBody[3].X = g.snakeBody[0].X
		case dirDown:
			g.snakeBody[0].Y++
			g.snakeBody[1].Y = g.snakeBody[0].Y
			g.snakeBody[2].Y = g.snakeBody[0].Y
			g.snakeBody[3].Y = g.snakeBody[0].Y - 1
		case dirUp:
			g.snakeBody[0].Y--
			g.snakeBody[1].Y = g.snakeBody[0].Y
			g.snakeBody[2].Y = g.snakeBody[0].Y
			g.snakeBody[3].Y = g.snakeBody[0].Y - 1
		}
	}
	g.timer++
	fmt.Println(g.timer)
	return nil
}

func (g *Game) needsToMoveSnake() bool {
	return g.timer%g.moveTime == 0
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, v := range g.snakeBody {
		ebitenutil.DrawRect(screen, float64(v.X*gridSize), float64(v.Y*gridSize), gridSize, gridSize, color.RGBA{0x80, 0xa0, 0xc0, 0xff})
	}
	ebitenutil.DrawRect(screen, float64(xNumInScreen/4*gridSize), float64(yNumInScreen/4*gridSize), gridSize, gridSize, color.RGBA{0xFF, 0x00, 0x00, 0xff})
	if g.moveDirection == dirNone {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("Press up/down/left/right to start"))
	} else {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func newGame() *Game {
	g := &Game{
		snakeBody: make([]Position, 4),
		moveTime:  4,
	}
	g.snakeBody[0].X = xNumInScreen / 2
	g.snakeBody[0].Y = yNumInScreen / 2
	g.snakeBody[1].X = g.snakeBody[0].X - 1
	g.snakeBody[1].Y = g.snakeBody[0].Y
	g.snakeBody[2].X = g.snakeBody[0].X + 1
	g.snakeBody[2].Y = g.snakeBody[0].Y
	g.snakeBody[3].X = g.snakeBody[0].X
	g.snakeBody[3].Y = g.snakeBody[0].Y - 1
	return g
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("弹幕游戏")
	if err := ebiten.RunGame(newGame()); err != nil {
		log.Fatal(err)
	}
}
