#!/bin/bash

# getopt命令用于处理命令行选项和参数, 主要用于解析短选项和长选项

# 使用getopt解析命令行选项, 可以解析命令行，结果是 -a -b test1 -c -d -- test2 test3
# b后面有个冒号说明b带参数 -cd也是选项 其余得都是参数加--区分
#>getopt ab:cd -a test0 -b test1 test2 test3 -cd


# set设置位置参数(包括选项和参数), getopt解析输入得命令行
# hve是选项, 其中e后面有个冒号说明e带参数
# 输入 -vhe 233 a b c 解析成 -v -h -e 233 -- a b c
eval set -- "$(getopt -q vhe: "$@")"

# 打印下最终的解析结果, 或者echo "parse args = $*"
echo "parse args =" "$@"


# 执行选项和参数
while [ -n "$1" ]; do
  case "$1" in
    # 输出帮助信息
    -h | --help)
      echo "Usage: $(basename "$0") [options]"
      echo "Options:"
      echo "  -h, --help    Show this help message"
      echo "  -v, --version Show version information"
      echo "  -e, --execute Execute process"
      # 这里exit的话 下面参数就没法解析了
#      exit 0
      ;;
    # 实际命令选项
    -v | --version)
      echo "Version 1.0.0"
#      exit 0
      ;;
    # 实际命令选项, 选项带参数值, 当前选项是$1, 那紧跟得参数就是$2, 处理完参数还要shift把参数去掉
    -e | --execute) param="$2"
      echo "executed process $param"
      shift
#      exit 0
      ;;
    --) shift # 处理双破折号, 表示后面没有选项了, 只剩下参数, shift是因为这里跳出循环, 所以要把“--”也移出
      echo "Options over"
      break
      ;;
    *)
      echo "Unknown option: $1"
      exit 1 # 错误就不向下执行了
      ;;
  esac
  shift 1 # 移动到下一个选项
done


# 处理剩下的参数
echo "$@"
count=1
for param in "$@"; do
  echo "Parameter $count: $param"
  count=$((count + 1))
done