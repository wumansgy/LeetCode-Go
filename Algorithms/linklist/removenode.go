package main

func main(){

}

//
func removeNthFromEnd(head *Node, n int) *Node {

	if n<=0||head==nil{
		return head
	}
	//定义快慢指针，利用快慢指针来实现删除倒数第几个结点
	fast:=new(Node)   //快指针
	slow:=new(Node)   //慢指针
	//  list:=new(ListNode)

	fast=head
	slow=head

	for i:=0;i<n;i++{
		fast=fast.Next   //快指针先走完需要走的步数
	}
	if fast==nil{
		return head.Next
	}

	for fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next   //一起走，等到快指针到结尾
	}

	slow.Next=slow.Next.Next
	return head
}

//求链表的中间结点 用快慢指针来实现，快指针走2步，慢指针走一步，快指针到底，慢指针就是中间结点
