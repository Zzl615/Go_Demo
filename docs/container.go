// List是列表，一种非连续存储的容器，由多个节点组成，节点通过一些变量记录彼此之间的关系。list有多种实现方法，如单向链表、双向链表等。
// Go语言中的list的实现原理是双向链表，list能高效地进行任意位置的元素插入和删除操作。
// Golang的标准库提供了高级的数据结构List，具体在包container/list。
// container/list包里主要有两个数据结构类型：“Element”、“List”；
// Element类型代表双向链表中的一个元素，相当于C++里面的“iterator”
// List代表一个双向链表。List零值为一个空的、可用的链表。
// Element有Prev和Next方法用于得到前一个或者下一个element，element可以直接调用value属性；

 

// 声明list
// 格式：变量名:= list.New()
//       var 变量名 list.List
// list与切片、Map不一样，没有具体元素类型的限制。
// 在java、c++等里面，list的成员必须是同一个数据类型，但是Go语言中却允许list里插入任意类型成员。
// 建议使用New()实现list。

 

// list的声明及添加数据

// 为了等下的代码能更直观的运行，先写两个print的函数

func printList(i string, l list.List) {
	fmt.Println(i)
	fmt.Printf("类型：%T, 值：%v, 长度：%d \n", l, l, l.Len())
}
 
func printIterateList(l list.List){
	for e:=l.Front();e!=nil; e= e.Next(){
		fmt.Printf("%T ,%v \t", e.Value, e.Value)
	}
}

// 用第一种方式声明及添加数据

func testListNew() {
	var list1 list.List
	printList("测试声明list", list1)
	//测试声明list
	// 类型：list.List, 值：{{<nil> <nil> <nil> <nil>} 0}, 长度：0
 
	//添加数据
	list1.PushFront(true)
	list1.PushFront("aaaa")
	list1.PushFront(1.23456)
	list1.PushFront(123)
	//尾部添加数据
	list1.PushBack("一二三")
	printList("添加数据以后list", list1)
	//添加数据以后list
	// 类型：list.List, 值：{{0xc000062360 0xc000062390 <nil> <nil>} 5}, 长度：5
 
	printIterateList(list1)
	//int ,123 	float64 ,1.23456 	string ,aaaa 	bool ,true 	string ,一二三
}

// 第二种方式声明及其他相关函数
func testList1(){
	list1 := list.New()
	//用list.New()声明的list，返回的是一个指针，所以必须加上一个*号
	printList("",*list1)
	list1.PushBack("bbbbb")
	list1.PushFront(1)
	element1 := list1.PushBack("ccccc")
	//在element1后面插入数据
	list1.InsertAfter("aaaaa",element1)
	//在element1前面插入数据
	list1.InsertBefore(50,element1)
	printIterateList(*list1)
 
	fmt.Println()
	//移除element1
	list1.Remove(element1)
	printIterateList(*list1)
 
	list2 := list.New()
	list2.PushFront("1234")
	list2.PushBack("5678")
 
	fmt.Println()
	//在list1前面添加list2列表
	list1.PushFrontList(list2)
	printIterateList(*list1)
 
	list3 := list.New()
	list3.PushFront("abcd")
	list3.PushBack("efgh")
 
	fmt.Println()
	//在lisi1后面添加list3、
	list1.PushBackList(list3)
	printIterateList(*list1)
}
 

// list的遍历：先建一个变量

var list2 list.List
// 然后开始遍历该list

//遍历list，顺序遍历
func iterateList1() {
	for e := list2.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
 
//遍历list，逆序遍历
func iterateList2() {
	for e := list2.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
 

// list是值类型：用一个函数来证明list是值类型。

func testList2(l list.List){
	l.PushFront("一")
	l.PushBack("三")
	printIterateList(l)
 
	fmt.Println(l.Len())
}
 
 
//说明list是值类型
func main() {
	printIterateList(list2)
	fmt.Println()
	testList2(list2)
	list2.PushFront("二")
	printIterateList(list2)
	fmt.Println()
	testList2(list2)
}
 