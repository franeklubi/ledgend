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
    var max_pixels float64
    if ( a.Direction ) {
        max_pixels = math.Floor((buffer_length_float64-start_index)*a.Length)
    } else {
        max_pixels = math.Floor(start_index*a.Length)
    }
    if ( max_pixels == 0 ) {
        return true
    }

    // time since start in milliseconds
    since_start := float64(time.Since(a.Start).Milliseconds())
    done := time.Now().After(a.Start.Add(a.Duration))

    time_length_multiplier := since_start / float64(a.Duration.Milliseconds())
    if ( time_length_multiplier > 1 ) {
        time_length_multiplier = 1
    }

    // length corrected for time
    curr_length := max_pixels * time_length_multiplier


    for x := 0; x < int(curr_length); x++ {
        // multipler for gradient
        m := float64(x) / max_pixels

        var index_to_draw int
        if ( a.Direction ) {
            index_to_draw = x + int(start_index)
        } else {
            index_to_draw = int(start_index) - x - 1
        }

        b.pixels[index_to_draw] = Gradient(a.Start_colour, a.End_colour, m)
    }

    return done
}

func Gradient(a, b Color, m float64) (Color) {
    R := float64(a.R) - math.Floor( (float64(a.R)-float64(b.R))*m )
    G := float64(a.G) - math.Floor( (float64(a.G)-float64(b.G))*m )
    B := float64(a.B) - math.Floor( (float64(a.B)-float64(b.B))*m )

    return Color{uint8(R), uint8(G), uint8(B)}
}
