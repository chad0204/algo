#!/bin/bash

# 每次迭代test都会被执行
a=10
while [ $a -gt 0 ]; do
  #   a=$[$a - 1]
  ((tmp=a-1))
  a=$tmp
#  echo "a=$a"
done

# 如果是多层
v=10
while echo "v=$v"; do
  if [ $v -eq 5 ]; then
    break
  fi
  v=$((v-1))
done




# 和while的test相反, 不满足才行
b=10
until [ $b -lt 0 ]; do
  b=$((b-1))
#  echo "b=$b"
done