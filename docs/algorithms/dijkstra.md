# Dijkstra's algorithm

Dijkstra's algorithm is an algorithm for finding the shortest paths between nodes in a weighted graph.
For a given source node in the graph, the algorithm finds the shortest path between that node and every other.
Here’s a good [video explanation](https://www.youtube.com/watch?v=pSqmAO-m7Lk&ab_channel=WilliamFiset).

The main idea is that if we follow the shortest path in each iteration, we'll find a shortest path eventually.

Alternatives:
- Simple DFS is unable to solve this task, because it can finds a path, not the shortest path.
But you can dance around it, keeping track of your depths and updating it in some kind of global array, I guess.
By the way, if you have grid-like structure graph, DFS would work as good as dijkstra.
- Simple BFS finds a path, not the shortest path.
But you can dance around it, keeping track of your depths and updating shortest paths in some kind of global array, I guess.
Dijkstra just works faster because BFS visits every node, and Dijkstra visits less nodes.

Eager dijkstra replaces irrelevant paths from our queue, taking up less space and working faster. I've implemented it too!

Source code:
- Dijkstra: [src/algorithms/dijkstra.go](../../src/algorithms/dijkstra.go)  
- Eager Dijkstra: [src/algorithms/eager_dijkstra.go](../../src/algorithms/eager_dijkstra.go)
