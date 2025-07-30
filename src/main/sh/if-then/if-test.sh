#!/bin/bash
a=1
if test $a; then
  echo "a has value"
else
  echo "a has no value"
fi

# test等价于[], [空格 命令 空格]
b=""
if [ $b ]; then
  echo "b has value"
else
  echo "b has no value"
fi


# 数值比较 eq ge gt le lt ne; 注意bash不能进行浮点数比较, zsh可以
c="1"
if [ $a -eq "$c" ]; then
  echo "a eq c"
else
  echo "a not eq c"
fi


# 字符比较 = != <(需要加\转义) > -n(字符串长度是否非0) -z(字符串长度是否为0) ; 注意bash不能进行浮点数比较, zsh可以
str1="abc"
str2=""
if [ -n "$str1" ]; then
  echo "str1 length not 0"
else
  echo "str1 length is 0"
fi



# 文件比较
# -d file 检查文件是否存在且是一个目录
# -e file 检查文件是否存在
# -f file 检查文件是否存在且是一个普通文件
# -r file 检查文件是否存在且可读
# -s file 检查文件是否存在且非空
# -w file 检查文件是否存在且可写
# -x file 检查文件是否存在且可执行
# -O file 检查文件是否存在且是当前用户的所有者
# -G file 检查文件是否存在且是当前用户组的所有者
# file1 -nt file2 # 检查file1是否比file2新
# file1 -ot file2 # 检查file1是否比file2旧

#!/bin/bash
jump_dir=/pengchao
if [ -d "$jump_dir" ]; then
 cd $jump_dir
 ls
else
 echo "$jump_dir not directory"
# 打印上一个命令的执行状态码
 echo $?
fi


# 复合条件 && ||
a=1
b=2
if [ $a -eq 1 ] && [ $b -eq 2 ]; then
  echo "a eq 1 and b eq 2"
else
  echo "a not eq 1 or b not eq 2"
fi

