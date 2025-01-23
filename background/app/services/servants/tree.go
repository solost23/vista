package servants

import (
	"context"
)

type TreeNode interface {
	ID() uint
	ParentID() uint
	AppendChildren(any)
}

const rootParentId = 0

// BuildTree 重建树 o(n^2)
// @param: array []TreeNode
// @return: treeNodes []TreeNode
func BuildTree(_ context.Context, array []TreeNode) []TreeNode {
	maxLen := len(array)
	var rootNodes = make([]TreeNode, 0, len(array)/2+1)
	// 找出根节点,根节点的特点，没有父节点
	for i := 0; i < maxLen; i++ {
		// 统计每个节点的父节点出现的次数，父节点出现0次就是根节点
		count := 0
		for j := 0; j < maxLen; j++ {
			// 如果有节点的ID == i的parentID 那么j就是父节点
			if array[j].ID() == array[i].ParentID() {
				count++
				array[j].AppendChildren(array[i])
			}
		}
		if count == 0 {
			rootNodes = append(rootNodes, array[i])
		}
	}
	return rootNodes
}

// BuildTree2 重建树 o(n)
// @param: array []TreeNode
// @return: treeNodes []TreeNode
func BuildTree2(_ context.Context, array []TreeNode) []TreeNode {
	if len(array) <= 0 {
		return array
	}

	maxLen := len(array)

	treeIdMap := make(map[uint]TreeNode, maxLen)
	for i := 0; i != maxLen; i++ {
		treeIdMap[array[i].ID()] = array[i]
	}

	trees := make([]TreeNode, 0, 1)
	for i := 0; i != maxLen; i++ {
		if array[i].ParentID() == rootParentId {
			trees = append(trees, array[i])
		} else {
			treeIdMap[array[i].ParentID()].AppendChildren(array[i])
		}
	}
	return trees
}
