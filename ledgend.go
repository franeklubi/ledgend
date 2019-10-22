package ledgend


type rgb struct {
    R, G, B uint8
}

type animation struct {
    name                        string
    start_colour, end_colour    rgb
}

type Buffer struct {
    length          uint16
    pixels          []rgb
    animation_queue []animation
}


func (b *Buffer) GetBuffer() ([]rgb) {
    return b.pixels;
}

func (b *Buffer) GetAnimationQueue() ([]animation) {
    return b.animation_queue;
}

func XOR(b1 *Buffer, b2 *Buffer) {
}
