#!/bin/bash
# ./option.sh -e 233 -v -- a b c

# 命令行选项, 选项和参数并没有实际的区分, 都是通过位置参数获取. 选项更像是一种约束
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
