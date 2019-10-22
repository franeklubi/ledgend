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
