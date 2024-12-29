#!/bin/bash

# 获取 CPU 使用率
get_cpu_usage() {
    # 使用 top 命令获取 CPU 使用率，取第二行，第一列数据
    cpu_usage=$(top -bn1 | grep "Cpu(s)" | awk '{print $2}' | cut -d '%' -f 1)
    echo "CPU Usage: $cpu_usage%"
}

# 获取内存使用率
get_memory_usage() {
    # 使用 free 命令获取内存信息
    # 计算内存使用率：(总内存 - 空闲内存) / 总内存 * 100
    total_memory=$(free | grep Mem | awk '{print $2}')
    free_memory=$(free | grep Mem | awk '{print $4}')
    used_memory=$((total_memory - free_memory))
    memory_usage=$((used_memory * 100 / total_memory))
    echo "Memory Usage: $memory_usage%"
}

# 主函数，调用上面两个函数
main() {
    get_cpu_usage
    get_memory_usage
}

# 调用主函数
main