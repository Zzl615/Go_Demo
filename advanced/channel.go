package advanced

import (
	"parallelReader"
)

// Read reads from readers in parallel. Returns p.dataBlocks number of bufs.
func (p *parallelReader) Read(dst [][]byte) ([][]byte, error) {
	
	newBuf := dst
    //省略不太相关代码
    var newBufLK sync.RWMutex

    //省略无关
    //channel开始创建，要发挥作用了。这里记住几个数字：
    //readTriggerCh大小是10，p.dataBlocks大小是5
    readTriggerCh := make(chan bool, len(p.readers))
    for i := 0; i < p.dataBlocks; i++ {
        // Setup read triggers for p.dataBlocks number of reads so that it reads in parallel.
        readTriggerCh <- true
    }

    healRequired := int32(0) // Atomic bool flag.
    readerIndex := 0
    var wg sync.WaitGroup
    // readTrigger 为 true, 意味着需要用disk.ReadAt() 读取下一个数据
    // readTrigger 为 false, 意味着读取成功了，不再需要读取
    for readTrigger := range readTriggerCh {
        newBufLK.RLock()
        canDecode := p.canDecode(newBuf)
        newBufLK.RUnlock()
        //判断是否有5个成功的，如果有，退出for循环
        if canDecode {
            break
        }
        //读取次数上限，不能大于10
        if readerIndex == len(p.readers) {
            break
        }
        //成功了，退出本次读取
        if !readTrigger {
            continue
        }
        wg.Add(1)
        //并发读取数据
        go func(i int) {
            defer wg.Done()
            //省略不太相关代码
            _, err := rr.ReadAt(p.buf[bufIdx], p.offset)
            if err != nil {
                //省略不太相关代码
                // 失败了，标记为true，触发下一个读取.
                readTriggerCh <- true
                return
            }
            newBufLK.Lock()
            newBuf[bufIdx] = p.buf[bufIdx]
            newBufLK.Unlock()
            // 成功了，标记为false，不再读取
            readTriggerCh <- false
        }(readerIndex)
        //控制次数，同时用来作为索引获取和存储数据
        readerIndex++
    }
    wg.Wait()

    //最终结果判断，如果OK了就正确返回，如果有失败的，返回error信息。
    if p.canDecode(newBuf) {
        p.offset += p.shardSize
        if healRequired != 0 {
            return newBuf, errHealRequired
        }
        return newBuf, nil
    }

    return nil, errErasureReadQuorum
}

// 最终成功模式

// 第一种思路：
// 先并发获取，存放起来，然后再一个个判断是否获取成功，
// 如果有的没有成功再重新获取,而且获取的文件不能重复。
// 这种方式是取到结果后进行判断是否成功，然后根据情况再决定是否重新获取,
// 要去重，要判断，业务逻辑比较复杂。

// 第二种思路：
// 并发的时候就保证成功，里面可能是个for循环，直到成功为止
// 然后再返回结果。这种思路缺陷也很明显
// 如果这个文件损坏，那么就会一直死循环下去，要避免死循环，就要加上重试次数。

// MinIO的实现方式
// 比较巧妙，它也是多协程，但是发现如果有文件读取不成功，
// 他会通过channel的方式标记，换一个文件读取。
// 因为一共10个文件呢，这个不行，换一个，不能在一个文件上等死，只要成功读取5个就可以了。
