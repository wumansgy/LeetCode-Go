package main

func main(){

}

// 递归的方法
func MergeList(L1 *Node,L2 *Node )*Node{  //有序合并两个单链表,这里内容存储int类型才能比较

	if L1==nil&&L2==nil{
		return nil   //两个都为空就返回nil
	}
	if L1==nil{
		return L2   //l1为空就返回l2
	}
	if L2==nil{
		return L1   //l2为空就返回l1
	}
	NewList:=new(Node)  //新的链表
	if L1.Data.(int)>L2.Data.(int){
		NewList=L2
		NewList.Next=MergeList(L1,L2.Next)   //这里递归操作
	}else{
		NewList=L1
		NewList.Next=MergeList(L1.Next,L2)   //递归
	}
	return NewList
}

// 循环方法来实现单链表合并
func MergeList02(L1 *Node,L2 *Node)*Node{   //普通循环遍历

	NewList:=new(Node)
	NewList.Next=nil  //创建新链表

	p1:=new(Node)  //声明两个指针分别指向两个链表的数据开始部分
	p2:=new(Node)
	p1=L1
	p2=L2

	last:=new(Node)   //声明一个指针指向新链表的最后一个结点，
	last=NewList

	for p1!=nil&&p2!=nil{

		if p1.Data.(int)<p2.Data.(int){ //p1结点的数据小，last指向p1结点，然后向后移
			last.Next=p1
			p1=p1.Next
			last=last.Next
		}else{  //p2结点的数据小，last指向p2结点，然后向后移
			last.Next=p2
			p2=p2.Next
			last=last.Next
		}
	}
	//当有一个链表结束时，将last指向另外一条链表
	if p1==nil{
		last.Next=p2
	} else if p2==nil{
		last.Next=p1
	}
	return NewList
}

