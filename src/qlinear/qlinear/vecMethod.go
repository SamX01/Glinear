package qlinear

import "math"

func (u Vec3f) crossProduct(v Vec3f) Vec3f {
	return Vec3f{F1: u.F2*v.F3 - u.F3*v.F2,
		F2: u.F3*v.F1 - u.F1*v.F3,
		F3: u.F1*v.F2 - u.F2*v.F1,
	}
}

func (u Vec3f) innerProduct(v Vec3f) float32 {
	return u.F1*v.F1 + u.F2*v.F2 + u.F3*v.F3
}

func (x Vec3f) norm() float32 {
	return sqrt(x.F1*x.F1 + x.F2*x.F2 + x.F3*x.F3)
}

func sqrt(x float32) float32 {
	var xhalf float32 = 0.5 * x // get bits for floating VALUE
	i := math.Float32bits(x)    // gives initial guess y0
	i = 0x5f375a86 - (i >> 1)   // convert bits BACK to float
	x = math.Float32frombits(i) // Newton step, repeating increases accuracy
	x = x * (1.5 - xhalf*x*x)
	x = x * (1.5 - xhalf*x*x)
	x = x * (1.5 - xhalf*x*x)
	return x
}
