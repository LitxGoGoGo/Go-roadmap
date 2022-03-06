package main

import "fmt"

//定义结构体节点
type HeroNode struct {
	No    int
	Name  string
	Left  *HeroNode
	Right *HeroNode
}

//前序遍历,先输出root然后再输出左子树,再输出又子树
func PreOrder(node *HeroNode) {
	if node != nil {
		fmt.Printf("no=%d,name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}

}

//中序遍历,先输出root的左子树,再输出root,最后输出右子树
func InfixOrder(node *HeroNode) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("no=%d,name=%s\n", node.No, node.Name)
		InfixOrder(node.Right)
	}

}

//后序遍历,先输出root的左子树,最后输出右子树,再输出root
func PostOrder(node *HeroNode) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("no=%d,name=%s\n", node.No, node.Name)
	}
}
func main() {
	//构建一个二叉树
	root := &HeroNode{
		No:   1,
		Name: "宋江",
	}

	left1 := &HeroNode{
		No:   2,
		Name: "吴用",
	}

	node10 := &HeroNode{
		No:   10,
		Name: "10",
	}
	node12 := &HeroNode{
		No:   12,
		Name: "12",
	}

	left1.Left = node10
	left1.Right = node12

	right1 := &HeroNode{
		No:   3,
		Name: "卢俊义",
	}

	root.Left = left1
	root.Right = right1

	right2 := &HeroNode{
		No:   4,
		Name: "林冲",
	}

	right1.Right = right2

	//PreOrder(root)

	//InfixOrder(root)

	PostOrder(root)
}
