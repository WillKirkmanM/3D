package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/WillKirkmanM/3D/pkg"
)

type Game struct {
	renderer *pkg.Renderer
	cube     *pkg.Mesh
	rotation float64
}

func NewGame() *Game {
	return &Game{
		renderer: pkg.NewRenderer(screenWidth, screenHeight),
		cube:     pkg.NewCubeMesh(),
	}
}

func (g *Game) Update() error {
	g.rotation += 0.02

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.renderer.Camera.Position.X -= 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.renderer.Camera.Position.X += 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.renderer.Camera.Position.Z += 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.renderer.Camera.Position.Z -= 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.renderer.Camera.Rotation -= 0.05
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.renderer.Camera.Rotation += 0.05
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.renderer.Clear()

	transform := pkg.RotationY(g.rotation)

	g.renderer.RenderMesh(g.cube, transform)

	screen.DrawImage(g.renderer.GetImage(), nil)

	ebitenutil.DebugPrint(screen, "FPS: "+fmt.Sprintf("%.2f", ebiten.ActualFPS()))
	ebitenutil.DebugPrint(screen, "\nControls: WASD - move, Arrow keys - rotate camera")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
