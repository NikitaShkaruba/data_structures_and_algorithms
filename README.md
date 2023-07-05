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

### My problems with standard go library

- No generics support - I've defined too much `Max`, `Min`, `Abs`, `...` functions in my solutions manually
- Awful heap interface - who came up with [this sh*t](https://pkg.go.dev/container/heap#example-package-IntHeap)??? It's too complicated!!!
In order to use it you need to specify a lot of error-prone code
- Not enough data structures. There is no `AVL tree`, `Trie`, `Union-find`, ...

### Project structure

- [build/leetcode_library.template](build/leetcode_library.template) - template I copy-paste to every solution, containing all the generic data structures and generic misc functions
- [docs](docs) - directory with algorithm and data structure explanations necessary for leetcode
- [src](src) - directory with generic data structures and generic misc functions necessary for leetcode. Each separate file gets concatenated to [leetcode_library.template](./build/leetcode_library.template)
- [tools](tools) - directory that helps to maintain [leetcode_library.template](./build/leetcode_library.template)
