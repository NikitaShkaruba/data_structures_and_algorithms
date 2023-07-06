# Union find

Is used for finding connected components and minimum spanning tree of for undirected graphs.
Here’s a good [video explanation](https://www.youtube.com/watch?v=0jNmHPfA_yE). Basically you can `union` connected components,
and `find` which connected component this node belongs to.

Because union-find compresses itself, it's main operations `Union` and `Find` work in `O(1)` time, `O(1)` space, which
is simply outstanding.

Source code: [src/data_structures/union_find.go](../../src/data_structures/union_find.go)
