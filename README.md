# parse-leetcode-testcase
Parse LeetCode Testcase

# Install

```sh
go install github.com/liuxueyang/parse-leetcode-testcase@latest
```

# Usage

```
Usage of parse-leetcode-testcase:
  -i string
    	The input file to read from (default "raw.txt")
  -p string
    	The suffix name of the file to write to
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

The following command will create two identical files "input.txt" and "input_a.txt".

```sh
parse-leetcode-testcase -p a
```
