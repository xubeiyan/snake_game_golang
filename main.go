package main

import (
	"bytes"
	"errors"
	"image/color"
	"log"
	"snake/common"
	"snake/entity"
	"snake/game"
	"snake/math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var mplusFaceSource *text.GoTextFaceSource

type Game struct {
	world      *game.World
	lastUpdate time.Time
	gameOver   bool
}

func (g *Game) Update() error {
	if g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			g.world.ClearAllEntity()
			g.world.AddEntity(
				entity.NewPlayer(
					math.Point{
						X: common.ScreenWidth / common.GridSize / 2,
						Y: common.ScreenHeight / common.GridSize / 2,
					},
					math.DirRight,
				),
			)

			for range 2 {
				g.world.AddEntity(
					entity.NewFood(),
				)
			}

			g.gameOver = false
		}
		return nil
	}

	playerRaw, ok := g.world.GetFirstEntity("player")
	if !ok {
		return errors.New("entity player was nit found")
	}
	player := playerRaw.(*entity.Player)

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		player.SetDirection(math.DirUp)
	} else if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		player.SetDirection(math.DirDown)
	} else if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		player.SetDirection(math.DirLeft)
	} else if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		player.SetDirection(math.DirRight)
	}

	if time.Since(g.lastUpdate) < common.GameSpeed {
		return nil
	}
	g.lastUpdate = time.Now()

	for _, entity := range g.world.Entities() {
		if entity.Update(g.world) {
			g.gameOver = true
			return nil
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, entity := range g.world.Entities() {
		entity.Draw(screen)
	}

	if g.gameOver {
		face := &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   48,
		}

		hintFace := &text.GoTextFace{
			Source: mplusFaceSource,
			Size:   24,
		}

		t := "Game Over"
		w, h := text.Measure(
			t,
			face,
			face.Size,
		)

		hintText := "Press Space to restart"

		hw, hh := text.Measure(
			hintText,
			hintFace,
			hintFace.Size,
		)

		op := &text.DrawOptions{}
		op.GeoM.Translate(
			common.ScreenWidth/2-w/2, common.ScreenHeight/2-h/2-20,
		)
		op.ColorScale.ScaleWithColor(color.White)
		text.Draw(
			screen,
			t,
			face,
			op,
		)

		op = &text.DrawOptions{}
		op.GeoM.Translate(
			common.ScreenWidth/2-hw/2, common.ScreenHeight/2-hh/2+30,
		)
		op.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 255})
		text.Draw(
			screen,
			hintText,
			hintFace,
			op,
		)
	}
}

func (g *Game) Layout(
	outsideWidth,
	outsideHeight int,
) (int, int) {
	return common.ScreenWidth, common.ScreenHeight
}

func main() {
	s, err := text.NewGoTextFaceSource(
		bytes.NewReader(
			fonts.MPlus1pRegular_ttf,
		),
	)

	if err != nil {
		log.Fatal(err)
	}
	mplusFaceSource = s

	world := game.NewWorld()
	world.AddEntity(
		entity.NewPlayer(
			math.Point{
				X: common.ScreenWidth / common.GridSize / 2,
				Y: common.ScreenHeight / common.GridSize / 2,
			},
			math.DirRight,
		),
	)

	for range 2 {
		world.AddEntity(
			entity.NewFood(),
		)
	}
	g := &Game{
		world: world,
	}

	ebiten.SetWindowSize(common.ScreenWidth, common.ScreenHeight)
	ebiten.SetWindowTitle("Snake")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
