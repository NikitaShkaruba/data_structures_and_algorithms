# Greedy

A greedy algorithm is an approach for solving a problem by selecting the best option available at the moment.
The algorithm never reverses the earlier decision even if the choice is wrong.
It works in a top-down approach.

A counterpart to greedy is dynamic programming, where we check all the available options, not the single one.

If there’s only one answer and theres an optimal path you see - go for a greedy solution.
You should challenge your greedy approach with counter examples to understood that it works.
If it doesn't - try dynamic programming.

Problem can be solved greedily when:
1. The problem can be broken down into "overlapping sub-problems" - smaller versions of the original problem that are re-used multiple times.
2. The problem has an "optimal substructure" - an optimal solution can be formed from optimal solutions to the overlapping sub-problems of the original problem.
3. At each sub-problem we can (!) select the most optimal one every time, without reconsidering the previous steps once chosen.

### Comparison with dynamic programming

It's described well at the [docs/algorithms/dynamic_programming.md#comparison-with-greedy](./dynamic_programming.md#comparison-with-greedy)
