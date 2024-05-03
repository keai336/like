"""
<<
name:emoai
describe:分析情绪的ai
parse:/emoa
>>
"""
import sys
sys.stdout.reconfigure(encoding='utf-8')
def orig_para(para):
    para= para.replace("₹"," ").replace("ℳ","/n")
    return para
from openai import OpenAI
import os
api_key="sk-3sAIcSk9UWxVcOglj5xTy0Lr8EIrHtUaQ5j7M1c22bkVS3cq"
if key:=os.environ.get("OPENAI_API_KEY"):
    api_key =key
# client = OpenAI(
#     api_key="sk-NeCzn2mssAsc7AOD1PnWKqMrSC9M5HkgkFSRsR1NdxHuSAgK",
#     base_url="https://api.moonshot.cn/v1",
# )
client = OpenAI(
    # defaults to os.environ.get("OPENAI_API_KEY")
    api_key=api_key,
    base_url="https://api.chatanywhere.tech/v1"
)
def oneround(inp):
    completion = client.chat.completions.create(
        model="gpt-3.5-turbo",
        messages=[
            {"role": "system",
             "content": "以后我发你一句话,你只用回复你能感受到的情绪及其强度,情绪名字典[好、乐、哀、怒、惧、恶、惊)],强度评级,满分1000,你只需返回 {情绪名:分数}, 其他什么不要多说 ,m3bro,一个返回例子{'好': 0,'乐': 4,'哀': 0,'怒': 0,'惧': 0,'恶': 0,'惊': 0}"},
            {"role": "user", "content":inp}
        ],
        temperature=0.2,
    )

    print(completion.choices[0].message.content,end="")

# 接收并打印传递的参数
arg=sys.argv[1]
oneround(orig_para(arg))
