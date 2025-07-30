#!/bin/bash

# break跳出循环, 如果是多层循环, 使用break n可跳出n层循环
# continue也能跳过多层循环
level1=3
level2=3
level3=3
for ((i=0; i<=level1; i++)); do
  for ((j=0; j<=level2; j++)); do
    for ((k=0; k<=level3; k++)); do
      echo "i=$i, j=$j, k=$k"
      if [ $k -eq 2 ]; then
        echo "Breaking out of the innermost loop"
#        break  # 跳出最内层循环
        break 3 # 跳出3层循环
      fi
    done
    echo "Exited innermost loop, j=$j"
  done
  echo "Exited middle loop, i=$i"
done

