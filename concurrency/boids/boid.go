package boid

import (
	"image/color"
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	position  Vector2d
	velocity  Vector2d
	id        int
	boidColor color.RGBA
}

func (boid *Boid) start() {
	for {
		boid.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func (boid *Boid) moveOne() {
	boid.velocity = boid.velocity.Add(boid.calculateAcceleration()).Limit(-1, 1)
	boidArray[int(boid.position.x)][int(boid.position.y)] = -1
	boid.position = boid.position.Add(boid.velocity)
	boidArray[int(boid.position.x)][int(boid.position.y)] = boid.id
	next := boid.position.Add(boid.velocity)
	if next.x >= screenWidth || next.x < 0 {
		boid.velocity = Vector2d{-boid.velocity.x, boid.velocity.y}
	}
	if next.y >= screenHeight || next.y < 0 {
		boid.velocity = Vector2d{boid.velocity.x, -boid.velocity.y}
	}
}

func (boid *Boid) calculateAcceleration() Vector2d {
	upper, lower := boid.position.AddV(viewRadius), boid.position.AddV(-viewRadius)
	averageVelocity := Vector2d{0, 0}
	count := 0.0
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			if otherBoidId := boidArray[int(i)][int(j)]; otherBoidId != -1 && otherBoidId != boid.id {
				if dist := boids[otherBoidId].position.Distance(boid.position); dist < viewRadius {
					count++
					averageVelocity = averageVelocity.Add(boids[otherBoidId].velocity)
				}
			}
		}
	}

	acceleration := Vector2d{0, 0}
	if count > 0 {
		averageVelocity = averageVelocity.DivisionV(count)
		acceleration = averageVelocity.Subtract(boid.velocity).MultiplyV(adjustingRate)
	}
	return acceleration
}

func CreateBoid(id int) {
	min := 0
	max := 255
	boidColor := color.RGBA{
		uint8(rand.Intn(max-min+1) + min),
		uint8(rand.Intn(max-min+1) + min),
		uint8(rand.Intn(max-min+1) + min),
		uint8(rand.Intn(max-min+1) + min),
	}

	boid := Boid{
		position:  Vector2d{rand.Float64() * screenWidth, rand.Float64() * screenHeight},
		velocity:  Vector2d{(rand.Float64() * 2) - 1.0, (rand.Float64() * 2) - 1.0},
		id:        id,
		boidColor: boidColor,
	}
	boids[id] = &boid
	boidArray[int(boid.position.x)][int(boid.position.y)] = boid.id
	go boid.start()
}
