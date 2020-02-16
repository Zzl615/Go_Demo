package heap

type MaxHeap struct {
	heap  []int
	h_len int
	h_cap int
}

/**
 *  NewMaxHeap returns an initialized MaxHeap.
 */
func NewMaxHeap(k int) *MaxHeap {
	return &MaxHeap{
		heap:  make([]int, 0, k),
		h_len: 0,
		h_cap: k,
	}
}

/**
 *  InsertElem：Insert an element into the max heap.
 */
func (this *MaxHeap) InsertElem(num int) bool {
	if this.h_cap == this.h_len {
		return false
	}
	this.heap[this.h_len] = num
	this.h_len++
	this.heap = this.heapifyBottom(this.heap, h_len)
}

/**
 *  DeleteElem：Delete an element from the max heap.
 */
func (this *MaxHeap) DeleteElem(num int) bool {
	if this.h_len == 0 {
		return false
	}
	this.heap[1], this.heap[this.h_len] = this.heap[this.h_len], this.heap[1]
	this.h_len--
	this.heapifyTop(&this.heap, this.h_len)
	return true
}

/**
 * buildHeap：
 *
 */
func (this *MaxHeap) BuildHeap(array []int, n int) {
	for i := (n - 1) / 2; i >= 0; i-- {
		this.heapifyTop(array, n, i)
	}
}

/**
 * heapifyTop
 */
func (this *MaxHeap) heapifyTop(array []int, n int, i int) *MaxHeap {
	for i = n {
		var maxPos := i
		if array[2*i+1] > array[i] {
			maxPos = 2*i+1
		} 
		if array[2*i+2] > array[i] {
			maxPos = 2*i+2
		}
		if maxPos != i {
			array[maxPos], array[i] = array[i], array[maxPos]
		}
		i ++
	}
	return &array
}

/**
 * heapifyBottom
 */
func (this *MaxHeap) heapifyBottom(arr []int, k int) *MaxHeap {
	// for i:= k; i >= 0 && arr[i] > arr[i/2]; i = i/2 {
	// 	arr[i], arr[i/2] = arr[i/2], arr[i]
	// }
	// return arr
}
