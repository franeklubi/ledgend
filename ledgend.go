package ledgend

import (
    "time"
)


type Color struct {
    R, G, B uint8
}

type Animation struct {
    Direction                   bool
    Length                      float32     // float between 0 and 1
    Start_colour, End_colour    Color
    Duration                    time.Duration
    Start                       time.Time
}

type Buffer struct {
    length          uint16
    pixels          []Color
    animation_queue []Animation
}


func GenBuffer(length uint16) (Buffer) {
    b := Buffer{}
    b.length = length

    b.pixels = make([]Color, length)

    return b
}

func (b *Buffer) GetPixels() ([]Color) {
    return b.pixels
}

func (b *Buffer) GetAnimationQueue() ([]Animation) {
    return b.animation_queue
}

func XOR(b1 *Buffer, b2 *Buffer) {
}

func (b *Buffer) ApplyQueue() {
}

func (b *Buffer) AddAnimation(a Animation, as ...Animation) {
    b.animation_queue = append(b.animation_queue, a)
    for _, v := range as {
        b.animation_queue = append(b.animation_queue, v)
    }
}
