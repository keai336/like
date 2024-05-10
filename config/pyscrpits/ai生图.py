"""
<<
name:ai生图
describe:ai生图,/imag+描述 [-b(大图) -3(第三代模型)]
parse:/imag
>>
"""
import requests
import json

import sys
sys.stdout.reconfigure(encoding='utf-8')
def orig_para(para):
    para= para.replace("₹"," ").replace("ℳ","/n")
    return para
arg=orig_para(sys.argv[1])

# arg = input("")
model = "dall-e-2"
size = "256x256"
if arg.endswith("-3"):
   arg = arg[:-2]
   model = "dall-e-3"
   size = "1024x1024"
elif arg.endswith("-b"):
   arg = arg[:-2]
   size = "1024x1024"

url = "https://api.chatanywhere.tech/v1/images/generations"
payload = json.dumps({
   "prompt": arg,
   "n": 1,
   "model": model,
   "size": size
})
headers = {
   'Authorization': 'Bearer sk-3sAIcSk9UWxVcOglj5xTy0Lr8EIrHtUaQ5j7M1c22bkVS3cq',
   'User-Agent': 'Apifox/1.0.0 (https://apifox.com)',
   'Content-Type': 'application/json'
}

response = requests.request("POST", url, headers=headers, data=payload)
rsp = response.json()
url = rsp["data"][0]["url"]
id = rsp["created"]
path = f"../config/资源/图/{id}.png"
def download_image(url, save_path):
   try:
      # 发送GET请求获取图片数据
      response = requests.get(url)
      # 确保请求成功
      if response.status_code == 200:
         # 保存图片到指定路径
         with open(save_path, 'wb') as f:
            f.write(response.content)
      else:
         print("下载失败，HTTP 状态码:", response.status_code)
   except Exception as e:
      print("下载失败:", str(e))

download_image(url,save_path=path)
print(f"~imag:{path}")