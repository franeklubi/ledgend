package ledgend


import (
    "errors"
)


type Buffer struct {
    length          uint16
    pixels          []Color
    animation_queue []Animation
}


type Change struct {
    Index   uint16
    Pixel   Color
}


// GenBuffer generates an empty buffer
func GenBuffer(length uint16) (Buffer) {
    b := Buffer{}
    b.length = length

    b.pixels = make([]Color, length)

    return b
}


// GetPixels copies and returns pixels data from Buffer
func (b *Buffer) GetPixels() ([]Color) {
    copied_pixels := make([]Color, len(b.pixels))
    copy(copied_pixels, b.pixels)

    return copied_pixels
}


// XORPixels returns Changes between Color slices
func XORPixels(p1 []Color, p2 []Color) ([]Change, error) {
    if ( len(p1) != len(p2) ) {
        return []Change{}, errors.New("Pixel slices must be the same length!")
    }

    var changes []Change
    for index, pixel := range p2 {
        if ( p1[index] != pixel ) {
            changes = append(changes, Change{uint16(index), pixel})
        }
    }

    return changes, nil
}


// GetAnimationQueue returns Buffer's animation queue
func (b *Buffer) GetAnimationQueue() ([]Animation) {
    return b.animation_queue
}


// ApplyQueue applies every animation from Buffer's animation_queue in order
//
// If the animation is finished, it removes it from the animation_queue
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


// AddAnimation adds an Animation to Buffer's animation queue
func (b *Buffer) AddAnimation(a Animation, as ...Animation) {
    b.animation_queue = append(b.animation_queue, a)

    for _, animation := range as {
        b.animation_queue = append(b.animation_queue, animation)
    }
}
