package ledgend


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


func XORPixels(b1 *Buffer, b2 *Buffer) {
}


func (b *Buffer) GetAnimationQueue() ([]Animation) {
    return b.animation_queue
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
