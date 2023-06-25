# Data Structures and Algorithms

![ci_badge](https://img.shields.io/github/actions/workflow/status/NikitaShkaruba/data_structures_and_algorithms/ci.yml)
![go_version_badge](https://img.shields.io/github/go-mod/go-version/NikitaShkaruba/data_structures_and_algorithms)
![license_badge](https://img.shields.io/github/license/NikitaShkaruba/data_structures_and_algorithms)

### About

Standard go library  weak and inconvenient for leetcoding, so I've decided to write a good library myself.
I've defined all the generic data structures and generic misc functions in `src`, then with `tools/leetcode_library_template_builder`
I concatenate all the files to `build/leetcode_library.template`, which I later copy-paste to every leetcode solution,
enjoying the library I deserve.

### My problems with standard go library

- No generics support - I've defined too much `Max`, `Min`, `Abs`, ... functions in my solutions manually
- Awful heap interface - who came up with [this sh*t](https://pkg.go.dev/container/heap)??? In order to use it you need
to define it, and there's a lot of error-prone code
- Not enough data structures. There is no `AVL tree`, `Trie`, `Union-find`, ...

### Project structure

- [build/leetcode_library.template](build/leetcode_library.template) - template I copy-paste to every solution, containing all the generic data structures and generic misc functions
- [docs/algorithms](docs/algorithms) - directory with algorithm explanations necessary for leetcode
- [src](src) - directory with generic data structures and generic misc functions necessary for leetcode. Each separate file gets concatenated to `leetcode_library.template`
- [tools](tools) - directory that helps to maintain `leetcode_library.template`
