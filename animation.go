package ledgend


import (
    "time"
    "math"
)


type Animation struct {
    Direction           bool
    Start_pos, Length   float64     // float between 0 and 1
    Start_col, End_col  Color
    Duration            time.Duration
    Start               time.Time
}


// applyAnimation renders an Animation on Buffer's pixels
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

        b.pixels[index_to_draw] = Gradient(a.Start_col, a.End_col, m)
    }

    return done
}
