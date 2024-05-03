"""
<<
name:美图
describe:随机美图
parse:/mt
paracheck:^\d?$
>>
"""
import sys
import random
import os
sys.stdout.reconfigure(encoding='utf-8')
path = "C:\\Users\keai3\Desktop\Beauty-pictures-crawling-master\美女图集"
def one():
    whos = os.listdir(path)
    who = whos[random.randint(0,len(whos)-1)]
    whichs = os.listdir(path+"/"+who)
    which = whichs[random.randint(0,len(whichs)-1)]
    one = "/".join([path,who,which])
    print(f"~imag:{one}")
arg = arg=sys.argv[1]
if arg:
    for i in range(int(arg)):
        one()
else:
    one()
