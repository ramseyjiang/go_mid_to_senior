A heap is a different thing altogether. It is a tree-like data structure that satisfies the heap property. A stack is like a stack of pebbles, but a
heap is a heap of pebbles.

A popular heap implementation is the binary heap, where each node has at most 2 child nodes. There are 2 types of heaps。 The first is the min heap where
the heap property is such that the key of the parent node is always smaller than that of the child nodes. The second is the max heap where the heap
property, is such that the key of the parent node is always larger than that of the child nodes.

How we can find the parent node is to subtract 1 from the child and divide by 2？ parent = (child - 1)/2 Left child = (2 * parent) + 1 Right child = (2 *
parent) + 2

In a max heap it does not matter which sibling is left or right as long as both are smaller than the parent. This means that there are many possible
ways a heap can organise itself.
