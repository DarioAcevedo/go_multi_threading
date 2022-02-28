package main

import (
	"math/rand"
	"time"
)

type Boid struct {
	
	position Vector2d
	velocity Vector2d
	id int
}

func (b *Boid) update() {
	b.position = b.position.Add(b.velocity)
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
			x: rand.Float64() * 2 - 1.0,
			y: rand.Float64() * 2 - 1.0,
		},
		id: bid,
	}
	boids[bid] = &b
	boidMap[int(b.position.x)][int(b.position.y)] = bid
	go b.start()
}