package ledgend

import (
    "time"
    "math"
)


type Color struct {
    R, G, B uint8
}

type Animation struct {
    Direction                   bool
    Start_pos, Length           float64     // float between 0 and 1
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
    var updated_queue []Animation

    for _, animation := range b.animation_queue {
        done := b.applyAnimation(animation)

        if ( !done ) {
            updated_queue = append(updated_queue, animation)
        }
    }

    b.animation_queue = updated_queue
}

func (b *Buffer) AddAnimation(a Animation, as ...Animation) {
    b.animation_queue = append(b.animation_queue, a)

    for _, animation := range as {
        b.animation_queue = append(b.animation_queue, animation)
    }
}

func (b *Buffer) applyAnimation(a Animation) (bool) {
    buffer_length_float64 := float64(b.length)

    start_index := math.Floor(buffer_length_float64*a.Start_pos)

    // determine how many pixels the animation covers at max
    max_pixels := math.Floor((buffer_length_float64-start_index)*a.Length)

    // time since start in milliseconds
    since_start := float64(time.Since(a.Start).Milliseconds())
    done := time.Now().After(a.Start.Add(a.Duration))

    time_length_multiplier := since_start / float64(a.Duration.Milliseconds())

    // length corrected for time
    curr_length := max_pixels * time_length_multiplier

    for x := 0; x < int(curr_length); x++ {
        b.pixels[x+int(start_index)] = a.Start_colour
    }

    return done
}
