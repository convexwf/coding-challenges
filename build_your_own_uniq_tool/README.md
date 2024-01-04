# Build Your Own Uniq Tool (Not implemented)

Inspired By [Build Your Own uniq Tool | Coding Challenges](https://codingchallenges.fyi/challenges/challenge-uniq)

基本功能：
1. 读取输入文件，逐行比较内容，将重复的行去除，输出去重后的内容到输出文件。
2. 可以指定忽略大小写进行比较。
3. 可以统计重复行的数量。

扩展功能：
1. 支持对文件内容进行排序后再去重。
2. 支持指定输出重复的行。
3. 支持指定输出不重复的行。
4. 支持指定输出重复行的数量和具体位置。
5. 支持处理大文件，采用分块读取和处理的方式，减少内存占用。
6. 支持多个输入文件进行去重操作。
7. 支持指定输出重复行的行号。
8. 支持输出去重后的行数统计。
9. 支持输出去重前后的行数对比统计。
10. 支持自定义输出格式，如指定分隔符、字段顺序等。
11. 支持根据指定字段进行去重操作，而不是整行内容进行比较。
12. 支持对去重后的内容进行进一步处理，如去除空行、去除特定字符等操作。

## bazel build

npm install -g @bazel/bazelisk