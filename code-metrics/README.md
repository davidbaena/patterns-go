
### Steps to add metrics to the code
1) Get modified files from a pull request

```
gh pr diff https://github.com/davidbaena/patterns-go/pull/1 --repo=davidbaena/patterns-go --name-only

creational/singleton/database.go
creational/singleton/database_test.go
creational/singleton/go.mod
creational/singleton/main.go
```

2) Run emerge to get fan-in, fan-out, louvain metrics
"patterns-go/creational/singleton/database.go": {
  "number-of-methods-in-file": 2,
  "sloc-in-file": 15,
  "file_result_dependency_graph_louvain-modularity-in-file": 3,
  "fan-in-dependency-graph": 0,
  "fan-out-dependency-graph": 1
}

### Metrics
Fan-In: Number of incoming dependencies. 
A <-- B
A <-- C
A <-- D
Fan-In of A = 3
sync <- database_test
sync <- database

Fan-Out: Numer of outgoing dependencies.
A --> B
A --> C
Fan-Out of A = 2
database_test -> sync
database_test -> testing

Louvain Modularity: Community detection algorithm to find communities in a graph. Example: How many communities are in the dependency graph

### References
https://www.math.unipd.it/~tullio/IS-1/2004/Approfondimenti/Fan-in_Fan-out.html#:~:text=Fan%2Din%20is%20the%20number,are%20counted%20with%20different%20rules