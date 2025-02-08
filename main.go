package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image/color"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	rl.InitWindow(800, 450, "hello game")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	player := Player{
		posX:           400,
		posY:           200,
		width:          30,
		height:         30,
		health:         3,
		color:          rl.Black,
		lastDamageTime: time.Now(),
		damageDuration: 500 * time.Millisecond,
	}

	var bullets []Bullet
	lastBulletSpawn := time.Now()

	for !rl.WindowShouldClose() {
		// Player movement
		if rl.IsKeyDown(rl.KeyW) {
			player.posY -= 5
		}
		if rl.IsKeyDown(rl.KeyA) {
			player.posX -= 5
		}
		if rl.IsKeyDown(rl.KeyS) {
			player.posY += 5
		}
		if rl.IsKeyDown(rl.KeyD) {
			player.posX += 5
		}

		if time.Since(lastBulletSpawn) > 10*time.Millisecond {
			newBullet := Bullet{
				posX:      int32(0),                 // Random X position
				posY:      int32(rand.Intn(450)),    // Random Y position
				diffX:     int32(rand.Intn(16) - 8), // Speed in X (-2 to 2)
				diffY:     int32(rand.Intn(16) - 8), // Speed in Y (-2 to 2)
				color:     rl.Red,
				createdAt: time.Now(),
			}
			bullets = append(bullets, newBullet)
			lastBulletSpawn = time.Now()
		}

		if time.Since(lastBulletSpawn) > 10*time.Millisecond {
			newBullet := Bullet{
				posX:      int32(rand.Intn(450)),    // Random X position
				posY:      int32(800),               // Random Y position
				diffX:     int32(rand.Intn(16) - 8), // Speed in X (-2 to 2)
				diffY:     int32(rand.Intn(16) - 8), // Speed in Y (-2 to 2)
				color:     rl.Red,
				createdAt: time.Now(),
			}
			bullets = append(bullets, newBullet)
			lastBulletSpawn = time.Now()
		}

		// Check bullet collision
		player.isBulletCollide(&bullets)

		// Reset player color after damage duration
		if player.damaged && time.Since(player.lastDamageTime) > player.damageDuration {
			player.damaged = false
		}

		// Remove bullets after 30 seconds
		bullets = removeOldBullets(bullets)

		// End game if health reaches 0
		if player.health <= 0 {
			return
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		// Player flashing effect
		playerColor := rl.Black
		if player.damaged {
			playerColor = rl.Red
		}
		rl.DrawRectangle(player.posX, player.posY, player.width, player.height, playerColor)

		// Draw and update bullets
		for i := range bullets {
			rl.DrawRectangle(bullets[i].posX, bullets[i].posY, 10, 10, bullets[i].color)
			bullets[i].posX += bullets[i].diffX
			bullets[i].posY += bullets[i].diffY
		}

		// Draw health
		rl.DrawText("Health: "+strconv.Itoa(player.health), 10, 10, 20, rl.Red)
		rl.EndDrawing()
	}
}

// Player struct
type Player struct {
	posX           int32
	posY           int32
	width          int32
	height         int32
	health         int
	color          color.RGBA
	lastDamageTime time.Time
	damageDuration time.Duration
	damaged        bool
}

// Bullet struct
type Bullet struct {
	posX      int32
	posY      int32
	diffX     int32
	diffY     int32
	color     color.RGBA
	createdAt time.Time
}

// Check if the player is hit by any bullet
func (p *Player) isBulletCollide(bullets *[]Bullet) {
	now := time.Now()
	if now.Sub(p.lastDamageTime) < time.Second {
		return
	}

	for i := 0; i < len(*bullets); i++ {
		bullet := (*bullets)[i]
		playerRect := rl.Rectangle{X: float32(p.posX), Y: float32(p.posY), Width: float32(p.width), Height: float32(p.height)}
		bulletRect := rl.Rectangle{X: float32(bullet.posX), Y: float32(bullet.posY), Width: 10, Height: 10}

		if rl.CheckCollisionRecs(playerRect, bulletRect) {
			p.health--
			p.lastDamageTime = now
			p.damaged = true
			break
		}
	}
}

// Removes bullets that are older than 30 seconds
func removeOldBullets(bullets []Bullet) []Bullet {
	newBullets := bullets[:0] // Create a new slice
	now := time.Now()
	for _, bullet := range bullets {
		if now.Sub(bullet.createdAt) < 30*time.Second {
			newBullets = append(newBullets, bullet)
		}
	}
	return newBullets
}
