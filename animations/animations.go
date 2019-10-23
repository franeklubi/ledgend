package animations

import (
    "time"
    . "github.com/franeklubi/ledgend"
)


func Sweep(
    direction bool,
    start_pos, length float64,
    start_colour, end_colour Color,
    duration time.Duration,
) (Animation) {
    return Animation{
        direction,
        start_pos, length,
        start_colour, end_colour,
        duration, time.Now(),
    }
}


func FromMiddleFullSweep(
    start_colour, end_colour Color,
    duration time.Duration,
) (Animation, Animation) {

    a := Sweep(
        false, 0.5, 1,
        start_colour, end_colour,
        duration,
    )

    b := a
    b.Direction = true

    return a, b
}


func Pulse(
    direction bool,
    start_pos, length float64,
    start_colour_a, end_colour_a Color,
    start_colour_b, end_colour_b Color,
    duration, duration_back time.Duration,
) (Animation, Animation) {
    a := Sweep(
        direction, start_pos, length,
        start_colour_a, end_colour_a,
        duration,
    )


    b_start_pos := start_pos + (length * (1 - start_pos))
    b_length := 1-start_pos
    if (!direction) {
        b_length = b_start_pos
        b_start_pos = 1 - b_start_pos
    }

    b := Sweep(
        !direction, b_start_pos, b_length,
        start_colour_b, end_colour_b,
        duration_back,
    )
    b.Start = b.Start.Add(duration)

    return a, b
}
