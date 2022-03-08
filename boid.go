package main

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	position Vector2d
	velocity Vector2d
	id       int
}

// Take into account the view radius to calculate
//the average change on velocity our boid will have
func CalcAcceleration(boid *Boid) Vector2d {
	upper, lower := boid.position.AddValue(viewRadius), boid.position.AddValue(-viewRadius)
	averageVelocity := Vector2d{x: 0, y: 0}
	count := 0
	for x := math.Max(lower.x, 0); x <= math.Min(upper.x, screenWidth); x++ {
		for y := math.Max(lower.y, 0); y <= math.Min(upper.y, screenHeight); y++ {
			if i := boidMap[int(x)][int(y)]; i != -1 && i != boid.id {
				otherBoid := boids[i]
				if otherBoid.position.Distance(boid.position) <= viewRadius {
					count += 1
					averageVelocity = averageVelocity.Add(otherBoid.velocity)
				}
			}
		}
	}
	if count > 0 {
		averageVelocity = averageVelocity.DivValue(float64(count))
		return averageVelocity.Substract(boid.velocity).MultValue(adjRate)
	}
	return averageVelocity
}

func (b *Boid) update() {
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	b.velocity = b.velocity.Add(CalcAcceleration(b))
	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.x)][int(b.position.y)] = b.id
	next := b.position.Add(b.velocity)
	if next.x > screenWidth || next.x < 0 {
		b.velocity.x = -b.velocity.x
	}
	if next.y > screenHeight || next.y < 0 {
		b.velocity.y = -b.velocity.y
	}
}

func (b *Boid) start() {
	for {
		b.update()
		time.Sleep(5 * time.Millisecond)
	}
}

func CreateBoid(bid int) {
	b := Boid{
		position: Vector2d{
			x: rand.Float64() * screenWidth,
			y: rand.Float64() * screenHeight,
		},
		velocity: Vector2d{
			x: rand.Float64()*2 - 1.0,
			y: rand.Float64()*2 - 1.0,
		},
		id: bid,
	}
	boids[bid] = &b
	boidMap[int(b.position.x)][int(b.position.y)] = bid
	go b.start()
}
