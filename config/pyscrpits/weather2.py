"""
<<
name:地址天气
describe:发地址给天气
parse:位置

>>
"""
import sys
sys.stdout.reconfigure(encoding='utf-8')
def orig_para(para):
    para= para.replace("₹"," ").replace("ℳ","/n")
    return para
arg=orig_para(sys.argv[1])


import json
import requests as rq
from openai import OpenAI

client = OpenAI(
    # defaults to os.environ.get("OPENAI_API_KEY")
    api_key="sk-9tqFMrPxR4eQVBZfnv2PF1QwJWpKLaaVrbfw4FfkWEdIkl9e",
    base_url="https://api.chatanywhere.tech/v1"
)

def oneroundai(message):
    completion = client.chat.completions.create(
        model="gpt-3.5-turbo",
        messages=message,
        temperature=0.2,
    )

    return completion.choices[0].message.content


def discribe_context(json):
    context1 = [
        {"role": "user",
         "content": "以后给定你一个json数据,根据数据生成一段天气状况描述,你要像天气播报员一样,简洁明细,给出一些建议,不要说无关的话"},
        {"role": "user", "content": json}
    ]
    return context1

def weatherhourly(inp):
    u = 'https://api.caiyunapp.com/v2.6'
    token = 'yctDWm6mKtnILZI4'
    ls = inp.split(",")
    locat = "{},{}".format(ls[2],ls[1])
    name = ls[0]
    rand = "hourly?hourlysteps=1"
    url ='/'.join([u,token,locat,rand])
#     print(url)
    r = rq.get(url)
    if r.status_code == 200:
        message = r.json()["result"]["hourly"]
        message["location_name"] = name
        message = json.dumps(message,ensure_ascii=False)
        # print(message)
        describe = oneroundai(discribe_context(message))
        # temp = message["temperature"][0]["value"]
        # app_temp = message["apparent_temperature"][0]["value"]
        # wind_v=message["wind"][0]["speed"]
        # humidity = message["humidity"][0]["value"]
        # print(temp,app_temp,wind_v)
        print(describe)
    else:
        print(r.status_code)
    r.close()


weatherhourly(arg)