# AVL Tree

AVL tree is a type of self-balanced search tree.
Tree is said to be balanced differences between the heights of left and right subtrees for every node are less than or equal to 1.

Basically simple search tree can become very unbalanced if you always insert bigger and bigger values, for example.
The tree will then look like a string `1(root) - 2 - 3 - 4`, which leads to `O(n)` time for searching inside of it.
Self-balanced search tree self-balances almost every time you insert into it for no cost whatsoever in terms of time complexity.

Inserting, pushing and searching in avl tree costs `O(logn)` time.

There are other types of self-balanced search trees (Red-Black tree for example), but AVL tree is pretty simple, so I've decided to use it in my library.

Source code: [src/data_structures/avl_tree.go](../../src/data_structures/avl_tree.go)

### Heap comparison

Checkout [docs/data_structures/heap.md#avl-tree-comparison](./heap.md#avl-tree-comparison)
