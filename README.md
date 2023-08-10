# Data Structures and Algorithms

![ci_badge](https://img.shields.io/github/actions/workflow/status/NikitaShkaruba/data_structures_and_algorithms/ci.yml)
![go_version_badge](https://img.shields.io/github/go-mod/go-version/NikitaShkaruba/data_structures_and_algorithms)
![license_badge](https://img.shields.io/github/license/NikitaShkaruba/data_structures_and_algorithms)

### About

Standard go library is pretty weak and inconvenient for leetcoding, so I've decided to write a good go library myself.
I've defined all the generic data structures and generic misc functions in [src](./src), then with [leetcode_library_template_builder](./tools/leetcode_library_template_builder/main.go)
I concatenate all the files to [leetcode_library.template](./build/leetcode_library.template), which I later copy-paste to every leetcode solution,
enjoying the library I deserve...

Also, I have all the algorithm and data structures explained in [docs](docs) directory, check it out!
Those explanations contain some code snippets I haven't included in [leetcode_library.template](./build/leetcode_library.template),
because I think you should quickly code them yourself, and don't rely on the library heavily.

### Why standard go library is not enough for leetcode

- No generics support - I've manually defined too much `Max`, `Min`, `Abs`, `...` functions in my solutions.
- Awful [heap interface](https://pkg.go.dev/container/heap#example-package-IntHeap).
It doesn't work out of the box - you need to define a lot of error-prone code, and it's inconvenient to use afterwards.
- Not enough data structures. There is no `AVL tree`, `Trie`, `Union-find`, ...

### Project structure

- [build/leetcode_library.template](build/leetcode_library.template) - template I copy-paste to every solution, containing all the generic data structures and generic misc functions
- [docs](docs) - directory with algorithm and data structure explanations necessary for leetcode
- [src](src) - directory with data structures, algorithms, and misc functions necessary for leetcode. Each separate file gets concatenated to [leetcode_library.template](./build/leetcode_library.template)
- [tools](tools) - directory that helps to maintain [leetcode_library.template](./build/leetcode_library.template)
