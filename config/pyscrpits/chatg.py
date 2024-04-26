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
    completion = client.chat.completions.create(
        model="gpt-3.5-turbo",
        messages=[
            {"role": "user",
             "content":"你是一个学者,你要以学者的知识回答我所问的问题"},
            {"role": "user", "content":inp}
        ],
        temperature=0.2,
    )

    print(completion.choices[0].message.content,end="")
import sys

# 接收并打印传递的参数
# -*- coding: utf-8 -*-
sys.stdout.reconfigure(encoding='utf-8')

# 接收并打印传递的参数
for arg in sys.argv[1:]:
    oneround(orig_para(arg))
