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


# 数值比较 -eq -ne -ge -gt -le -lt; 注意bash不能进行浮点数比较, zsh可以
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


# 双括号命令允许你在比较过程中使用高级数学表达式。test命令只能在比较中使用简单的算术操作。test重点是比较, 双括号是运算
# val++; val--; ++val; --val; !; **(幂运算); <<; >>; &; |; &&; ||
val=10
if (( val ** 2 > 90 )); then
  ((val2 = val ** 2))
  echo "The square of $val is $val2"
fi

a=10
while [ $a -gt 0 ]; do
  ((tmp=a-1))
  a=$tmp
  echo "a=$a"
done

# 双方括号, 比较字符串
if [[ $USER == r* ]]
then
 echo "Hello $USER"
else
 echo "Sorry, I do not know you"
fi


