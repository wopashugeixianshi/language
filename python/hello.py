#!/usr/bin/env python3
# 加载BCC库
from bcc import BPF

# 加载eBPF内核态程序
b = BPF(src_file="hello.c")

# 将eBPF程序挂载到kprobe(openat类似于open)
b.attach_kprobe(event="do_sys_openat2", fn_name="hello_world")

# 读取并打印eBPF内核程序输出的数据
b.trace_print()
