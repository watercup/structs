/***************************************************************************
 *
 * @desc
 * @author <wanyang@wanyang.site>
 * @version 05/12/2017 11:03
 * @project structs
 **************************************************************************/
package rb_tree

import "fmt"

func main() {
	root := &RbNode{}
	insert(&root, 40, "f")
	insert(&root, 30, "t")
	insert(&root, 20, "tw")
	insert(&root, 10, "ten")
	insert(&root, 5, "five")
	insert(&root, 50, "fif")
	insert(&root, 60, "sf")
	insert(&root, 70, "sf")
	insert(&root, 15, "sf")
	insert(&root, 100, "sf")
	insert(&root, 65, "sf")
	preOrder(root)
}

const (
	Red = iota
	Black
)

type Color uint8

type RbNode struct {
	color  Color
	left   *RbNode
	right  *RbNode
	parent *RbNode
	key    int
	value  string
}

func preOrder(root *RbNode) {
	if root != nil {
		preOrder(root.left)
		fmt.Print(root.key, " ")
		preOrder(root.right)
	}
}

func insert(root **RbNode, key int, value string) {
	node := &RbNode{
		color: Red,
		key:   key,
		value: value,
	}
	var cursor *RbNode
	x := *root
	for x != nil && x.key != 0 {
		cursor = x
		if key > x.key {
			x = x.right
		} else {
			x = x.left
		}
	}
	node.parent = cursor
	if cursor != nil {
		if node.key > cursor.key {
			cursor.right = node
		} else {
			cursor.left = node
		}
	} else {
		*root = node
	}
	insertFixUp(root, node)
}

func insertFixUp(root **RbNode, fixNode *RbNode) {
	for fixNode.parent != nil && fixNode.parent.parent != nil && fixNode.parent.color == Red {
		//父节点为祖父的左子
		if fixNode.parent.parent.left == fixNode.parent {
			//如果叔叔也是红的
			if fixNode.parent.parent.right != nil && fixNode.parent.parent.right.color == Red {
				//则将父|叔为黑，祖父为红，从祖父节点继续向上
				fixNode.parent.parent.right.color = Black
				fixNode.parent.color = Black
				fixNode.parent.parent.color = Red
				fixNode = fixNode.parent.parent
				continue
			}
			if fixNode == fixNode.parent.right {
				//如果当前节点是右孩子，需要左旋变成 将子节点上升
				leftRotate(root, fixNode.parent)
				//继续从最低影响节点开始继续
				fixNode = fixNode.left
			}
			fixNode.parent.color = Black
			fixNode.parent.parent.color = Red
			rightRotate(root, fixNode.parent.parent)
		} else {
			//如果叔叔也是红的
			if fixNode.parent.parent.left != nil && fixNode.parent.parent.left.color == Red {
				//则将父|叔为黑，祖父为红，从祖父节点继续向上
				fixNode.parent.parent.left.color = Black
				fixNode.parent.color = Black
				fixNode.parent.parent.color = Red
				fixNode = fixNode.parent.parent
				continue
			}
			if fixNode == fixNode.parent.left {
				rightRotate(root, fixNode.parent)
				fixNode = fixNode.right
			}
			fixNode.parent.color = Black
			fixNode.parent.parent.color = Red
			leftRotate(root, fixNode.parent.parent)
		}
	}
	(*root).color = Black
}

func leftRotate(root **RbNode, x *RbNode) {
	y := x.right
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}
	y.parent = x.parent
	if x.parent == nil {
		*root = x
	} else {
		if x.parent.left == x {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	}
	y.left = x
	x.parent = y
}

func rightRotate(root **RbNode, y *RbNode) {
	x := y.left
	//将x的有孩子 送给y
	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}
	//x替换y的所有信号
	x.parent = y.parent
	if y.parent == nil {
		*root = x
	} else {
		if y == y.parent.left {
			y.parent.left = x
		} else {
			y.parent.right = x
		}
	}
	//y称为x的子节点
	y.parent = x
	x.right = y
}

func remove(root *RbNode, z *RbNode) {

}

func removeFixUp(root *RbNode, z *RbNode) {

}
