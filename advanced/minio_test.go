package advanced


var bp = NewBytePoolCap(500, 1024, 1024)
var sp = &sync.Pool{
    New: func() interface{} {
        return make([]byte, 1024, 1024)
    },
}