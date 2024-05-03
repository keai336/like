"""
<<
name:音乐下载
describe:输入音乐名字,下载音乐
parse:/music
>>
"""
import sys
sys.stdout.reconfigure(encoding='utf-8')
def orig_para(para):
    para= para.replace("₹"," ").replace("ℳ","/n")
    return para
from spotdl import Spotdl
arg = sys.argv[1]
arg = orig_para(arg)
spotdl = Spotdl(client_id='9fcc5c8a71324aaf9e50d87f35c91ce2', client_secret='11c37a0018d14e44ad7579e202d98ddb')
if songs:=spotdl.search([arg]):
    song,path = spotdl.download(songs[0])
    print(f"file:{path}")
