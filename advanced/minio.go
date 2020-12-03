package advanced

// https://mp.weixin.qq.com/s?__biz=MzI3MjU4Njk3Ng==&mid=2247484598&idx=1&sn=ea91646ad078518e9ffbc39994abe424&chksm=eb310539dc468c2f81bc79f2ff55e3e8e314506137746ed6bb785420d9ee49f1a289a67897d6&scene=178&cur_album_id=1368991005374234626#rd

type BytePoolCap struct {
	c    chan []byte
	w    int
	wcap int
}

func (bp *BytePoolCap) Get() (b []byte) {
	select {
	case b = <-bp.c:
	// reuse existing buffer
	default:
		// create new buffer
		if bp.wcap > 0 {
			b = make([]byte, bp.w, bp.wcap)
		} else {
			b = make([]byte, bp.w)
		}
	}
	return
}

func (bp *BytePoolCap) Put(b []byte) {
	select {
	case bp.c <- b:
		// buffer went back into pool
	default:
		// buffer didn't go back into pool, just discard
	}
}

func NewBytePoolCap(maxSize int, width int, capwidth int) (bp *BytePoolCap) {
	// 工厂函数，用于生成*BytePoolCap
	return &BytePoolCap{
		c:    make(chan []byte, maxSize),
		w:    width,
		wcap: capwidth,
	}
}
