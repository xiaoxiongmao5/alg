package main

import(
	"fmt"
	slink "xj/chapter1_linklist/level1/single_linklist"
	dlink "xj/chapter1_linklist/level1/double_linklist"
	rslink "xj/chapter1_linklist/level1/ring_singlelink"
)

func TestSinglelink(){
	fmt.Println("————打印空链表————")
	catHead := &slink.SingleCat{}
	slink.ShowLink(catHead)

	cat1 := &slink.SingleCat{
		Id:1,
	}
	cat2 := &slink.SingleCat{
		Id:2,
	}
	cat3 := &slink.SingleCat{
		Id:3,
	}
	cat4 := &slink.SingleCat{
		Id:4,
	}
	cat5 := &slink.SingleCat{
		Id:5,
	}
	fmt.Println("————链表添加 尾结点1————")
	slink.AddLinkNode(catHead, cat1, 0)
	slink.ShowLink(catHead)
	fmt.Println("————链表添加 不存在位置的结点2————")
	slink.AddLinkNode(catHead, cat2, 5)
	slink.ShowLink(catHead)
	fmt.Println("————链表添加 头结点3————")
	slink.AddLinkNode(catHead, cat3, 1)
	slink.ShowLink(catHead)
	fmt.Println("————链表添加 第二个结点4————")
	slink.AddLinkNode(catHead, cat4, 2)
	slink.ShowLink(catHead)
	fmt.Println("————链表添加 不存在位置的结点5————")
	slink.AddLinkNode(catHead, cat5, 15)
	slink.ShowLink(catHead)

	fmt.Println("————链表删除 最后一个节点————")
	slink.DelLinkNode(catHead, 0)
	slink.ShowLink(catHead)
	fmt.Println("————链表删除 第一个节点————")
	slink.DelLinkNode(catHead, 1)
	slink.ShowLink(catHead)
	fmt.Println("————链表删除 不存在的节点————")
	slink.DelLinkNode(catHead, 18)
	slink.ShowLink(catHead)
	fmt.Println("————链表删除 第二个节点————")
	slink.DelLinkNode(catHead, 2)
	slink.ShowLink(catHead)
	fmt.Println("————链表删除 第二个节点————")
	slink.DelLinkNode(catHead, 2)
	slink.ShowLink(catHead)
	fmt.Println("————链表删除 最后一个节点————")
	slink.DelLinkNode(catHead, 0)
	slink.ShowLink(catHead)
	fmt.Println("————链表删除 第一个节点————")
	slink.DelLinkNode(catHead, 1)
	slink.ShowLink(catHead)
}

func TestDoublelink(){
	fmt.Println("————打印空链表————")
	catHead := &dlink.SingleCat{}
	dlink.ShowLink(catHead)

	cat1 := &dlink.SingleCat{
		Id:1,
	}
	cat2 := &dlink.SingleCat{
		Id:2,
	}
	cat3 := &dlink.SingleCat{
		Id:3,
	}
	cat4 := &dlink.SingleCat{
		Id:4,
	}
	cat5 := &dlink.SingleCat{
		Id:5,
	}
	fmt.Println("————链表添加 尾结点1————")
	dlink.AddLinkNode(catHead, cat1, 0)
	dlink.ShowLink(catHead)
	fmt.Println("————链表添加 不存在位置的结点2————")
	dlink.AddLinkNode(catHead, cat2, 5)
	dlink.ShowLink(catHead)
	fmt.Println("————链表添加 头结点3————")
	dlink.AddLinkNode(catHead, cat3, 1)
	dlink.ShowLink(catHead)
	fmt.Println("————链表添加 第二个结点4————")
	dlink.AddLinkNode(catHead, cat4, 2)
	dlink.ShowLink(catHead)
	fmt.Println("————链表添加 不存在位置的结点5————")
	dlink.AddLinkNode(catHead, cat5, 15)
	dlink.ShowLink(catHead)
	fmt.Println("————链表逆序打印————")
	dlink.ShowLink2(catHead)
	
	fmt.Println("————链表删除 最后一个节点————")
	dlink.DelLinkNode(catHead, 0)
	dlink.ShowLink(catHead)
	fmt.Println("————链表删除 第一个节点————")
	dlink.DelLinkNode(catHead, 1)
	dlink.ShowLink(catHead)
	fmt.Println("————链表删除 不存在的节点————")
	dlink.DelLinkNode(catHead, 18)
	dlink.ShowLink(catHead)
	fmt.Println("————链表删除 第二个节点————")
	dlink.DelLinkNode(catHead, 2)
	dlink.ShowLink(catHead)
	fmt.Println("————链表删除 第二个节点————")
	dlink.DelLinkNode(catHead, 2)
	dlink.ShowLink(catHead)
	fmt.Println("————链表删除 最后一个节点————")
	dlink.DelLinkNode(catHead, 0)
	dlink.ShowLink(catHead)
	fmt.Println("————链表删除 第一个节点————")
	dlink.DelLinkNode(catHead, 1)
	dlink.ShowLink(catHead)
}

func TestRingSinglelink(){
	fmt.Println("————打印空链表————")
	catHead := &rslink.SingleCat{}
	rslink.ShowLink(catHead)

	cat1 := &rslink.SingleCat{
		Id:1,
	}
	cat2 := &rslink.SingleCat{
		Id:2,
	}
	cat3 := &rslink.SingleCat{
		Id:3,
	}
	cat4 := &rslink.SingleCat{
		Id:4,
	}
	cat5 := &rslink.SingleCat{
		Id:5,
	}
	fmt.Println("————链表添加 尾结点1————")
	catHead = rslink.AddLinkNode(catHead, cat1, 0)
	rslink.ShowLink(catHead)
	fmt.Println("————链表添加 不存在位置的结点2————")
	catHead = rslink.AddLinkNode(catHead, cat2, 5)
	rslink.ShowLink(catHead)
	fmt.Println("————链表添加 头结点3————")
	catHead = rslink.AddLinkNode(catHead, cat3, 1)
	rslink.ShowLink(catHead)
	fmt.Println("————链表添加 第二个结点4————")
	catHead = rslink.AddLinkNode(catHead, cat4, 2)
	rslink.ShowLink(catHead)
	fmt.Println("————链表添加 尾结点5————")
	catHead = rslink.AddLinkNode(catHead, cat5, 0)
	rslink.ShowLink(catHead)

	fmt.Println("————链表删除 最后一个节点————")
	catHead = rslink.DelLinkNode(catHead, 0)
	rslink.ShowLink(catHead)
	fmt.Println("————链表删除 第一个节点————")
	catHead = rslink.DelLinkNode(catHead, 1)
	rslink.ShowLink(catHead)
	fmt.Println("————链表删除 不存在的节点————")
	catHead = rslink.DelLinkNode(catHead, 18)
	rslink.ShowLink(catHead)
	fmt.Println("————链表删除 第二个节点————")
	catHead = rslink.DelLinkNode(catHead, 2)
	rslink.ShowLink(catHead)
	fmt.Println("————链表删除 第二个节点————")
	catHead = rslink.DelLinkNode(catHead, 2)
	rslink.ShowLink(catHead)
	fmt.Println("————链表删除 最后一个节点————")
	catHead = rslink.DelLinkNode(catHead, 0)
	rslink.ShowLink(catHead)
	fmt.Println("————链表删除 第一个节点————")
	catHead = rslink.DelLinkNode(catHead, 1)
	rslink.ShowLink(catHead)
}

func main() {
	fmt.Println("hi, linklist level1")

	// fmt.Println("================>下面测试单向链表")
	// TestSinglelink()

	// fmt.Println("================>下面测试双向链表")
	// TestDoublelink()

	fmt.Println("================>下面测试环形单向链表")
	TestRingSinglelink()
	
}