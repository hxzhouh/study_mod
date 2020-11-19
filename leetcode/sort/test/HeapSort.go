package test

type Heap struct {
	a        [11]int // 堆内元素
	capacity int     // 堆的容量
	count    int     // 已经存储的数据
}

func InitHeap() *Heap {
	a := [11]int{}
	return &Heap{
		a:        a,
		capacity: 10, // 因为第一个位置没有存储数据
		count:    0,
	}
}

func (t *Heap) Insert(data int) {
	if t.count >= 10 {
		return
	}
	t.count += 1
	t.a[t.count] = data //先写到最后
	i := t.count

	for i/2 > 0 && t.a[i] > t.a[i/2] {
		t.swap(i, i/2)
		i = i / 2
	}
}

func (t *Heap) swap(i, i1 int) {
	temp := t.a[i]
	t.a[i] = t.a[i1]
	t.a[i1] = temp
}

//删除堆顶元素
func (t *Heap) removeMax() {
	t.a[1] = t.a[t.count]
	t.count -= 1 //将最后一个元素放到堆顶，然后使用自上而下堆化过程
	t.heapify()
}
func (t *Heap) heapify() {
	i := 1
	for {
		maxPos := i
		if i*2 <= t.count && t.a[i] < t.a[i*2] { // 如果左节点大于 根节点
			maxPos = i * 2
		}
		if i*2+1 <= t.count && t.a[maxPos] < t.a[i*2+1] { // 如果右节点大于根节点
			maxPos = i*2 + 1
		}
		if maxPos == i {
			break
		}
		t.swap(i, maxPos)
		i = maxPos
	}
}

func (t *Heap) Delete(i int) {

}
