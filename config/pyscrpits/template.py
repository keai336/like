"""
<<
name:testa
describe:测试
parse:/test
paracheck:^\d{2,5}$
>>
"""
import sys
sys.stdout.reconfigure(encoding='utf-8')
def orig_para(para):
    para= para.replace("₹"," ").replace("ℳ","/n")
    return para
arg=orig_para(sys.argv[1])
print(arg)