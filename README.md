# Data Structures and Algorithms

![ci_badge](https://img.shields.io/github/actions/workflow/status/NikitaShkaruba/data_structures_and_algorithms/ci.yml)
![go_version_badge](https://img.shields.io/github/go-mod/go-version/NikitaShkaruba/data_structures_and_algorithms)
![license_badge](https://img.shields.io/github/license/NikitaShkaruba/data_structures_and_algorithms)

### About

Standard go library  weak and inconvenient for leetcoding, so I've decided to write a good library myself.
I've defined all the data structures and misc functions in `src`, then with `tools/leetcode_library_template_builder`
I concatenate all the files to `build/leetcode_library.template`, which I later copy-paste to every leetcode solution,
enjoying the library I deserve.

### My problems with standard go library

- No generics support - I've defined too much `Max`, `Min`, `Abs`, ... functions in my solutions manually
- Awful heap interface - who came up with [this sh*t](https://pkg.go.dev/container/heap)??? In order to use it you need
to define it, and there's a lot of error-prone code
- Not enough data structures. There is no `AVL tree`, `Trie`, `Union-find`, ...

### Project structure

- [src/algorithms](src/algorithms) - directory with algorithm explanations necessary for leetcode
- [src/data_structures](src/data_structures) - directory with generic data structures necessary for leetcode
- [src/stdlib](src/stdlib) - directory with generic misc functions that can be used everywhere
- [build/leetcode_library.template](build/leetcode_library.template) - template I copy-paste to every solution, containing all the data structures and misc functions
- [tools](tools) - directory that helps to maintain `leetcode_library.template`
