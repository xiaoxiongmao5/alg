package main

import (
	"fmt"
	slink "xj/chapter1_linklist/homework/single_linklist"
	"os"
)

type Student struct {
	Name string
	Language string
}
// 返回查找目标位置的前一个节点
func FindNode(headNode *slink.LinkNode[Student], newNode *slink.LinkNode[Student]) (bool, *slink.LinkNode[Student]){
	if headNode.Next == nil {
		fmt.Println("该学员链表中暂无学员")
		return false, headNode
	}
	// 第一个节点
	tmp := headNode.Next
	find := false
	lan := newNode.Val.Language
	for {
		if tmp.Next == nil {
			find = true
			break
		}
		// 找到该链表该语言的最后位置了
		if tmp.Val.Language == lan && tmp.Next.Val.Language != lan {
			find = true
			break
		}
		tmp = tmp.Next
	}
	if find {
		return find, tmp
	}
	return false, tmp
}
// 显式学员列表信息
func ShowLinkInfo(linklist *slink.LinkList[Student]) {
	fmt.Println("————当前学员列表信息如下：————")
	if linklist.Head == nil || linklist.Head.Next == nil {
		fmt.Println("暂无学员")
		return
	}
	tmp := linklist.Head.Next
	for {
		fmt.Printf("{\"%s\",\"%s\"}, ", tmp.Val.Name, tmp.Val.Language)
		if tmp.Next == nil {
			break
		}
		if tmp.Val.Language != tmp.Next.Val.Language {
			fmt.Println("")
		}
		tmp = tmp.Next
	}
	fmt.Println("")
}

// 添加初始已定的6位学员
func AddInitStu6List(linklist *slink.LinkList[Student]){
	stu1Node := &slink.LinkNode[Student]{
		Val:Student{
			Name:"aa",
			Language:"Java",
		},
	}
	stu2Node := &slink.LinkNode[Student]{
		Val:Student{
			Name:"ab",
			Language:"Java",
		},
	}
	stu3Node := &slink.LinkNode[Student]{
		Val:Student{
			Name:"ba",
			Language:"CPP",
		},
	}
	stu4Node := &slink.LinkNode[Student]{
		Val:Student{
			Name:"bb",
			Language:"CPP",
		},
	}
	stu5Node := &slink.LinkNode[Student]{
		Val:Student{
			Name:"ca",
			Language:"Python",
		},
	}
	stu6Node := &slink.LinkNode[Student]{
		Val:Student{
			Name:"cb",
			Language:"Python",
		},
	}	
	// 初始化
	linklist.AddLastNode(stu1Node)
	linklist.AddLastNode(stu2Node)
	linklist.AddLastNode(stu3Node)
	linklist.AddLastNode(stu4Node)
	linklist.AddLastNode(stu5Node)
	linklist.AddLastNode(stu6Node)
}

// 添加已定的3位学员
func AddStu3List(linklist *slink.LinkList[Student]){
	stu7Node := &slink.LinkNode[Student]{
		Val:Student{
			Name:"cc",
			Language:"Python",
		},
	}
	stu8Node := &slink.LinkNode[Student]{
		Val:Student{
			Name:"bc",
			Language:"CPP",
		},
	}
	stu9Node := &slink.LinkNode[Student]{
		Val:Student{
			Name:"ac",
			Language:"Java",
		},
	}
	// stu10Node := &slink.LinkNode[Student]{
	// 	Val:Student{
	// 		Name:"ff",
	// 		Language:"Go",
	// 	},
	// }
	_, node := FindNode(linklist.Head, stu7Node)
	linklist.AddAfter(node, stu7Node)

	_, node = FindNode(linklist.Head, stu8Node)
	linklist.AddAfter(node, stu8Node)

	_, node = FindNode(linklist.Head, stu9Node)
	linklist.AddAfter(node, stu9Node)

	// _, node = FindNode(linklist.Head, stu10Node)
	// linklist.AddAfter(node, stu10Node)
}

func main() {
	// 头结点
	headNode := &slink.LinkNode[Student]{}
	linklist := &slink.LinkList[Student]{
		Head:headNode,
	}

	var num int
	var InitStuListFlag bool
	var stuName string
	var stuLan string
	for {
		fmt.Println("===============学员管理系统===============")
		fmt.Println("===============输入 1 添加初始已定的6位学员")
		fmt.Println("===============输入 2 添加后面已定的3位学员")
		fmt.Println("===============输入 3 显式学员列表信息")
		fmt.Println("===============输入 4 添加新的学员")
		fmt.Println("===============输入 5 退出系统")
		fmt.Println("请出入操作数字:")
		fmt.Scanln(&num)
		switch num {
			case 1:
				if (!InitStuListFlag){
					InitStuListFlag = true
					AddInitStu6List(linklist)
					fmt.Println("添加成功！")
				} else {
					fmt.Println("初始学员已经添加到系统中，无需二次添加")
				}
				ShowLinkInfo(linklist)
			case 2:
				AddStu3List(linklist)
				fmt.Println("添加成功！")
				ShowLinkInfo(linklist)
			case 3:
				ShowLinkInfo(linklist)
			case 4:
				fmt.Println("请出入新生姓名: ")
				fmt.Scanln(&stuName)
				fmt.Println("请出入新生专业(比如:Java、CPP、Python): ")
				fmt.Scanln(&stuLan)
				stuNode := &slink.LinkNode[Student]{
					Val:Student{
						Name:stuName,
						Language:stuLan,
					},
				}
				_, node := FindNode(linklist.Head, stuNode)
				linklist.AddAfter(node, stuNode)
				fmt.Println("添加成功！")
				ShowLinkInfo(linklist)
			case 5:
				os.Exit(0)
			default:
				fmt.Println("您的输入有误，请重新输入正确的操作数字")
		}
	}
}