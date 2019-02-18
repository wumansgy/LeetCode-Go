package main

type Node struct{
	Data interface{}
	Next *Node
}

func main(){

}

// 利用Go的多重赋值
func reverseList01(head *Node)*Node{
	cur:=head
	var pre *Node=nil
	for cur!=nil{
		pre,cur,cur.Next=cur,cur.Next,pre
	}
	return pre
}

// 迭代循环来反转
func reverseList02(head *Node)*Node{
	nextNode:=new(Node)
	preNode:=new(Node)
	nextNode=nil
	preNode=nil

	for head!=nil{
		nextNode=head.Next   //先保存链表的下一个=

		head.Next=preNode   //反转

		preNode=head   //相当于往下跑一个

		head=nextNode   //head也要往后面跑

	}

	return preNode
}

// 递归来反转
var endnode *Node
var newlist *Node
func recursiveTraverse(head *Node){
	if head==nil{
		return    //出口
	}
	if head.Next==nil{
		endnode=head
		newlist=head
		return     //出口
	}

	recursiveTraverse(head.Next)

	endnode.Next=head   //反转
	endnode=head   //相当于下移
	endnode.Next=nil  //最后需要尾指针
}