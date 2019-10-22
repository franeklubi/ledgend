package animations

import (
    "time"
    . "github.com/franeklubi/ledgend"
)

func Sweep(dir bool, len float32, s_c, e_c Color, d time.Duration) (Animation) {
    return Animation{dir, len, s_c, e_c, d, time.Now()}
}
