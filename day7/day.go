package day7

import (
	"AdventOfCode2022/util"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const TOTAL = 70000000
const NECESSARY = 30000000

type Tree struct {
	superNode   *Tree
	name        string
	size        int
	isDirectory bool
	nodes       []*Tree
}

func isLS(line string) bool {
	return strings.HasPrefix(line, "$ ls")
}

func isCD(line string) (*string, bool) {
	res := strings.Split(line, "$ cd ")
	if len(res) > 1 {
		return &res[1], true
	}
	return nil, false
}

func isDir(line string) (*string, bool) {
	res := strings.Split(line, "dir ")
	if len(res) > 1 {
		return &res[1], true
	}
	return nil, false
}

func isFile(line string) (*string, int, bool) {
	split := strings.Split(line, " ")
	if len(split) != 2 {
		return nil, 0, false
	}
	size, err := strconv.Atoi(split[0])
	if err != nil {
		return nil, 0, false
	}
	return &split[1], size, true
}

func printTree(tree *Tree, layer int) {
	toPrint := ""
	for i := 0; i < layer*3; i++ {
		toPrint += " "
	}
	toPrint += "- " + tree.name
	if tree.isDirectory {
		toPrint += " (dir)"
	} else {
		toPrint += " (file, size=" + strconv.Itoa(tree.size) + ")"
	}
	fmt.Println(toPrint)
	for _, node := range tree.nodes {
		printTree(node, layer+1)
	}
}

func part1(tree *Tree) int {
	values := 0
	for _, node := range tree.nodes {
		values += part1(node)
	}
	if tree.size > 0 && tree.size <= 100000 && tree.isDirectory {
		values += tree.size
	}
	return values
}

func part2Fn(tree *Tree, minimumSize int) int {
	minSize := -1
	for _, node := range tree.nodes {
		if !node.isDirectory {
			continue
		}
		if node.size > minimumSize {
			temp := part2Fn(node, minimumSize)
			if minSize == -1 || minSize > temp {
				minSize = temp
			}
		}
	}

	if tree.size > minimumSize {
		if minSize == -1 || minSize > tree.size {
			minSize = tree.size
		}
	}
	return minSize
}

func addSizesToSuperNode(tree *Tree, size int) {
	tree.size += size
	if tree.superNode != nil {
		addSizesToSuperNode(tree.superNode, size)
	}
}

func getValue(file string, part2 bool) {
	results, err := util.GetFileContentsSplit(file)
	util.PanicOnError(err)

	filesystem := Tree{
		name:        "/",
		size:        0,
		isDirectory: true,
		nodes:       make([]*Tree, 0),
	}
	currentTree := &filesystem
	for _, line := range results {
		if dir, proceed := isCD(line); proceed {
			if *dir == "/" {
				continue
			}
			if *dir == ".." {
				currentTree = currentTree.superNode
				continue
			}
			var targetDir *Tree
			for _, node := range currentTree.nodes {
				if node.name == *dir {
					targetDir = node
				}
			}

			if targetDir == nil {
				panic("directory " + *dir + " not found")
			}
			currentTree = targetDir
			continue
		}
		if isLS(line) {
			continue
		}
		if dir, proceed := isDir(line); proceed {
			currentTree.nodes = append(currentTree.nodes, &Tree{
				superNode:   currentTree,
				name:        *dir,
				size:        0,
				isDirectory: true,
				nodes:       make([]*Tree, 0),
			})
		}
		if file, size, proceed := isFile(line); proceed {
			addSizesToSuperNode(currentTree, size)
			currentTree.nodes = append(currentTree.nodes, &Tree{
				superNode:   currentTree,
				name:        *file,
				size:        size,
				isDirectory: false,
				nodes:       make([]*Tree, 0),
			})
		}

	}

	printTree(&filesystem, 0)
	if part2 {
		left := TOTAL - filesystem.size
		needed := NECESSARY - left
		res := part2Fn(&filesystem, needed)
		log.Println("left over: ", left)
		log.Println("needed: ", needed)
		log.Println("res: ", res)
	} else {
		log.Println(part1(&filesystem))
	}

}
