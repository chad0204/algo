#!/bin/bash

# test1 if后的命令行执行状态码是0, 执行then后的命令, then 可以执行多行命令
if pwd;ls; then
  echo "It work"
  echo "It work too"
fi

# test2 if后的命令行执行状态码非0, 不执行then后的命令
testuser=root
if grep $testuser /etc/passwd; then
  echo "This command id work"
else
  echo "The user $testuser does not exist on this system."
fi
echo "abc"



# test3 嵌套elif
if pwd; then
  echo "It work"
elif grep $testuser /etc/passwd; then
  echo "This user $testuser is exist"
elif ls -d /home/$testuser; then
  echo "This $testuser is a directory"
else
  echo "The user $testuser does not exist on this system."
fi

