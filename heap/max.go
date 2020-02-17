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
	this.heap = this.HeapifyBottom(this.heap, h_len)
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
	this.HeapifyTop(&this.heap, this.h_len, 0)
	return true
}

/**
 * BuildHeap:
 *
 */
func (this *MaxHeap) BuildHeap(array []int, n int) *maxHeap{
	for i := (n - 1) / 2; i >= 0; i-- {
		this.HeapifyTop(array, n, i)
	}
	return &array
}

/**
 * heapifyTop
 */
func (this *MaxHeap) HeapifyTop(array []int, n int, i int) *maxHeap {
	var maxPos = -1
	var maxParent = (n-1)/2 
	for ; maxPos != i; i = maxPos {
		maxPos := i
		if i > maxParent && array[2*i+1] > array[i] { 
			maxPos = 2*i + 1
		}
		if i > maxParent && array[2*i+2] > array[i] {
			maxPos = 2*i + 2
		}
		if maxPos != i {
			array[maxPos], array[i] = array[i], array[maxPos]
		}
	}
	return &array
}

/**
 * heapifyBottom
 */
func (this *MaxHeap) HeapifyBottom(arr []int, i int) *maxHeap {
	for i > 0 && arr[i] > arr[parentPos]{
		parentPos = (i-1)/2
		if arr[i] > arr[parentPos] {
			arr[i], arr[i/2] = arr[i/2], arr[i]
		}
		i = parentPos
	}
	return &arr
}
