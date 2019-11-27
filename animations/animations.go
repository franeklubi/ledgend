package animations

import (
    "time"
    . "github.com/franeklubi/ledgend"
)


func Sweep(
    direction bool,
    start_pos, length float64,
    start_col, end_col Color,
    duration time.Duration,
    start_time time.Time,
) (Animation) {
    return Animation{
        direction,
        start_pos, length,
        start_col, end_col,
        duration, start_time,
    }
}


func FromMiddleFullSweep(
    start_col, end_col Color,
    duration time.Duration,
    start_time time.Time,
) (Animation, Animation) {

    a := Sweep(
        false, 0.5, 1,
        start_col, end_col,
        duration, start_time,
    )

    b := a
    b.Direction = true

    return a, b
}


func Pulse(
    direction bool,
    start_pos, length float64,
    start_col_a, end_col_a,
    start_col_b, end_col_b Color,
    duration, duration_back time.Duration,
    start_time time.Time,
) (Animation, Animation) {
    a := Sweep(
        direction, start_pos, length,
        start_col_a, end_col_a,
        duration, start_time,
    )


    b_start_pos := start_pos + (length * (1 - start_pos))
    b_length := 1-start_pos
    if (!direction) {
        b_length = b_start_pos
        b_start_pos = 1 - b_start_pos
    }

    b := Sweep(
        !direction, b_start_pos, b_length,
        start_col_b, end_col_b,
        duration_back, start_time,
    )
    b.Start = b.Start.Add(duration)

    return a, b
}


func Strobo(
    start_col_a, end_col_a,
    start_col_b, end_col_b Color,
    duration, interval time.Duration,
    start_time time.Time,
) ([]Animation) {
    passes := int(duration.Milliseconds()/interval.Milliseconds())

    var (
        anims           []Animation
        blank           Color
    )

    if ( passes%2 != 0 ) {
        passes++
    }

    strobe := Sweep(
        true, 0, 1,
        start_col_a, end_col_a,
        time.Millisecond, start_time,
    )

    for x := 0; x < passes; x++ {

        if ( x%2 == 0 ) {
            g := float64(x)/float64(passes)
            strobe.Start_col = Gradient(start_col_a, start_col_b, g)
            strobe.End_col = Gradient(end_col_a, end_col_b, g)
        } else {
            strobe.Start_col = blank
            strobe.End_col = blank
        }

        anims = append(anims, strobe)

        strobe.Start = strobe.Start.Add(interval)
    }

    return anims
}


func GradientOverTime(
    direction bool,
    start_pos, length float64,
    start_col_a, end_col_a,
    start_col_b, end_col_b Color,
    duration, interval time.Duration,
    start_time time.Time,
) ([]Animation) {
    passes := int(duration.Milliseconds()/interval.Milliseconds())

    var anims []Animation

    strobe := Sweep(
        direction, start_pos, length,
        start_col_a, end_col_a,
        time.Millisecond, start_time,
    )

    for x := 0; x < passes; x++ {
        g := float64(x)/float64(passes)
        strobe.Start_col = Gradient(start_col_a, start_col_b, g)
        strobe.End_col = Gradient(end_col_a, end_col_b, g)

        anims = append(anims, strobe)

        strobe.Start = strobe.Start.Add(interval)
    }
}
