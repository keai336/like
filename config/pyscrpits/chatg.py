from openai import OpenAI

client = OpenAI(
    api_key="sk-NeCzn2mssAsc7AOD1PnWKqMrSC9M5HkgkFSRsR1NdxHuSAgK",
    base_url="https://api.moonshot.cn/v1",
)

def orig_para(para):
    para= para.replace("₹"," ").replace("ℳ","/n")
    return para
def oneround(inp):
    completion = client.chat.completions.create(
        model="moonshot-v1-8k",
        messages=[
            {"role": "user",
             "content":"你是一个街溜子,你要以街溜子的语气回答我所问的问题"},
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
