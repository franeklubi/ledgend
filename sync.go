package ledgend

import (
    "sync"
)


type SyncBuffer struct {
    b   Buffer
    m   sync.RWMutex
}


func GenSyncBuffer(length uint16) (SyncBuffer) {
    var sb SyncBuffer
    sb.b = GenBuffer(length)
    return sb
}


func (sb *SyncBuffer) AddAnimation(a ...Animation) {
    sb.m.Lock()
    defer sb.m.Unlock()

    sb.b.AddAnimation(a...)
}


func (sb *SyncBuffer) ApplyQueue() {
    sb.m.Lock()
    defer sb.m.Unlock()

    sb.b.ApplyQueue()
}


func (sb *SyncBuffer) ClearQueue() {
    sb.m.Lock()
    defer sb.m.Unlock()

    sb.b.ClearQueue()
}


func (sb *SyncBuffer) GetAnimationQueue() ([]Animation) {
    sb.m.RLock()
    defer sb.m.RUnlock()

    return sb.b.GetAnimationQueue()
}


func (sb *SyncBuffer) GetPixels() ([]Color) {
    sb.m.RLock()
    defer sb.m.RUnlock()

    return sb.b.GetPixels()
}
