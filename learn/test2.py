messages=[]
"""
<<
name:aibot
describe:问答ai
parse:/ai
>>
"""
from openai import OpenAI

# client = OpenAI(
#     api_key="sk-NeCzn2mssAsc7AOD1PnWKqMrSC9M5HkgkFSRsR1NdxHuSAgK",
#     base_url="https://api.moonshot.cn/v1",
# )

import os
api_key="sk-3sAIcSk9UWxVcOglj5xTy0Lr8EIrHtUaQ5j7M1c22bkVS3cq"
if key:=os.environ.get("OPENAI_API_KEY"):
    api_key =key
client = OpenAI(
    # defaults to os.environ.get("OPENAI_API_KEY")
    api_key=api_key,
    base_url="https://api.chatanywhere.tech/v1"
)
def orig_para(para):
    para= para.replace("₹"," ").replace("ℳ","/n")
    return para


def oneround(inp):
    messages.append({"role":"user",
                     "content":inp})
    completion = client.chat.completions.create(
        model="gpt-4",
        messages=messages,
        temperature=1.0,
    )
    # print(completion)
    rep = completion.choices[0].message.content
    print(rep)
    messages.append({"role":"system",
                     "content":rep})

import sys

# 接收并打印传递的参数
# -*- coding: utf-8 -*-
sys.stdout.reconfigure(encoding='utf-8')

# 接收并打印传递的参数
oneround("我们玩一个游戏,你给我一个六级英语单词,我回答中文翻译,如果对了,给我加一分,如果错了,游戏结束,结算我的积分")
while True:
    oneround(input("~"))
