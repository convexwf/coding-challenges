# Coding Challenges - Build Your Own Word Count Tool

## 前言

首先看一下 GNU/Linux 系统中的 `wc` 命令的用法：

```bash
$ wc --help
Usage: wc [OPTION]... [FILE]...
  or:  wc [OPTION]... --files0-from=F
Print newline, word, and byte counts for each FILE, and a total line if
more than one FILE is specified.  A word is a non-zero-length sequence of
characters delimited by white space.

With no FILE, or when FILE is -, read standard input.

The options below may be used to select which counts are printed, always in
the following order: newline, word, character, byte, maximum line length.
  -c, --bytes            print the byte counts
  -m, --chars            print the character counts
  -l, --lines            print the newline counts
      --files0-from=F    read input from the files specified by
                           NUL-terminated names in file F;
                           If F is - then read names from standard input
  -L, --max-line-length  print the maximum display width
  -w, --words            print the word counts
      --help     display this help and exit
      --version  output version information and exit

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Full documentation at: <https://www.gnu.org/software/coreutils/wc>
or available locally via: info '(coreutils) wc invocation'
```

这里面的功能包括统计行数、单词数、字符数、字节数、最长行长度等。

当实现一个自己的wc工具时，除了基本的行数、字数、字符数统计功能外，还可以考虑扩展以下功能：

1. 统计代码行数：可以添加一个选项来统计文件中的代码行数，排除空行和注释行。
2. 统计单词频率：可以添加一个选项来统计文件中每个单词出现的频率。
3. 统计文件大小：可以添加一个选项来统计文件的大小，以字节为单位。
4. 统计文件类型：可以添加一个选项来统计不同类型文件的数量，如文本文件、图片文件、音频文件等。
5. 输出结果可视化：可以添加一个选项来将统计结果以图表或图形的形式展示出来，更直观地呈现数据。
6. 多文件同时统计：可以添加一个选项来同时统计多个文件的信息，比较不同文件之间的差异。
7. 自定义输出格式：可以添加一个选项来让用户自定义输出结果的格式，如指定输出字段的顺序或格式化输出样式。
8. 支持多种文件编码：可以添加一个选项来支持统计不同编码格式的文件，如UTF-8、GBK等。
9. 实时统计：可以添加一个选项来实时监控文件的变化并实时更新统计结果，方便用户及时了解文件的变化情况。

