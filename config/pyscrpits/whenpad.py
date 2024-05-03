"""
<<
name:随机一句
describe:随机一句话
parse:拍一拍

>>
"""
import sys
sys.stdout.reconfigure(encoding='utf-8')
def orig_para(para):
    para= para.replace("₹"," ").replace("ℳ","/n")
    return para
arg=orig_para(sys.argv[1])


import random
import requests as rq

def one_saying():
    def format_hitokoto(data):
        # 提取数据
        hitokoto = data['hitokoto']
        source = data['from']
        from_who = data['from_who']

        # 计算最长长度
        max_length = max(len(hitokoto), len(source), len(from_who) if from_who else 0)

        # 格式化句子
        formatted_hitokoto = hitokoto + '\n'
        if source:
            formatted_hitokoto += ' ' * (max_length - len(source)) + ' & ' + source + '\n'
        if from_who:
            formatted_hitokoto += ' ' * (max_length - len(from_who)) + ' - ' + from_who

        return formatted_hitokoto
    url = "https://v1.hitokoto.cn/"
    g = rq.get(url)
    if g.status_code==200:
        js = g.json()
        # print(js)
        print(format_hitokoto(js),end="")
    else:
        print("嘿嘿",end="")
        # print(response)
    g.close()

def one_saying2():
    def format_data(data):
        formatted_output = ""
        for item in data['data']:
            content = item['content']
            author = item['author']
            if author:  # 如果作者不为空
                formatted_output += f"{content}\n{' ' * (len(content) - len(author))} \t--{author}"
            else:  # 如果作者为空，则只输出句子
                formatted_output += f"{content}"
        return formatted_output
    url="https://www.mxnzp.com/api/daily_word/recommend?count=1&app_id=syoplkiosvdbgolq&app_secret=yYXnwIdg7GwZEX1ycDBDYwcruNHG1kFM"
    g = rq.get(url)
    if g.status_code==200:
        js = g.json()
        print(format_data(js),end="")


def zen_of_git():
    url = "https://api.github.com/zen"
    try:
        g = rq.get(url)
        print(g.text)
    except rq.exceptions.ConnectionError:
        print("嘿嘿",end="")

def ai_English():
    import os
    api_key="sk-3sAIcSk9UWxVcOglj5xTy0Lr8EIrHtUaQ5j7M1c22bkVS3cq"
    if key:=os.environ.get("OPENAI_API_KEY"):
        api_key =key
    from openai import OpenAI
    client = OpenAI(
    # defaults to os.environ.get("OPENAI_API_KEY")
        api_key=api_key,
        base_url="https://api.chatanywhere.tech/v1"
    )


    completion = client.chat.completions.create(
        model="gpt-3.5-turbo",
        messages=[
            {"role": "user",
             "content":"返回一句英文六级的长一点美句,一定要按照格式为:原句\n 翻译,其他什么都不要说"},
        ],
        temperature=1.4,
    )

    print(completion.choices[0].message.content,end="")

ls = [zen_of_git,one_saying,one_saying2,ai_English,ai_English,ai_English]
ls[random.randint(0,len(ls)-1)]()