import json
import sys

import requests

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


def towhhich(arg):
    arg=arg.replace("₹"," ")
    if arg.endswith("-j"):
        return tranlate(arg[:-2], "auto2ja")
    if arg.endswith("-e"):
        return tranlate(arg[:-2], "auto2en")
    if arg.endswith("-c"):
        return tranlate(arg[:-2], "auto2zh")
    return tranlate(arg[:], "auto2zh")

def orig_para(para):
    para= para.replace("₹"," ").replace("ℳ","/n")
    return para
for arg in sys.argv[1:]:
    print(towhhich(orig_para(arg)),end="")
