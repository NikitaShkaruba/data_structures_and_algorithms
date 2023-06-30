# Segment Tree

### Simple Segment tree

Is an abstraction similar to Prefix Sum. You create it for `O(n)` time and make range sum queries in `O(1)` time.
The new advantage is that you can also update elements in it in `O(logn)` time. 

Checkout `TestSegmentTree_Update` in [segment_tree_test.go](../../src/data_structures/segment_tree_test.go) for usage examples.

### Switch Segment tree

Basically stores booleans instead of actual data. You can use it to count how many elements are enabled in you data.
Switch segment tree also has very quick `flipRange` method to update the whole range for only `O(logn)` time, which is good!
Maybe it's possible to write the same method for the simple segment tree, but I have never given it a lot of thought currently.
