#!/bin/bash

# shift命令可以将命令列表整体向左移, 并从左边开始丢弃命令
while [ $# -ne 0 ]; do
    echo "$1"
    shift
done
