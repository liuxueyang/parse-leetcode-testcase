# parse-leetcode-testcase
Parse LeetCode Testcase

# Install

```sh
go install github.com/liuxueyang/parse-leetcode-testcase@latest
```

# Usage

```
Usage of ./parse-leetcode-testcase:
  -companion
        Export to competitive companion format
  -i string
        The input file to read from (default "raw.txt")
  -o string
        The output file to write to (default "input.txt")
  -p string
        The suffix name of the file to write to
  -v    Enable verbose output
```

Copy LeetCode testcase to text file "raw.txt". Example

```txt
示例 1：

输入：nums = [1,1,1,0,0,0,1,1,1,1,0], K = 2
输出：6
解释：[1,1,1,0,0,1,1,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 6。
示例 2：

输入：nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], K = 3
输出：10
解释：[0,0,1,1,1,1,1,1,1,1,1,1,0,0,0,1,1,1,1]
粗体数字从 0 翻转到 1，最长的子数组长度为 10。
```

Run

```sh
parse-leetcode-testcase
```

Outputs a file "input.txt" with a format similar to Codeforces input.

```txt
11
1 1 1 0 0 0 1 1 1 1 0 
2
19
0 0 1 1 0 0 1 1 1 0 1 1 0 0 0 1 1 1 1 
3
```

Another example of input file "raw.txt":

```txt
Example 1:

Input: nums = [2,4,6,5,7]

Output: 3

Explanation:

Swapping 5 and 6, the array becomes [2,4,5,6,7]

Swapping 5 and 4, the array becomes [2,5,4,6,7]

Swapping 6 and 7, the array becomes [2,5,4,7,6]. The array is now a valid arrangement. Thus, the answer is 3.

Example 2:

Input: nums = [2,4,5,7]

Output: 1

Explanation:

By swapping 4 and 5, the array becomes [2,5,4,7], which is a valid arrangement. Thus, the answer is 1.

Example 3:

Input: nums = [1,2,3]

Output: 0

Explanation:

The array is already a valid arrangement. Thus, no operations are needed.

Example 4:

Input: nums = [4,5,6,8]

Output: -1

Explanation:

No valid arrangement is possible. Thus, the answer is -1.
```

Output file:

```txt
5
2 4 6 5 7 
4
2 4 5 7 
3
1 2 3 
4
4 5 6 8
```

The following command will create two identical files "input.txt" and "input_a.txt".

```sh
parse-leetcode-testcase -p a
```

When your code failed to pass the LeetCode, you can copy the test case from the LeetCode page and paste it into "raw.txt", then run the command again to generate the input file.

Example of input file "raw.txt":

```txt
示例 1：

输入： start = [6,0,3], d = 2

输出： 4

解释：

可以选择整数 8, 0 和 4 获得最大可能得分，得分为 min(|8 - 0|, |8 - 4|, |0 - 4|)，等于 4。

示例 2：

输入： start = [2,6,13,13], d = 5

输出： 5

解释：

可以选择整数 2, 7, 13 和 18 获得最大可能得分，得分为 min(|2 - 7|, |2 - 13|, |2 - 18|, |7 - 13|, |7 - 18|, |13 - 18|)，等于 5。

start =
[1000000000,0]
d =
1000000000
----
start =
[1000000000,0]
d =
1000000000
```
