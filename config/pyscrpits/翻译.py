"""
<<
name:翻译
describe:"支持中英日"
parse:/translate
>>
"""
import json
import sys
import re
import requests

default_dic = {"e":"zh",
               "j":"zh",
               "c":"en"}
sys.stdout.reconfigure(encoding='utf-8')

def tranlate(source, direction):
    url = "http://api.interpreter.caiyunai.com/v1/translator"

    # WARNING, this token is a test token for new developers,
    # and it should be replaced by your token
    token = "pg6h27odnfyz4f8gx8pl"

    payload = {
        "source": source,
        "trans_type": direction,
        "request_id": "demo",
        "detect": True,
    }

    headers = {
        "content-type": "application/json",
        "x-authorization": "token " + token,
    }

    response = requests.request("POST", url, data=json.dumps(payload), headers=headers)

    return json.loads(response.text)["target"]



def orig_para(para):
    para= para.replace("₹"," ").replace("ℳ","/n")
    return para

def analyze_language(text):
    # 定义正则表达式来匹配中文、英文和日文字符
    chinese_pattern = re.compile(r'[\u4e00-\u9fff]')
    english_pattern = re.compile(r'[a-zA-Z]')
    japanese_pattern = re.compile(r'[\u3040-\u309f\u30a0-\u30ff\u31f0-\u31ff\u4e00-\u9fff]')

    # 统计各语言字符的数量
    chinese_count = len(re.findall(chinese_pattern, text))
    english_count = len(re.findall(english_pattern, text))
    japanese_count = len(re.findall(japanese_pattern, text))

    # 计算各语言字符的总数
    total_count = chinese_count + english_count + japanese_count

    # 计算各语言字符的比例
    chinese_percentage = (chinese_count / total_count) * 100 if total_count > 0 else 0
    english_percentage = (english_count / total_count) * 100 if total_count > 0 else 0
    japanese_percentage = (japanese_count / total_count) * 100 if total_count > 0 else 0
    dic_1 = {
        "c": chinese_percentage,
        "e": english_percentage,
        "j": japanese_percentage
    }
    max_value = max(dic_1,key=lambda x: dic_1[x])
    return max_value


def towhhich(arg):
    if arg.endswith("-j"):
        return tranlate(arg[:-2], "auto2ja")
    if arg.endswith("-e"):
        return tranlate(arg[:-2], "auto2en")
    if arg.endswith("-c"):
        return tranlate(arg[:-2], "auto2zh")
    return tranlate(arg[:], "auto2{}".format(default_dic[analyze_language(arg[:])]))

for arg in sys.argv[1:]:
    print(towhhich(orig_para(arg)),end="")
