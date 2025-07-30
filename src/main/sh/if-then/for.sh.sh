#!/bin/bash

# for循环假定每个值都是用空格分割的, 除了空格, 还有制表符和换行符(环境变量IFS中, 可以通过IFS=指定, 比如IFS=:指定:为分隔符)
list1="1 2 3 4 5"

# 如果带有特殊字符或者空格，需要转义或者用双引号括起来
list2="I don't know if this'll work"

for var in I don\'t know if "this'll" work
do
  echo "var: $var"
done


for var in $list2
do
  echo "var: $var"
done


# 从命令中读取列表
for var in "$(ls)"
do
  echo "var: $var"
done


# 通配符
for file in  /pengchao/*
do
 if [ -f "$file" ]; then
  echo "$file is file"
 elif [ -d "$file" ]; then
  echo "$file is directory"
 else
  echo "$file is other"
 fi
done



# c语言风格：
for ((i=0, j=10; i<5; i++,j--))
do
  echo "i: $i, j: $j"
done


# 处理循环的输出
# 输出到文件
for ((val=0; val<10; val++))
do
 echo "val: $val"
done > file.txt

# 排序后输出
for ((val=0; val<10; val++))
do
 echo "val: $val"
done | sort




