package main

func main(){

}

// 环检测
func IsLoop(head *Node)bool{
	fast:=new(Node)  //快指针
	slow:=new(Node)  //慢指针
	fast=head
	slow=head   //指向头结点

	for slow!=nil&&fast.Next!=nil{
		slow=slow.Next       //慢指针走一步
		fast=fast.Next.Next  //快指针走两步

		if fast==slow{   //有环 ，如果有环迟早会相遇
			return true
		}
	}

	return false
}

// 环入口检测
func listLoopEn(head *Node)*Node{   //寻找链表的环的入点
	fast:=new(Node)  //快指针
	slow:=new(Node)  //慢指针
	fast=head
	slow=head   //指向头结点

	for slow!=nil&&fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next  //快指针走两步

		if fast==slow{   //有环
			break    //记录相遇的点
		}
	}
	if slow==nil||fast.Next==nil{
		return nil    //说明不存在环，直接返回
	}
	ptr1:=new(Node) //指向第一个结点
	ptr2:=new(Node) //指向之前相遇的结点
	ptr1=head
	ptr2=slow

	for ptr1!=ptr2{
		ptr1=ptr1.Next
		ptr2=ptr2.Next   //一直向下走一步一步
	}
	return ptr1   //返回这个相遇的点，就是环的入口的点

}

// 环长度检测
func ListLoopLength(head *Node)int{   //环长度检测
	//求出环的长度
	//思路1：记录下相遇节点存入临时变量tempPtr，然后让slow(或者fast，都一样)
	// 继续向前走slow = slow -> next；一直到slow == tempPtr; 此时经过的步数就是环上节点的个数；
	//思路2： 从相遇点开始slow和fast继续按照原来的方式向前走slow = slow -> next; fast = fast -> next -> next；
	// 直到二者再次项目，此时经过的步数就是环上节点的个数 。
	fast:=new(Node)  //快指针
	slow:=new(Node)  //慢指针
	fast=head
	slow=head   //指向头结点
	tmp:=new(Node)   //保存临时点
	var length int
	for slow!=nil&&fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next  //快指针走两步

		if fast==slow{   //有环
			tmp=slow
			break
		}
	}
	length++
	tmp=tmp.Next  //先向前走一步
	for tmp!=slow{
		tmp=tmp.Next
		length++
	}
	return length
}

// 拓展
//求链表长度，
//可以根据前面求出的条件求出起点到入口的距离然后加上环的长度，就是链表的长度

//求出环上距离任意一个节点最远的点（对面节点）
//对于换上任意的一个点ptr0, 我们要找到它的”对面点“，可以这样思考：同样使用上面的快慢指针的方法，让slow和fast都指向ptr0，每一步都执行与上面相同的操作（slow每次跳一步，fast每次跳两步），

//当fast = ptr0或者fast = prt0->next的时候slow所指向的节点就是ptr0的”对面节点“。

//对于问题6（扩展）如何判断两个无环链表是否相交，和7（扩展）如果相交，求出第一个相交的节点，其实就是做一个问题的转化：
//设有连个链表listA和listB，如果两个链表都无环，并且有交点，那么我们可以让其中一个链表（不妨设是listA）的为节点连接到其头部，这样在listB中就一定会出现一个环。

//因此我们将问题6和7分别转化成了问题1和2