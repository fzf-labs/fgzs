package tree

type Tree struct {
	Id       int64       `json:"id"`
	ParentId int64       `json:"parentId"`
	Title    string      `json:"title"`
	Other    interface{} `json:"other"`
	Children []Tree      `json:"children"`
}

type Trees []Tree

func (t Trees) Len() int {
	return len(t)
}

func (t Trees) Less(i, j int) bool {
	return t[i].Id < t[j].Id
}

func (t Trees) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func GenerateTree(list []Tree) []Tree {
	var trees []Tree
	// Define the top-level root and child nodes
	var roots, childs []Tree
	for _, v := range list {
		if v.ParentId == 0 {
			// Determine the top-level root node
			roots = append(roots, v)
		}
		childs = append(childs, v)
	}

	for _, v := range roots {
		childTree := &Tree{
			Id:       v.Id,
			ParentId: v.ParentId,
			Title:    v.Title,
			Other:    v.Other,
			Children: make([]Tree, 0),
		}
		// recursive
		recursiveTree(childTree, childs)

		trees = append(trees, *childTree)
	}
	return trees
}

func recursiveTree(tree *Tree, allNodes []Tree) {
	for _, v := range allNodes {
		if v.ParentId == 0 {
			// If the current node is the top-level root node, skip
			continue
		}
		if tree.Id == v.ParentId {
			childTree := &Tree{
				Id:       v.Id,
				ParentId: v.ParentId,
				Title:    v.Title,
				Other:    v.Other,
				Children: make([]Tree, 0),
			}
			recursiveTree(childTree, allNodes)
			tree.Children = append(tree.Children, *childTree)
		}
	}
}

// FindSubNode 查询子级
func FindSubNode(node *Tree, allNodes []Tree) {
	for _, v := range allNodes {
		if node.Id == v.ParentId {
			FindSubNode(&v, allNodes)
			node.Children = append(node.Children, v)
		}
	}
}

// FindParentNode 查询父级
func FindParentNode(node *Tree, allNodes []Tree) {
	for _, v := range allNodes {
		temp := v
		if node.ParentId == temp.Id {
			temp.Children = append(temp.Children, *node)
			*node = temp
			FindParentNode(node, allNodes)
		}
	}
}
