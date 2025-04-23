package main

import (
	"fmt"
	"github.com/mzohreva/GoGraphviz/graphviz"
	"strings"
	"syscall/js"
)

var (
	window = js.Global().Get("window")
)

func main() {
	drawGraph()

	<-make(chan bool)
}

func drawGraph() {
	var sb strings.Builder
	arr := []int{1, 3, 4, 5, 9, 11, 15, 23, 27, 49}
	tree := newBinaryTree(arr, 0, len(arr)-1)
	G := visualizeBinaryTree(tree)
	G.GenerateDOT(&sb)
	dot := sb.String()

	window.Call("drawGraph", dot)
}

func visualizeBinaryTree(root *binaryTree) *graphviz.Graph {
	G := &graphviz.Graph{}
	addSubTree(root, G)
	G.DefaultNodeAttribute(graphviz.Shape, graphviz.ShapeCircle)
	G.GraphAttribute(graphviz.NodeSep, "0.3")
	G.MakeDirected()
	return G
}

func addSubTree(root *binaryTree, G *graphviz.Graph) int {
	if root == nil {
		null := G.AddNode("")
		G.NodeAttribute(null, graphviz.Shape, graphviz.ShapePoint)
		return null
	}
	rootNode := G.AddNode(fmt.Sprint(root.value))
	leftNode := addSubTree(root.left, G)
	rightNode := addSubTree(root.right, G)
	G.AddEdge(rootNode, leftNode, "")
	G.AddEdge(rootNode, rightNode, "")
	return rootNode
}

type binaryTree struct {
	value       int
	left, right *binaryTree
}

func newBinaryTree(arr []int, start, end int) *binaryTree {
	if start == end {
		return &binaryTree{arr[start], nil, nil}
	}
	if start > end {
		return nil
	}
	mid := (start + end) / 2
	root := &binaryTree{arr[mid], nil, nil}
	root.left = newBinaryTree(arr, start, mid-1)
	root.right = newBinaryTree(arr, mid+1, end)
	return root
}
