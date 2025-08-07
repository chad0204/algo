#!/bin/bash
# 命令行参数
# 位置参数$0表示当前脚本名 ./pos_param.sh 和bash pos_param.sh返回的结果不一样, 可以通过basename只输出脚本名
echo $0

TEST_SH_HOME=$(dirname "$0")
TEST_SH_NAME=$(basename "$0")
TEST_SH_PARENT_HOME=$(cd $(dirname "$0");cd ../;  pwd)
# exit处理cd失败的情况, 避免脚本向下错误执行; $(...)加双引号避免结果有空格
TEST_SH_PARENT_HOME2=$(cd "$(dirname "$0")" || exit ;cd ../;  pwd)
#``写法
TEST_SH_PARENT_HOME3=$(cd "`dirname "$0"`" || exit ;cd ../;  pwd)


echo "TEST_HOME is: $TEST_SH_HOME"
echo "TEST_SH_NAME is: $TEST_SH_NAME"
echo "TEST_SH_PARENT_HOME is: $TEST_SH_PARENT_HOME"
echo "TEST_SH_PARENT_HOME2 is: $TEST_SH_PARENT_HOME2"
echo "TEST_SH_PARENT_HOME3 is: $TEST_SH_PARENT_HOME3"




# $1表示第一个位置参数 一直到$9. 如果变量超过10个, 可以使用${10} ${11}等方式访问
echo $1
echo ${10}


# 可以统计执行脚本的命令行带有多少参数
echo $#
# 小技巧，echo $#是命令总数n, ${n}就是最后一个参数的值, 但是不能写${$#}
echo ${!#}

# 将所有参数作为一个整体字符串输出, 但是遍历时是一个单词
echo "$*"
# 将每个参数作为独立字符串输出, 可遍历出每一个单词, 适合for in "$@"
echo "$@"

