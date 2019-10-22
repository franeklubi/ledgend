package ledgend

import (
    "time"
)


type Color struct {
    R, G, B uint8
}

type animation struct {
    name                        string
    start_colour, end_colour    Color
    length                      float32     // float between 0 and 1
    start                       time.Time
    duration                    time.Duration
}

type Buffer struct {
    length          uint16
    pixels          []Color
    animation_queue []animation
}


func GenBuffer(length uint16) (Buffer) {
    b := Buffer{}
    b.length = length

    return b
}

func (b *Buffer) GetPixels() ([]Color) {
    return b.pixels
}

func (b *Buffer) GetAnimationQueue() ([]animation) {
    return b.animation_queue
}

func XOR(b1 *Buffer, b2 *Buffer) {
}

func (b *Buffer) ApplyQueue() {
}

func (b *Buffer) AddAnimation(a animation, as ...animation) {
    b.animation_queue = append(b.animation_queue, a)
    for _, v := range as {
        b.animation_queue = append(b.animation_queue, v)
    }
}
