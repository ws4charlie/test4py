# test4py

Algorithm practice in Python and Go. Both languages use the same directory names, so each problem's two versions line up.

```
python/
  leetcode/   pNNNN_<slug>.py
  basics/     sorting, heap, tree traversal, graph, ...
  misc/       interview questions and scratch code
go/
  leetcode/   pNNNN_<slug>/solution.go + solution_test.go (one package per problem)
```

Run the Go tests: `cd go && go test ./...`

## LeetCode

| #    | Problem | Tags | Python | Go |
|------|---------|------|--------|----|
| 1    | [Two Sum](https://leetcode.com/problems/two-sum/) | hash | [py](python/leetcode/p0001_two_sum.py) | [go](go/leetcode/p0001_two_sum/solution.go) |
| 2    | [Add Two Numbers](https://leetcode.com/problems/add-two-numbers/) | linked list | [py](python/leetcode/p0002_add_two_numbers.py) | |
| 3    | [Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/) | sliding window, hash | [py](python/leetcode/p0003_longest_substring_without_repeating_characters.py) | [go](go/leetcode/p0003_longest_substring_without_repeating_characters/solution.go) |
| 5    | [Longest Palindromic Substring](https://leetcode.com/problems/longest-palindromic-substring/) | string, expand around center | [py](python/leetcode/p0005_longest_palindromic_substring.py) | [go](go/leetcode/p0005_longest_palindromic_substring/solution.go) |
| 6    | [Zigzag Conversion](https://leetcode.com/problems/zigzag-conversion/) | string, simulation | [py](python/leetcode/p0006_zigzag_conversion.py) | |
| 7    | [Reverse Integer](https://leetcode.com/problems/reverse-integer/) | math | [py](python/leetcode/p0007_reverse_integer.py) | |
| 9    | [Palindrome Number](https://leetcode.com/problems/palindrome-number/) | math, two pointers | [py](python/leetcode/p0009_palindrome_number.py) | |
| 11   | [Container With Most Water](https://leetcode.com/problems/container-with-most-water/) | two pointers | [py](python/leetcode/p0011_container_with_most_water.py) | [go](go/leetcode/p0011_container_with_most_water/solution.go) |
| 12   | [Integer to Roman](https://leetcode.com/problems/integer-to-roman/) | math, greedy | [py](python/leetcode/p0012_integer_to_roman.py) | |
| 13   | [Roman to Integer](https://leetcode.com/problems/roman-to-integer/) | hash, string | [py](python/leetcode/p0013_roman_to_integer.py) | |
| 14   | [Longest Common Prefix](https://leetcode.com/problems/longest-common-prefix/) | string | [py](python/leetcode/p0014_longest_common_prefix.py) | [go](go/leetcode/p0014_longest_common_prefix/solution.go) |
| 15   | [3Sum](https://leetcode.com/problems/3sum/) | sort, two pointers | [py](python/leetcode/p0015_3sum.py) | [go](go/leetcode/p0015_3sum/solution.go) |
| 17   | [Letter Combinations of a Phone Number](https://leetcode.com/problems/letter-combinations-of-a-phone-number/) | backtracking | [py](python/leetcode/p0017_letter_combinations_of_a_phone_number.py) | |
| 19   | [Remove Nth Node From End of List](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) | linked list, two pointers | [py](python/leetcode/p0019_remove_nth_node_from_end_of_list.py) | |
| 20   | [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) | stack | [py](python/leetcode/p0020_valid_parentheses.py) | [go](go/leetcode/p0020_valid_parentheses/solution.go) |
| 21   | [Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/) | linked list | [py](python/leetcode/p0021_merge_two_sorted_lists.py) | |
| 22   | [Generate Parentheses](https://leetcode.com/problems/generate-parentheses/) | backtracking | [py](python/leetcode/p0022_generate_parentheses.py) | |
| 26   | [Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/) | two pointers | [py](python/leetcode/p0026_remove_duplicates_from_sorted_array.py) | [go](go/leetcode/p0026_remove_duplicates_from_sorted_array/solution.go) |
| 28   | [Implement strStr()](https://leetcode.com/problems/implement-strstr/) | string | [py](python/leetcode/p0028_implement_strstr.py) | |
| 29   | [Divide Two Integers](https://leetcode.com/problems/divide-two-integers/) | math, bit | [py](python/leetcode/p0029_divide_two_integers.py) | |
| 33   | [Search in Rotated Sorted Array](https://leetcode.com/problems/search-in-rotated-sorted-array/) | binary search | [py](python/leetcode/p0033_search_in_rotated_sorted_array.py) | |
| 34   | [Find First and Last Position of Element in Sorted Array](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/) | binary search | [py](python/leetcode/p0034_find_first_and_last_position_of_element_in_sorted_array.py) | |
| 36   | [Valid Sudoku](https://leetcode.com/problems/valid-sudoku/) | hash | [py](python/leetcode/p0036_valid_sudoku.py) | |
| 38   | [Count and Say](https://leetcode.com/problems/count-and-say/) | string, recursion | [py](python/leetcode/p0038_count_and_say.py) | |
| 41   | [First Missing Positive](https://leetcode.com/problems/first-missing-positive/) | array | [py](python/leetcode/p0041_first_missing_positive.py) | |
| 45   | [Jump Game II](https://leetcode.com/problems/jump-game-ii/) | greedy, DP | [greedy](python/leetcode/p0045_jump_game_ii.py), [top-down](python/leetcode/p0045_jump_game_ii_topdown.py) | |
| 48   | [Rotate Image](https://leetcode.com/problems/rotate-image/) | matrix | [py](python/leetcode/p0048_rotate_image.py) | |
| 49   | [Group Anagrams](https://leetcode.com/problems/group-anagrams/) | hash, sort | [py](python/leetcode/p0049_group_anagrams.py) | |
| 53   | [Maximum Subarray](https://leetcode.com/problems/maximum-subarray/) | DP, Kadane | [py](python/leetcode/p0053_maximum_subarray.py) | |
| 54   | [Spiral Matrix](https://leetcode.com/problems/spiral-matrix/) | matrix, simulation | [py](python/leetcode/p0054_spiral_matrix.py) | |
| 55   | [Jump Game](https://leetcode.com/problems/jump-game/) | greedy, DP | [py](python/leetcode/p0055_jump_game.py) | |
| 168  | [Excel Sheet Column Title](https://leetcode.com/problems/excel-sheet-column-title/) | math | [py](python/leetcode/p0168_excel_sheet_column_title.py) | |
| 588  | [Design In-Memory File System](https://leetcode.com/problems/design-in-memory-file-system/) | design, trie | [py](python/leetcode/p0588_design_in_memory_file_system.py) | |
| 1306 | [Jump Game III](https://leetcode.com/problems/jump-game-iii/) | DFS, BFS | [py](python/leetcode/p1306_jump_game_iii.py) | |

## Basics

| Topic | Python | Go |
|-------|--------|----|
| Sorting (bubble, insertion, selection, merge, heap, quick) and shuffling | [sort.py](python/basics/sort.py) | |
| Binary heap | [heap.py](python/basics/heap.py) | |
| Binary tree traversal (recursive and iterative) | [binary_tree.py](python/basics/binary_tree.py) | |
| Graph DFS / BFS | [graph.py](python/basics/graph.py) | |
| Fibonacci (recursion, memoization, bottom-up) | [fibonacci.py](python/basics/fibonacci.py) | |
| Longest increasing subsequence | [longest_increasing_subsequence.py](python/basics/longest_increasing_subsequence.py) | |
| Tower of Hanoi | [hanoi.py](python/basics/hanoi.py) | |
| Find second largest | [find_second_largest.py](python/basics/find_second_largest.py) | |

## Misc

- [order_book.py](python/misc/order_book.py): order book matching (unfinished)
- [coffee_machine.py](python/misc/coffee_machine.py): best place for an office coffee machine (unfinished)
- [scratch.py](python/misc/scratch.py): scratch code
