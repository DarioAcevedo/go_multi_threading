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
func (boid *Boid) CalcAcceleration() Vector2d {
	upper, lower := boid.position.AddValue(viewRadius), boid.position.AddValue(-viewRadius)
	averagePosition, averageVelocity, separation := Vector2d{}, Vector2d{}, Vector2d{}
	count := 0.0
	lock.RLock()
	for x := math.Max(lower.x, 0); x <= math.Min(upper.x, screenWidth); x++ {
		for y := math.Max(lower.y, 0); y <= math.Min(upper.y, screenHeight); y++ {
			if i := boidMap[int(x)][int(y)]; i != -1 && i != boid.id {
				otherBoid := boids[i]
				if dist := otherBoid.position.Distance(boid.position); dist < viewRadius {
					count++
					averageVelocity = averageVelocity.Add(otherBoid.velocity)
					averagePosition = averagePosition.Add(otherBoid.position)
					positionDifference := boid.position.Substract(otherBoid.position).DivValue(dist)
					separation = separation.Add(positionDifference)

				}
			}
		}
	}

	lock.RUnlock()

	accel := Vector2d{
		x: boid.borderBounce(boid.position.x, screenWidth),
		y: boid.borderBounce(boid.position.y, screenHeight),
	}

	if count > 0 {
		averageVelocity = averageVelocity.DivValue(count)
		averagePosition = averagePosition.DivValue(count)

		alignAccel := averageVelocity.Substract(boid.velocity).MultValue(adjRate)
		cohesionAccel := averagePosition.Substract(boid.position).MultValue(adjRate)
		separationAccel := separation.MultValue(adjRate)

		return accel.Add(alignAccel).Add(cohesionAccel).Add(separationAccel)
	}

	return accel
}

func (b *Boid) borderBounce(p, mBp float64) float64 {
	if p < viewRadius {
		return 1 / p
	} else if p > mBp-viewRadius {
		return 1 / (p - mBp)
	}

	return 0
}

func (b *Boid) update() {
	acceleration := b.CalcAcceleration()

	lock.Lock()

	b.velocity = b.velocity.Add(acceleration).limit(-1, 1)
	boidMap[int(b.position.x)][int(b.position.y)] = -1
	b.position = b.position.Add(b.velocity)
	boidMap[int(b.position.x)][int(b.position.y)] = b.id

	lock.Unlock()
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
