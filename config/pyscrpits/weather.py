"""
<<
name:地名天气
describe:给定一个地名返回天气信息但是不太准确,通过ai根据地名获得的经纬,然后调用彩云api 格式 /weather+地名
parse:/weather
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
def oneroundai(message):
    completion = client.chat.completions.create(
        model="gpt-3.5-turbo",
        messages=message,
        temperature=0.2,
    )

    return completion.choices[0].message.content

def location_context(name):
    context1 = [
        {"role": "user",
         "content": "以后给定你一个地名,你返回其经纬度,格式:经度,纬度,其他无用的话不要回答,连个空格都不要多"},
        {"role": "user", "content": name}
    ]
    return context1

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
    locat = oneroundai(location_context(inp)).replace(" ","")
    rand = "hourly?hourlysteps=1"
    url ='/'.join([u,token,locat,rand])
    # print(url)
    r = rq.get(url)
    if r.status_code == 200:
        message = r.json()["result"]["hourly"]
        message["location_name"] = inp
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


# 接收并打印传递的参数
weatherhourly(arg)