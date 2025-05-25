package pkg

import (
	"image/color"
)

type Triangle struct {
	Vertices [3]Vec3
	Color    color.RGBA
}

type Mesh struct {
	Triangles []Triangle
}

func NewCubeMesh() *Mesh {
	return &Mesh{
		Triangles: []Triangle{
			{[3]Vec3{{-1, -1, 1}, {1, -1, 1}, {1, 1, 1}}, color.RGBA{255, 0, 0, 255}},
			{[3]Vec3{{-1, -1, 1}, {1, 1, 1}, {-1, 1, 1}}, color.RGBA{255, 0, 0, 255}},
			{[3]Vec3{{1, -1, -1}, {-1, -1, -1}, {-1, 1, -1}}, color.RGBA{0, 255, 0, 255}},
			{[3]Vec3{{1, -1, -1}, {-1, 1, -1}, {1, 1, -1}}, color.RGBA{0, 255, 0, 255}},
			{[3]Vec3{{-1, -1, -1}, {-1, -1, 1}, {-1, 1, 1}}, color.RGBA{0, 0, 255, 255}},
			{[3]Vec3{{-1, -1, -1}, {-1, 1, 1}, {-1, 1, -1}}, color.RGBA{0, 0, 255, 255}},
			{[3]Vec3{{1, -1, 1}, {1, -1, -1}, {1, 1, -1}}, color.RGBA{255, 255, 0, 255}},
			{[3]Vec3{{1, -1, 1}, {1, 1, -1}, {1, 1, 1}}, color.RGBA{255, 255, 0, 255}},
			{[3]Vec3{{-1, 1, 1}, {1, 1, 1}, {1, 1, -1}}, color.RGBA{255, 0, 255, 255}},
			{[3]Vec3{{-1, 1, 1}, {1, 1, -1}, {-1, 1, -1}}, color.RGBA{255, 0, 255, 255}},
			{[3]Vec3{{-1, -1, -1}, {1, -1, -1}, {1, -1, 1}}, color.RGBA{0, 255, 255, 255}},
			{[3]Vec3{{-1, -1, -1}, {1, -1, 1}, {-1, -1, 1}}, color.RGBA{0, 255, 255, 255}},
		},
	}
}
