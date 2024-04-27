"""
<<
name:testa
describe:测试
parse:/test
paracheck:^\d{2,5}$
>>
"""
import sys
def orig_para(para):
    para= para.replace("₹"," ").replace("ℳ","/n")
    return para
# -*- coding: utf-8 -*-
sys.stdout.reconfigure(encoding='utf-8')
for arg in sys.argv[1:]:
    print(arg)