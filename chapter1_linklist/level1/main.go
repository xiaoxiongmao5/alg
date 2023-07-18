package main

import(
	"fmt"
	slink "xj/chapter1_linklist/level1/single_linklist"
	dlink "xj/chapter1_linklist/level1/double_linklist"
	rslink "xj/chapter1_linklist/level1/ring_singlelink"
	rdlink "xj/chapter1_linklist/level1/ring_doublelink"
)

func TestSinglelink(){
	fmt.Println("————打印空链表————")
	catHead := &slink.Cat{}
	slink.ShowLink(catHead)

	cat1 := &slink.Cat{
		Id:1,
	}
	cat2 := &slink.Cat{
		Id:2,
	}
	cat3 := &slink.Cat{
		Id:3,
	}
	cat4 := &slink.Cat{
		Id:4,
	}
	cat5 := &slink.Cat{
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
	catHead := &dlink.Cat{}
	dlink.ShowLink(catHead)

	cat1 := &dlink.Cat{
		Id:1,
	}
	cat2 := &dlink.Cat{
		Id:2,
	}
	cat3 := &dlink.Cat{
		Id:3,
	}
	cat4 := &dlink.Cat{
		Id:4,
	}
	cat5 := &dlink.Cat{
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
	catHead := &rslink.Cat{}
	rslink.ShowLink(catHead)

	cat1 := &rslink.Cat{
		Id:1,
	}
	cat2 := &rslink.Cat{
		Id:2,
	}
	cat3 := &rslink.Cat{
		Id:3,
	}
	cat4 := &rslink.Cat{
		Id:4,
	}
	cat5 := &rslink.Cat{
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

func TestRingDoublelink() {
	fmt.Println("————打印空链表————")
	catHead := &rdlink.Cat{}
	rdlink.ShowLink(catHead)

	cat1 := &rdlink.Cat{
		Id:1,
	}
	cat2 := &rdlink.Cat{
		Id:2,
	}
	cat3 := &rdlink.Cat{
		Id:3,
	}
	cat4 := &rdlink.Cat{
		Id:4,
	}
	cat5 := &rdlink.Cat{
		Id:5,
	}
	fmt.Println("————链表添加 尾结点1————")
	catHead = rdlink.AddLinkNode(catHead, cat1, 0)
	rdlink.ShowLink(catHead)
	fmt.Println("————链表添加 不存在位置的结点2————")
	catHead = rdlink.AddLinkNode(catHead, cat2, 5)
	rdlink.ShowLink(catHead)
	fmt.Println("————链表添加 头结点3————")
	catHead = rdlink.AddLinkNode(catHead, cat3, 1)
	rdlink.ShowLink(catHead)
	fmt.Println("————链表添加 第二个结点4————")
	catHead = rdlink.AddLinkNode(catHead, cat4, 2)
	rdlink.ShowLink(catHead)
	fmt.Println("————链表添加 尾结点5————")
	catHead = rdlink.AddLinkNode(catHead, cat5, 0)
	rdlink.ShowLink(catHead)
	fmt.Println("————链表逆序打印————")
	rdlink.ShowLink2(catHead)

	fmt.Println("————链表删除 最后一个节点————")
	catHead = rdlink.DelLinkNode(catHead, 0)
	rdlink.ShowLink(catHead)
	fmt.Println("————链表删除 第一个节点————")
	catHead = rdlink.DelLinkNode(catHead, 1)
	rdlink.ShowLink(catHead)
	fmt.Println("————链表删除 不存在的节点————")
	catHead = rdlink.DelLinkNode(catHead, 18)
	rdlink.ShowLink(catHead)
	fmt.Println("————链表删除 第二个节点————")
	catHead = rdlink.DelLinkNode(catHead, 2)
	rdlink.ShowLink(catHead)
	fmt.Println("————链表删除 第二个节点————")
	catHead = rdlink.DelLinkNode(catHead, 2)
	rdlink.ShowLink(catHead)
	fmt.Println("————链表删除 最后一个节点————")
	catHead = rdlink.DelLinkNode(catHead, 0)
	rdlink.ShowLink(catHead)
	fmt.Println("————链表删除 第一个节点————")
	catHead = rdlink.DelLinkNode(catHead, 1)
	rdlink.ShowLink(catHead)
}

func createSList() *slink.Cat{
	list := &slink.Cat{}
	cat1 := &slink.Cat{
        Id : 1,
    }
	cat2 := &slink.Cat{
        Id : 2,
    }
	cat30 := &slink.Cat{
        Id : 30,
    }
	cat40 := &slink.Cat{
        Id : 40,
    }
	cat50 := &slink.Cat{
        Id : 50,
    }
	slink.AddLinkNode(list, cat1, 0)
	slink.AddLinkNode(list, cat2, 0)
	slink.AddLinkNode(list, cat30, 0)
	slink.AddLinkNode(list, cat40, 0)
	slink.AddLinkNode(list, cat50, 0)
	return list
}
func createDList() *dlink.Cat{
	list := &dlink.Cat{}
	cat1 := &dlink.Cat{
        Id : 1,
    }
	cat2 := &dlink.Cat{
        Id : 2,
    }
	cat30 := &dlink.Cat{
        Id : 30,
    }
	cat40 := &dlink.Cat{
        Id : 40,
    }
	cat50 := &dlink.Cat{
        Id : 50,
    }
	dlink.AddLinkNode(list, cat1, 0)
	dlink.AddLinkNode(list, cat2, 0)
	dlink.AddLinkNode(list, cat30, 0)
	dlink.AddLinkNode(list, cat40, 0)
	dlink.AddLinkNode(list, cat50, 0)
	return list
}
func createRSList() *rslink.Cat{
	list := &rslink.Cat{}
	cat1 := &rslink.Cat{
        Id : 1,
    }
	cat2 := &rslink.Cat{
        Id : 2,
    }
	cat30 := &rslink.Cat{
        Id : 30,
    }
	cat40 := &rslink.Cat{
        Id : 40,
    }
	cat50 := &rslink.Cat{
        Id : 50,
    }
	rslink.AddLinkNode(list, cat1, 0)
	rslink.AddLinkNode(list, cat2, 0)
	rslink.AddLinkNode(list, cat30, 0)
	rslink.AddLinkNode(list, cat40, 0)
	rslink.AddLinkNode(list, cat50, 0)
	return list
}
func createRDList() *rdlink.Cat{
	list := &rdlink.Cat{}
	cat1 := &rdlink.Cat{
        Id : 1,
    }
	cat2 := &rdlink.Cat{
        Id : 2,
    }
	cat30 := &rdlink.Cat{
        Id : 30,
    }
	cat40 := &rdlink.Cat{
        Id : 40,
    }
	cat50 := &rdlink.Cat{
        Id : 50,
    }
	rdlink.AddLinkNode(list, cat1, 0)
	rdlink.AddLinkNode(list, cat2, 0)
	rdlink.AddLinkNode(list, cat30, 0)
	rdlink.AddLinkNode(list, cat40, 0)
	rdlink.AddLinkNode(list, cat50, 0)
	return list
}

func main() {
	fmt.Println("hi, linklist level1")

	fmt.Println("================================>下面测试单向链表")
	TestSinglelink()
	list1 := createSList()
	slink.ShowLink(list1)

	fmt.Println("================================>下面测试双向链表")
	TestDoublelink()
	list2 := createDList()
	dlink.ShowLink(list2)

	// fmt.Println("================================>下面测试环形单向链表")
	// TestRingSinglelink()
	// list3 := createRSList()
	// rslink.ShowLink(list3)

	// fmt.Println("================================>下面测试环形双向链表")
	// TestRingDoublelink()
	// list4 := createRDList()
	// rdlink.ShowLink(list4)
	


	
}