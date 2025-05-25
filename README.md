<p align="center">
  <img src="https://avatars.githubusercontent.com/u/138057124?s=200&v=4" width="150" />
</p>
<h1 align="center">3D</h1>

<p align="center">A 3D Game Engine leveraging a software-based 3D rendering for Perspective Projection, Z-Buffering & Mesh Rendering</p>

## Features

- **3D Math Library**: Vector and matrix operations for 3D transformations
- **Software Rasterisation**: Triangle rasterisation with barycentric coordinates
- **Z-Buffer**: Depth testing for proper 3D rendering
- **Perspective Projection**: 3D to 2D projection with field of view
- **Back-face Culling**: Optimisation to skip invisible triangles
- **Camera System**: First-person style camera with position and rotation
- **Mesh Rendering**: Support for triangular meshes with colored faces

## Controls

- **WASD**: Move camera forward/backward and strafe left/right
- **Arrow Keys**: Rotate camera left/right
- **Q/E**: Move camera up/down

## Project Structure

```
3D/
├── cmd/
│   ├── game.go          # Game Loop & Input Handling
│   └── main.go          # Entry Point
├── pkg/
│   ├── renderer.go      # 3D Renderer Implementation
│   ├── math.go          # 3D Math Utilities (Vec3, Mat4)
│   └── mesh.go          # Mesh & Triangle Definitions
├── go.mod
└── README.md
```

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/WillKirkmanM/3D
cd 3D
```

2. Install Dependencies:
```bash
go mod tidy
```

3. Run the engine:
```bash
go run cmd/main.go
```

## Usage Example

```go
// Create a new game instance
game := NewGame()

// Create a renderer
renderer := pkg.NewRenderer(800, 600)

// Create a cube mesh
cube := pkg.NewCubeMesh()

// Render the mesh with a rotation transform
transform := pkg.RotationY(time.Now().Sub(startTime).Seconds())
renderer.RenderMesh(cube, transform)
```

## Technical Details

### Rendering Pipeline

1. **Model Transform**: Apply object transformations (rotation, translation, scale)
2. **View Transform**: Apply camera position and rotation
3. **Projection**: Convert 3D coordinates to 2D screen coordinates
4. **Clipping**: Remove triangles outside the view frustum
5. **Back-face Culling**: Skip triangles facing away from camera
6. **Rasterisation**: Fill triangles pixel by pixel using barycentric coordinates
7. **Z-Testing**: Use depth buffer to handle overlapping surfaces

## Extending the Engine

### Adding New Mesh Types

```go
func NewPyramidMesh() *pkg.Mesh {
    return &pkg.Mesh{
        Triangles: []pkg.Triangle{
            // Define triangle vertices and colors
        },
    }
}
```

### Adding Lighting

The engine can be extended with basic lighting by calculating dot products between triangle normals and light direction vectors.

### Adding Textures

Texture mapping can be implemented by interpolating UV coordinates across triangles during rasterisation.

## Dependencies

- [Ebiten v2](https://github.com/hajimehoshi/ebiten) - 2D game library for Go
- Go standard library (math, image, etc.)