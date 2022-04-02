package main

import "math"

type Vector2d struct {
	x, y float64
}

func (v1 Vector2d) Add(v2 Vector2d) Vector2d {
	return Vector2d{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
	}
}

func (v1 Vector2d) Substract(v2 Vector2d) Vector2d {
	return Vector2d{
		x: v1.x - v2.x,
		y: v1.y - v2.y,
	}
}

func (v1 Vector2d) Multiply(v2 Vector2d) Vector2d {
	return Vector2d{
		x: v1.x * v2.x,
		y: v1.y * v2.y,
	}
}

func (v1 Vector2d) AddValue(d float64) Vector2d {
	return Vector2d{
		x: v1.x + d,
		y: v1.y + d,
	}
}

func (v1 Vector2d) MultValue(d float64) Vector2d {
	return Vector2d{
		x: v1.x * d,
		y: v1.y * d,
	}
}

func (v1 Vector2d) DivValue(d float64) Vector2d {
	return Vector2d{
		x: v1.x / d,
		y: v1.y / d,
	}
}

func (v1 Vector2d) Distance(v2 Vector2d) float64 {
	x1 := v1.x - v2.x
	x2 := v1.y - v2.y
	x := math.Sqrt(math.Pow(x1, 2) + math.Pow(x2, 2))
	return x
}

func (v1 Vector2d) limit(lower, upper float64) Vector2d {
	return Vector2d{math.Min(math.Max(v1.x, lower), upper),
		math.Min(math.Max(v1.y, lower), upper)}
}
