package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"syscall"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func HumanSize(size int64) string {
	const (
		_ = 1 << (10 * iota)
		KiB
		MiB
		GiB
		TiB
		PiB
		EiB
	)

	if size < 2<<10 {
		return fmt.Sprintf("%6d B", size)
	}
	if size < 2<<20 {
		return fmt.Sprintf("%6.1f K", float64(size)/KiB)
	}
	if size < 2<<30 {
		return fmt.Sprintf("%6.1f M", float64(size)/MiB)
	}
	if size < 2<<40 {
		return fmt.Sprintf("%6.1f G", float64(size)/GiB)
	}
	if size < 2<<50 {
		return fmt.Sprintf("%6.1f T", float64(size)/TiB)
	}
	if size < 2<<60 {
		return fmt.Sprintf("%6.1f P", float64(size)/PiB)
	}

	return fmt.Sprintf("%6.1f E", float64(size)/EiB)
}

type nodeMeta struct {
	Name string
	Size int64
}

func add(target *tview.TreeNode, path string) int64 {
	var size int64

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	var newNodes []*tview.TreeNode

	for _, file := range files {
		var thissize int64
		thissize = 0
		name := filepath.Join(path, file.Name())
		node := tview.NewTreeNode(file.Name()).
			SetSelectable(file.IsDir())
		node.SetExpanded(false)

		if file.IsDir() {
			thissize = add(node, name)
			node.SetColor(tcell.ColorGreen)
		} else {
			thissize += file.Sys().(*syscall.Stat_t).Blocks * 512
		}

		node.SetText(fmt.Sprintf("  %v %s", HumanSize(thissize), node.GetText()))

		nodeMeta := nodeMeta{
			Name: name,
			Size: thissize,
		}
		node.SetReference(nodeMeta)

		size += thissize
		newNodes = append(newNodes, node)
	}

	sort.Slice(newNodes, func(i, j int) bool {
		return newNodes[i].GetReference().(nodeMeta).Size > newNodes[j].GetReference().(nodeMeta).Size
	})

	for _, node := range newNodes {
		target.AddChild(node)
	}

	return size
}

func main() {
	rootDir := "."

	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorGreen)

	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

	size := add(root, rootDir)
	root.SetText(fmt.Sprintf("     %v .", HumanSize(size)))

	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		node.SetExpanded(!node.IsExpanded())
	})

	if err := tview.NewApplication().SetRoot(tree, true).Run(); err != nil {
		panic(err)
	}
}
