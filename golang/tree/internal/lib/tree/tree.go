package tree

import (
	"encoding/json"
	"strings"

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
		pathParts := make([]string, 0)

		for index, pathElement := range path {
			pathParts = append(pathParts, pathElement)
			parentPath := path[:len(pathParts)-1]
			pathKey := strings.Join(pathParts, ".")
			parentPathKey := strings.Join(parentPath, ".")

			if _, ok := t.Refs[pathKey]; !ok {
				items := make([]*tree.Node[P], 0)
				newNode := tree.Node[P]{
					Name:  pathElement,
					Nodes: &items,
				}

				if index == len(path)-1 {
					// adding payload
					newNode.Payload = payload
				}

				toNode := t.Refs[parentPathKey]
				toNodeNodesDeref := *toNode.Nodes
				toNodeNodesDeref = append(toNodeNodesDeref, &newNode)
				toNode.Nodes = &toNodeNodesDeref

				t.Refs[pathKey] = &newNode
			}
		}
	}
}

func (t *Tree[P]) ToJSON() string {
	jsonData, _ := json.MarshalIndent(t.Root, "", "  ")
	return string(jsonData)
}
