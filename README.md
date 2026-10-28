# test4py

Algorithm practice in Python and Go. Both languages use the same directory names, so each problem's two versions line up.

```
python/
  leetcode/   pNNNN_<slug>.py
  interview_easy/<category>/  20 most-asked Easy problems
  basics/     sorting, heap, tree traversal, graph, ...
  misc/       interview questions and scratch code
go/
  leetcode/   pNNNN_<slug>/solution.go + solution_test.go (one package per problem)
  interview_easy/<category>/pNNNN_<slug>/
  basics/     same topics as python/basics, one package each
  misc/       finished versions of the misc problems
  internal/   shared helpers (linkedlist.ListNode, bintree.TreeNode)
```

Run the Go tests: `cd go && go test ./...`

## LeetCode

| #    | Problem | Tags | Python | Go |
|------|---------|------|--------|----|
| 1    | [Two Sum](https://leetcode.com/problems/two-sum/) | hash | [py](python/leetcode/p0001_two_sum.py) | [go](go/leetcode/p0001_two_sum/solution.go) |
| 2    | [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/) | linked list | [py](python/leetcode/p0002_add_two_numbers.py) | [go](go/leetcode/p0002_add_two_numbers/solution.go) |
| 3    | [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | sliding window, hash | [py](python/leetcode/p0003_longest_substring_without_repeating_characters.py) | [go](go/leetcode/p0003_longest_substring_without_repeating_characters/solution.go) |
| 5    | [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/) | string, expand around center | [py](python/leetcode/p0005_longest_palindromic_substring.py) | [go](go/leetcode/p0005_longest_palindromic_substring/solution.go) |
| 6    | [Zigzag Conversion](https://leetcode.com/problems/zigzag-conversion/) | string, simulation | [py](python/leetcode/p0006_zigzag_conversion.py) | [go](go/leetcode/p0006_zigzag_conversion/solution.go) |
| 7    | [Reverse Integer](https://leetcode.com/problems/reverse-integer/) | math | [py](python/leetcode/p0007_reverse_integer.py) | [go](go/leetcode/p0007_reverse_integer/solution.go) |
| 9    | [Palindrome Number](https://leetcode.com/problems/palindrome-number/) | math, two pointers | [py](python/leetcode/p0009_palindrome_number.py) | [go](go/leetcode/p0009_palindrome_number/solution.go) |
| 11   | [Container With Most Water](https://leetcode.com/problems/container-with-most-water/) | two pointers | [py](python/leetcode/p0011_container_with_most_water.py) | [go](go/leetcode/p0011_container_with_most_water/solution.go) |
| 12   | [Integer to Roman](https://leetcode.com/problems/integer-to-roman/) | math, greedy | [py](python/leetcode/p0012_integer_to_roman.py) | [go](go/leetcode/p0012_integer_to_roman/solution.go) |
| 13   | [Roman to Integer](https://leetcode.com/problems/roman-to-integer/) | hash, string | [py](python/leetcode/p0013_roman_to_integer.py) | [go](go/leetcode/p0013_roman_to_integer/solution.go) |
| 14   | [Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/) | string | [py](python/leetcode/p0014_longest_common_prefix.py) | [go](go/leetcode/p0014_longest_common_prefix/solution.go) |
| 15   | [3Sum](https://leetcode.com/problems/3sum/) | sort, two pointers | [py](python/leetcode/p0015_3sum.py) | [go](go/leetcode/p0015_3sum/solution.go) |
| 17   | [Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/) | backtracking | [py](python/leetcode/p0017_letter_combinations_of_a_phone_number.py) | [go](go/leetcode/p0017_letter_combinations_of_a_phone_number/solution.go) |
| 19   | [Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) | linked list, two pointers | [py](python/leetcode/p0019_remove_nth_node_from_end_of_list.py) | [go](go/leetcode/p0019_remove_nth_node_from_end_of_list/solution.go) |
| 20   | [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) | stack | [py](python/leetcode/p0020_valid_parentheses.py) | [go](go/leetcode/p0020_valid_parentheses/solution.go) |
| 21   | [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) | linked list | [py](python/leetcode/p0021_merge_two_sorted_lists.py) | [go](go/leetcode/p0021_merge_two_sorted_lists/solution.go) |
| 22   | [Generate Parentheses](https://leetcode.com/problems/generate-parentheses/) | backtracking | [py](python/leetcode/p0022_generate_parentheses.py) | [go](go/leetcode/p0022_generate_parentheses/solution.go) |
| 26   | [Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/) | two pointers | [py](python/leetcode/p0026_remove_duplicates_from_sorted_array.py) | [go](go/leetcode/p0026_remove_duplicates_from_sorted_array/solution.go) |
| 28   | [Implement strStr()](https://leetcode.com/problems/implement-strstr/) | string | [py](python/leetcode/p0028_implement_strstr.py) | [go](go/leetcode/p0028_implement_strstr/solution.go) |
| 29   | [Divide Two Integers](https://leetcode.com/problems/divide-two-integers/) | math, bit | [py](python/leetcode/p0029_divide_two_integers.py) | [go](go/leetcode/p0029_divide_two_integers/solution.go) |
| 33   | [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/) | binary search | [py](python/leetcode/p0033_search_in_rotated_sorted_array.py) | [go](go/leetcode/p0033_search_in_rotated_sorted_array/solution.go) |
| 34   | [Find First and Last Position of Element in Sorted Array](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/) | binary search | [py](python/leetcode/p0034_find_first_and_last_position_of_element_in_sorted_array.py) | [go](go/leetcode/p0034_find_first_and_last_position_of_element_in_sorted_array/solution.go) |
| 36   | [Valid Sudoku](https://leetcode.com/problems/valid-sudoku/) | hash | [py](python/leetcode/p0036_valid_sudoku.py) | [go](go/leetcode/p0036_valid_sudoku/solution.go) |
| 38   | [Count and Say](https://leetcode.com/problems/count-and-say/) | string, recursion | [py](python/leetcode/p0038_count_and_say.py) | [go](go/leetcode/p0038_count_and_say/solution.go) |
| 41   | [First Missing Positive](https://leetcode.com/problems/first-missing-positive/) | array | [py](python/leetcode/p0041_first_missing_positive.py) | [go](go/leetcode/p0041_first_missing_positive/solution.go) |
| 45   | [Jump Game II](https://leetcode.com/problems/jump-game-ii/) | greedy, DP | [greedy](python/leetcode/p0045_jump_game_ii.py), [top-down](python/leetcode/p0045_jump_game_ii_topdown.py) | [go](go/leetcode/p0045_jump_game_ii/solution.go) |
| 48   | [Rotate Image](https://leetcode.com/problems/rotate-image/) | matrix | [py](python/leetcode/p0048_rotate_image.py) | [go](go/leetcode/p0048_rotate_image/solution.go) |
| 49   | [Group Anagrams](https://leetcode.com/problems/group-anagrams/) | hash, sort | [py](python/leetcode/p0049_group_anagrams.py) | [go](go/leetcode/p0049_group_anagrams/solution.go) |
| 53   | [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/) | DP, Kadane | [py](python/leetcode/p0053_maximum_subarray.py) | [go](go/leetcode/p0053_maximum_subarray/solution.go) |
| 54   | [Spiral Matrix](https://leetcode.com/problems/spiral-matrix/) | matrix, simulation | [py](python/leetcode/p0054_spiral_matrix.py) | [go](go/leetcode/p0054_spiral_matrix/solution.go) |
| 55   | [Jump Game](https://leetcode.com/problems/jump-game/) | greedy, DP | [py](python/leetcode/p0055_jump_game.py) | [go](go/leetcode/p0055_jump_game/solution.go) |
| 168  | [Excel Sheet Column Title](https://leetcode.com/problems/excel-sheet-column-title/) | math | [py](python/leetcode/p0168_excel_sheet_column_title.py) | [go](go/leetcode/p0168_excel_sheet_column_title/solution.go) |
| 588  | [Design In-Memory File System](https://leetcode.com/problems/design-in-memory-file-system/) | design, trie | [py](python/leetcode/p0588_design_in_memory_file_system.py) | [go](go/leetcode/p0588_design_in_memory_file_system/solution.go) |
| 1306 | [Jump Game III](https://leetcode.com/problems/jump-game-iii/) | DFS, BFS | [py](python/leetcode/p1306_jump_game_iii.py) | [go](go/leetcode/p1306_jump_game_iii/solution.go) |

## Interview Easy

The 20 most frequently asked LeetCode Easy problems, grouped by category under `python/interview_easy/<category>/` and `go/interview_easy/<category>/`. Kept separate from the problems above.

| Category | # | Problem | Python | Go |
|----------|---|---------|--------|----|
| Array / Hash | 121 | [Best Time to Buy and Sell Stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/) | [py](python/interview_easy/array_hash/p0121_best_time_to_buy_and_sell_stock.py) | [go](go/interview_easy/array_hash/p0121_best_time_to_buy_and_sell_stock/solution.go) |
| Array / Hash | 169 | [Majority Element](https://leetcode.com/problems/majority-element/) | [py](python/interview_easy/array_hash/p0169_majority_element.py) | [go](go/interview_easy/array_hash/p0169_majority_element/solution.go) |
| Array / Hash | 217 | [Contains Duplicate](https://leetcode.com/problems/contains-duplicate/) | [py](python/interview_easy/array_hash/p0217_contains_duplicate.py) | [go](go/interview_easy/array_hash/p0217_contains_duplicate/solution.go) |
| Array / Hash | 242 | [Valid Anagram](https://leetcode.com/problems/valid-anagram/) | [py](python/interview_easy/array_hash/p0242_valid_anagram.py) | [go](go/interview_easy/array_hash/p0242_valid_anagram/solution.go) |
| Two Pointers | 88 | [Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/) | [py](python/interview_easy/two_pointers/p0088_merge_sorted_array.py) | [go](go/interview_easy/two_pointers/p0088_merge_sorted_array/solution.go) |
| Two Pointers | 125 | [Valid Palindrome](https://leetcode.com/problems/valid-palindrome/) | [py](python/interview_easy/two_pointers/p0125_valid_palindrome.py) | [go](go/interview_easy/two_pointers/p0125_valid_palindrome/solution.go) |
| Two Pointers | 283 | [Move Zeroes](https://leetcode.com/problems/move-zeroes/) | [py](python/interview_easy/two_pointers/p0283_move_zeroes.py) | [go](go/interview_easy/two_pointers/p0283_move_zeroes/solution.go) |
| Linked List | 141 | [Linked List Cycle](https://leetcode.com/problems/linked-list-cycle/) | [py](python/interview_easy/linked_list/p0141_linked_list_cycle.py) | [go](go/interview_easy/linked_list/p0141_linked_list_cycle/solution.go) |
| Linked List | 206 | [Reverse Linked List](https://leetcode.com/problems/reverse-linked-list/) | [py](python/interview_easy/linked_list/p0206_reverse_linked_list.py) | [go](go/interview_easy/linked_list/p0206_reverse_linked_list/solution.go) |
| Linked List | 234 | [Palindrome Linked List](https://leetcode.com/problems/palindrome-linked-list/) | [py](python/interview_easy/linked_list/p0234_palindrome_linked_list.py) | [go](go/interview_easy/linked_list/p0234_palindrome_linked_list/solution.go) |
| Stack / Queue | 232 | [Implement Queue using Stacks](https://leetcode.com/problems/implement-queue-using-stacks/) | [py](python/interview_easy/stack_queue/p0232_implement_queue_using_stacks.py) | [go](go/interview_easy/stack_queue/p0232_implement_queue_using_stacks/solution.go) |
| Binary Search | 278 | [First Bad Version](https://leetcode.com/problems/first-bad-version/) | [py](python/interview_easy/binary_search/p0278_first_bad_version.py) | [go](go/interview_easy/binary_search/p0278_first_bad_version/solution.go) |
| Binary Search | 704 | [Binary Search](https://leetcode.com/problems/binary-search/) | [py](python/interview_easy/binary_search/p0704_binary_search.py) | [go](go/interview_easy/binary_search/p0704_binary_search/solution.go) |
| Tree | 101 | [Symmetric Tree](https://leetcode.com/problems/symmetric-tree/) | [py](python/interview_easy/tree/p0101_symmetric_tree.py) | [go](go/interview_easy/tree/p0101_symmetric_tree/solution.go) |
| Tree | 104 | [Maximum Depth of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/) | [py](python/interview_easy/tree/p0104_maximum_depth_of_binary_tree.py) | [go](go/interview_easy/tree/p0104_maximum_depth_of_binary_tree/solution.go) |
| Tree | 226 | [Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/) | [py](python/interview_easy/tree/p0226_invert_binary_tree.py) | [go](go/interview_easy/tree/p0226_invert_binary_tree/solution.go) |
| Tree | 543 | [Diameter of Binary Tree](https://leetcode.com/problems/diameter-of-binary-tree/) | [py](python/interview_easy/tree/p0543_diameter_of_binary_tree.py) | [go](go/interview_easy/tree/p0543_diameter_of_binary_tree/solution.go) |
| Dynamic Programming | 70 | [Climbing Stairs](https://leetcode.com/problems/climbing-stairs/) | [py](python/interview_easy/dp/p0070_climbing_stairs.py) | [go](go/interview_easy/dp/p0070_climbing_stairs/solution.go) |
| Bit Manipulation | 136 | [Single Number](https://leetcode.com/problems/single-number/) | [py](python/interview_easy/bit/p0136_single_number.py) | [go](go/interview_easy/bit/p0136_single_number/solution.go) |
| Bit Manipulation | 268 | [Missing Number](https://leetcode.com/problems/missing-number/) | [py](python/interview_easy/bit/p0268_missing_number.py) | [go](go/interview_easy/bit/p0268_missing_number/solution.go) |

## Basics

| Topic | Python | Go |
|-------|--------|----|
| Sorting (bubble, insertion, selection, merge, heap, quick) and shuffling | [sort.py](python/basics/sort.py) | [go](go/basics/sorting/sorting.go) |
| Binary heap | [heap.py](python/basics/heap.py) | [go](go/basics/heap/heap.go) |
| Binary tree traversal (recursive and iterative) | [binary_tree.py](python/basics/binary_tree.py) | [go](go/basics/binarytree/binarytree.go) |
| Graph DFS / BFS | [graph.py](python/basics/graph.py) | [go](go/basics/graph/graph.go) |
| Fibonacci (recursion, memoization, bottom-up) | [fibonacci.py](python/basics/fibonacci.py) | [go](go/basics/fibonacci/fibonacci.go) |
| Longest increasing subsequence | [longest_increasing_subsequence.py](python/basics/longest_increasing_subsequence.py) | [go](go/basics/lis/lis.go) |
| Tower of Hanoi | [hanoi.py](python/basics/hanoi.py) | [go](go/basics/hanoi/hanoi.go) |
| Find second largest | [find_second_largest.py](python/basics/find_second_largest.py) | [go](go/basics/secondlargest/secondlargest.go) |

## Misc

- [order_book.py](python/misc/order_book.py): order book matching (unfinished)
- coffee machine placement (minimize total Manhattan distance to developers): [py](python/misc/coffee_machine.py) (unfinished), [go](go/misc/coffeemachine/coffeemachine.go) (median per axis)
- [scratch.py](python/misc/scratch.py): scratch code
