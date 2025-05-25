package pkg

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer struct {
	width, height int
	zBuffer       []float64
	frameBuffer   *image.RGBA
	Camera        Camera
	projection    Mat4
}

type Camera struct {
	Position Vec3
	Rotation float64
}

func NewRenderer(width, height int) *Renderer {
	return &Renderer{
		width:       width,
		height:      height,
		zBuffer:     make([]float64, width*height),
		frameBuffer: image.NewRGBA(image.Rect(0, 0, width, height)),
		Camera:      Camera{Position: Vec3{0, 0, -5}},
		projection:  PerspectiveMatrix(math.Pi/4, float64(width)/float64(height), 0.1, 100.0),
	}
}

func (r *Renderer) Clear() {
	for i := range r.frameBuffer.Pix {
		r.frameBuffer.Pix[i] = 0
	}

	for i := range r.zBuffer {
		r.zBuffer[i] = math.Inf(1)
	}
}

func (r *Renderer) RenderMesh(mesh *Mesh, transform Mat4) {
	view := RotationY(-r.Camera.Rotation)
	view = view.Multiply(Mat4{
		1, 0, 0, -r.Camera.Position.X,
		0, 1, 0, -r.Camera.Position.Y,
		0, 0, 1, -r.Camera.Position.Z,
		0, 0, 0, 1,
	})

	mvp := r.projection.Multiply(view.Multiply(transform))

	for _, triangle := range mesh.Triangles {
		v0 := mvp.MultiplyVec3(triangle.Vertices[0])
		v1 := mvp.MultiplyVec3(triangle.Vertices[1])
		v2 := mvp.MultiplyVec3(triangle.Vertices[2])

		edge1 := v1.Sub(v0)
		edge2 := v2.Sub(v0)
		normal := edge1.Cross(edge2)
		if normal.Z > 0 {
			continue
		}

		p0 := r.toScreenCoords(v0)
		p1 := r.toScreenCoords(v1)
		p2 := r.toScreenCoords(v2)

		r.drawTriangle(p0, p1, p2, triangle.Color)
	}
}

func (r *Renderer) toScreenCoords(v Vec3) Vec3 {
	x := (v.X + 1) * float64(r.width) / 2
	y := (1 - v.Y) * float64(r.height) / 2
	return Vec3{x, y, v.Z}
}

func (r *Renderer) drawTriangle(v0, v1, v2 Vec3, col color.RGBA) {
	minX := int(math.Max(0, math.Min(math.Min(v0.X, v1.X), v2.X)))
	maxX := int(math.Min(float64(r.width-1), math.Max(math.Max(v0.X, v1.X), v2.X)))
	minY := int(math.Max(0, math.Min(math.Min(v0.Y, v1.Y), v2.Y)))
	maxY := int(math.Min(float64(r.height-1), math.Max(math.Max(v0.Y, v1.Y), v2.Y)))

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			p := Vec3{float64(x), float64(y), 0}
			bary := r.barycentric(v0, v1, v2, p)

			if bary.X >= 0 && bary.Y >= 0 && bary.Z >= 0 {
				z := v0.Z*bary.X + v1.Z*bary.Y + v2.Z*bary.Z
				idx := y*r.width + x

				if z < r.zBuffer[idx] {
					r.zBuffer[idx] = z
					r.frameBuffer.SetRGBA(x, y, col)
				}
			}
		}
	}
}

func (r *Renderer) barycentric(v0, v1, v2, p Vec3) Vec3 {
	v0v1 := v1.Sub(v0)
	v0v2 := v2.Sub(v0)
	v0p := p.Sub(v0)

	dot00 := v0v2.Dot(v0v2)
	dot01 := v0v2.Dot(v0v1)
	dot02 := v0v2.Dot(v0p)
	dot11 := v0v1.Dot(v0v1)
	dot12 := v0v1.Dot(v0p)

	invDenom := 1 / (dot00*dot11 - dot01*dot01)
	u := (dot11*dot02 - dot01*dot12) * invDenom
	v := (dot00*dot12 - dot01*dot02) * invDenom

	return Vec3{1 - u - v, v, u}
}

func (r *Renderer) GetImage() *ebiten.Image {
	return ebiten.NewImageFromImage(r.frameBuffer)
}
