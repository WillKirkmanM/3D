package pkg

import "math"

type Vec3 struct {
    X, Y, Z float64
}

func (v Vec3) Add(other Vec3) Vec3 {
    return Vec3{v.X + other.X, v.Y + other.Y, v.Z + other.Z}
}

func (v Vec3) Sub(other Vec3) Vec3 {
    return Vec3{v.X - other.X, v.Y - other.Y, v.Z - other.Z}
}

func (v Vec3) Scale(s float64) Vec3 {
    return Vec3{v.X * s, v.Y * s, v.Z * s}
}

func (v Vec3) Dot(other Vec3) float64 {
    return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v Vec3) Cross(other Vec3) Vec3 {
    return Vec3{
        v.Y*other.Z - v.Z*other.Y,
        v.Z*other.X - v.X*other.Z,
        v.X*other.Y - v.Y*other.X,
    }
}

func (v Vec3) Length() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vec3) Normalize() Vec3 {
    length := v.Length()
    if length == 0 {
        return Vec3{}
    }
    return Vec3{v.X / length, v.Y / length, v.Z / length}
}

type Mat4 [16]float64

func NewMat4Identity() Mat4 {
    return Mat4{
        1, 0, 0, 0,
        0, 1, 0, 0,
        0, 0, 1, 0,
        0, 0, 0, 1,
    }
}

func (m Mat4) Multiply(other Mat4) Mat4 {
    var result Mat4
    for i := 0; i < 4; i++ {
        for j := 0; j < 4; j++ {
            for k := 0; k < 4; k++ {
                result[i*4+j] += m[i*4+k] * other[k*4+j]
            }
        }
    }
    return result
}

func (m Mat4) MultiplyVec3(v Vec3) Vec3 {
    x := m[0]*v.X + m[1]*v.Y + m[2]*v.Z + m[3]
    y := m[4]*v.X + m[5]*v.Y + m[6]*v.Z + m[7]
    z := m[8]*v.X + m[9]*v.Y + m[10]*v.Z + m[11]
    w := m[12]*v.X + m[13]*v.Y + m[14]*v.Z + m[15]
    
    if w != 0 {
        return Vec3{x / w, y / w, z / w}
    }
    return Vec3{x, y, z}
}

func PerspectiveMatrix(fov, aspect, near, far float64) Mat4 {
    f := 1.0 / math.Tan(fov/2.0)
    return Mat4{
        f / aspect, 0, 0, 0,
        0, f, 0, 0,
        0, 0, (far + near) / (near - far), (2 * far * near) / (near - far),
        0, 0, -1, 0,
    }
}

func RotationY(angle float64) Mat4 {
    cos := math.Cos(angle)
    sin := math.Sin(angle)
    return Mat4{
        cos, 0, sin, 0,
        0, 1, 0, 0,
        -sin, 0, cos, 0,
        0, 0, 0, 1,
    }
}