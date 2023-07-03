# Dynamic programming

Dynamic programming (DP) is an approach for solving a problem by checking all the options available.
In other words it's just brute force with optimizations, nothing more.
Everybody is scared of this pattern, because it's unusual to think of at first, but it's easy, really.

Problem can be solved by DP when:
1. The problem can be broken down into "overlapping sub-problems" - smaller versions of the original problem that are re-used multiple times.
2. The problem has an "optimal substructure" - an optimal solution can be formed from optimal solutions to the overlapping sub-problems of the original problem.
3. At each sub-problem we can't (!) select the most optimal one every time, without reconsidering the previous steps once chosen.

### How to use

Approach:

1. Find state variables
2. Find recurrence relation
3. Find base cases
4. Wire all those things together and start your DP

### When to use

Everybody explains it like this, even though I think this thinking is very shallow.

You should use DP when question contains one of those phrases:

- What is the minimum cost of doing ...
- What is the maximum profit from ...
- How many ways are there to do ...
- What is the longest possible ...
- Is it possible to reach a certain point ...

It helps to select DP for a beginner, but actually I don't like those phrase to algorithm mappings, so here's a better one for you.

You should use DP when:
- Future decisions depend on earlier decisions
- You can't solve it greedily

### Top-down or bottom-up

DP can be solved bottom-up (start from base-case, go for the target sub-problem) or top-down (start from target sub-problem, break it down till you go to base-case)

- A **bottom-up** implementation's runtime is usually faster, as iteration does not have the overhead that recursion does.
- A **top-down** implementation is usually much easier to write. This is because with recursion, the ordering of sub-problems does not matter, whereas with tabulation, we need to go through a logical ordering of solving sub-problems.

Usually top-down implementation is easy to implement, even though its cost is not perfect.
Start with it, and later try to convert to bottom-up approach if needed.
Or just solve everything top-down because why bother if there’s no TLE?
I always go for top-down :)

### Computing time and space complexity:

- One of the main points with DP is that we never repeat calculations, whether by tabulation or memoization, we only compute a state once.
  Because of this, the time complexity of a DP algorithm is directly tied to the number of possible states, so
  `Time complexity = number of states * time spending in each state`
- Usually is directly tied to the caching, so
  `Space complexity = number of states`

### Comparison with greedy

Both DP and [greedy](./greedy.md) are tools for optimization.
However, greedy algorithms look for a greedy choice, in the hopes of finding a global optimum.
Dynamic programming, on the other hand, finds the optimal solution to sub-problems and then makes an informed choice to combine the results of those sub-problems to find the most optimum solution.

You should always try to solve your problem with the greedy approach before going for DP.
If there’s only one answer and there's an optimal path you see - go for a greedy solution.
You should challenge your greedy approach with counter examples to understood that it works.
If it doesn't - try dynamic programming.
