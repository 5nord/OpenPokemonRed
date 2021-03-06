package store

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

// SpriteData wSpriteStateData1, wSpriteStateData2
var SpriteData [16]*Sprite

// SpriteImage sprite tile image
type SpriteImage struct {
	Index  int             // C1x2: pointer to current displayed tile image
	Images []*ebiten.Image // VRAM: all tile image
}

func (s *SpriteImage) Length() int {
	for i := 0; i < len(s.Images); i++ {
		if s.Images[i] == nil {
			return i
		}
	}
	return 10
}

// Sprite data
type Sprite struct {
	MovmentStatus              byte // C1x1 if bit7
	ScreenXPixel, ScreenYPixel int  // Pixel C1x4, C1x5
	AnimationFrame             uint // C1x7, C1x8 update on UpdateSprites
	Direction                  uint // C1x9
	WalkCounter                int
	MapXCoord, MapYCoord       int  // Coord C2x4, C2x5
	Delay                      uint // C2x8
	VRAM                       SpriteImage
	Simulated                  []uint
	MovementBytes              [2]byte // movement byte 1,2
	DeltaX, DeltaY             int
	RightHand                  bool // used to walk animation
	TextID                     int
	Hidden                     bool
	DoubleSpd                  bool
}

// AnimationCounter getter for animation counter
func (s *Sprite) AnimationCounter() uint {
	return s.AnimationFrame >> 2
}
func (s *Sprite) intraAnimationCounter() uint {
	return s.AnimationFrame % 4
}

// IsInvalidSprite check Sprite is valid data
func IsInvalidSprite(offset uint) bool {
	s := SpriteData[offset]
	return s == nil
}

// NumSprites a number of sprites at current map
func NumSprites() int {
	i := 0
	for SpriteData[i] != nil {
		i++
	}
	return i
}
