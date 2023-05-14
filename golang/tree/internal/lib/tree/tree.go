package tree

import (
	"tree/internal/domain/tree"
)

type Tree[P tree.PayloadType] struct {
	Root *tree.Node[P]
	Refs map[string]*tree.Node[P]
}

func New[P tree.PayloadType]() *Tree[P] {
	items := make([]*tree.Node[P], 0)
	rootNode := &tree.Node[P]{
		Nodes: &items,
	}
	refs := make(map[string]*tree.Node[P])
	refs[""] = rootNode

	return &Tree[P]{
		Root: rootNode,
		Refs: refs,
	}
}

func (t *Tree[P]) AddNode(path []string, payload P) {
	if len(path) > 0 {
		//pathParts := make([]string, 0)
		//
		//for index, pathElement := range path {
		//	pathParts = append(pathParts, pathElement)
		//	parentPath := path[:len(pathParts)-1]
		//	pathKey := strings.Join(pathParts, ".")
		//	parentPathKey := strings.Join(parentPath, ".")
		//
		//	// if element was not created before
		//	if _, ok := refs[pathKey]; !ok {
		//		items := make([]*collection.Node, 0)
		//		newNode := collection.Node{
		//			Name: pathElement,
		//			Item: &items,
		//		}
		//
		//		if index == len(parseResult.Package.Path)-1 {
		//			// adding a service here
		//			for _, service := range parseResult.Package.Services {
		//				itemsDeref := *newNode.Item
		//				itemsDeref = append(itemsDeref, g.createServiceNode(service))
		//				newNode.Item = &itemsDeref
		//			}
		//		}
		//
		//		toNode := refs[parentPathKey]
		//		toNodeItemsDeref := *toNode.Item
		//		toNodeItemsDeref = append(toNodeItemsDeref, &newNode)
		//		toNode.Item = &toNodeItemsDeref
		//
		//		refs[pathKey] = &newNode
		//	}
		//}
	}
}

func (t *Tree[P]) ToJSON() string {
	return ""
}
