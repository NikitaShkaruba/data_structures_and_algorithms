# Dijkstra's algorithm

Dijkstra's algorithm is an algorithm for finding the shortest paths between nodes in a weighted graph.
For a given source node in the graph, the algorithm finds the shortest path between that node and every other.
Here’s a good [video explanation](https://www.youtube.com/watch?v=pSqmAO-m7Lk&ab_channel=WilliamFiset).

Simple DFS will not work, because it can only determine the path with the unweighted edges.
By the way, if you have grid-like structure graph, DFS would work as good as dijkstra.

I’m implementing eager dijkstra, which is the fastest.

Source code: [src/algorithms/dijkstra.go](../../src/algorithms/dijkstra.go)
