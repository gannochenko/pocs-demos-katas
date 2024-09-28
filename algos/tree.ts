type TreeNode<P> = {
	id: number;
	payload: P;
};

type TreeNodeInput<P> = {
	id?: number;
	payload: P;
};

class Tree<P> {
	root: TreeNode<P> | null = null;
	nodes: Map<number, TreeNode<P>>;
	edges: Map<number, number[]>;

	constructor() {
		this.nodes = new Map<number, TreeNode<P>>();
		this.edges = new Map<number, number[]>();
	}

	public appendChild(parentId: number | null, nodeInput: TreeNodeInput<P>) {
		if (!nodeInput.id) {
			nodeInput.id = Date.now();
		}

		if (this.nodes.has(nodeInput.id)) {
			// the node must be re-attached to some other parent
			throw new Error("not implemented");
		} else {
			const node = {
				id: nodeInput.id,
				payload: nodeInput.payload,
			};

			this.edges.set(node.id, []);
			this.nodes.set(node.id, node);

			if (parentId === null) {
				this.root = node;
			} else {
				this.edges.get(parentId)?.push(node.id);
			}
		}
	}

	public traverseInBreadth(cb: (node: TreeNode<P>, depth: number) => void) {
		if (!this.root) {
			return;
		}

		let stack = [{
			nodeID: this.root.id,
			depth: 0,
		}];

		while(stack.length) {
			let stackItem = stack.shift();
			if (!stackItem) {
				return;
			}

			let node = this.nodes.get(stackItem.nodeID);
			if (!node) {
				return;
			}

			cb(node, stackItem.depth);

			this.edges.get(stackItem.nodeID)?.forEach(connectedNodeId => {
				stack.push({
					nodeID: connectedNodeId,
					depth: stackItem.depth + 1,
				});
			});
		}
	}

	public traverseInDepth(cb: (node: TreeNode<P>, depth: number) => void) {
		if (!this.root) {
			return;
		}

		const traverse = (node: TreeNode<P>, depth: number) => {
			if (!node) {
				return;
			}

			cb(node, depth);

			this.edges.get(node.id)?.forEach(connectedNodeId => {
				const childNode = this.nodes.get(connectedNodeId);
				if (childNode) {
					traverse(childNode, depth+1);
				}
			});
		};

		traverse(this.root, 0);
	}
}

// ////////////////
// ////////////////
// ////////////////

type NodeID = number;
type MaybeMissingNodeID = NodeID | null;

class BinaryTree<P> {
	root: TreeNode<P> | null = null;
	nodes: Map<NodeID, TreeNode<P>>;
	edges: Map<NodeID, [MaybeMissingNodeID, MaybeMissingNodeID]>;

	constructor() {
		this.nodes = new Map<NodeID, TreeNode<P>>();
		this.edges = new Map<NodeID, [MaybeMissingNodeID, MaybeMissingNodeID]>();
	}

	appendChildLeft(parentId: NodeID | null, nodeInput: TreeNodeInput<P>) {
		this.appendChild(parentId, nodeInput, true);
	}

	appendChildRight(parentId: NodeID | null, nodeInput: TreeNodeInput<P>) {
		this.appendChild(parentId, nodeInput, false);
	}

	private appendChild(parentId: NodeID | null, nodeInput: TreeNodeInput<P>, left: boolean) {
		if (!nodeInput.id) {
			nodeInput.id = Date.now();
		}

		if (this.nodes.has(nodeInput.id)) {
			// the node must be re-attached to some other parent
			throw new Error("not implemented");
		} else {
			const node = {
				id: nodeInput.id,
				payload: nodeInput.payload,
			};

			this.edges.set(node.id, [null, null]);
			this.nodes.set(node.id, node);

			if (parentId === null) {
				this.root = node;
			} else {
				const edges = this.edges.get(parentId);
				if (edges) {
					edges[left ? 0 : 1] = node.id;
				}
			}
		}
	}

	public traverseInBreadth(cb: (node: TreeNode<P>, depth: NodeID) => void) {
		if (!this.root) {
			return;
		}

		let stack = [{
			nodeID: this.root.id,
			depth: 0,
		}];

		while(stack.length) {
			let stackItem = stack.shift();
			if (!stackItem) {
				return;
			}

			let node = this.nodes.get(stackItem.nodeID);
			if (!node) {
				return;
			}

			cb(node, stackItem.depth);

			this.edges.get(stackItem.nodeID)?.forEach(connectedNodeId => {
				if (connectedNodeId !== null) {
					stack.push({
						nodeID: connectedNodeId,
						depth: stackItem.depth + 1,
					});
				}
			});
		}
	}
}

interface AnyTree<P> {
	traverseInBreadth(cb: (node: TreeNode<P>, depth: NodeID) => void): void;
}

function printTree<P>(tree: AnyTree<P>) {
	let currentDepth = 0;
	tree.traverseInBreadth((node, depth) => {
		if (currentDepth !== depth) {
			process.stdout.write("\n");
			currentDepth = depth;
		}
		process.stdout.write(node.payload+("("+depth+") \t".repeat(depth)));
	});
	process.stdout.write("\n");
}

function main() {
	const tree = new BinaryTree<number>();
	tree.appendChildLeft(null, {
		id: 1,
		payload: 1,
	})
	tree.appendChildLeft(1, {
		id: 2,
		payload: 2,
	})
	tree.appendChildRight(1, {
		id: 3,
		payload: 3,
	})
	console.log(tree);
	
	// tree.appendChildLeft(2, {
	// 	id: 4,
	// 	payload: 4,
	// })
	//
	// console.log(tree);

	printTree(tree);
}

main();