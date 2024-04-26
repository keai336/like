import sys
"""
<<
name:test
describe:测试
parse:/test
>>
"""
# 接收并打印传递的参数
# -*- coding: utf-8 -*-
sys.stdout.reconfigure(encoding='utf-8')

# 接收并打印传递的参数
for arg in sys.argv[1:]:
    print("Hello", arg)

