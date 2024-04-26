import re
import time
from datetime import date
import sys
sys.stdout.reconfigure(encoding='utf-8')
from lunarcalendar import Lunar


class Notify:
    def __init__(self, start=None, end=None, connect=None):
        self.start = start
        self.end = end
        self.connect = connect

    def nots(self):
        if self.start:
            if self.connect:
                if_print(self.connect, self.start)
            else:
                print(self.start)

    def note(self):
        if self.end:
            if self.connect:
                if_print(self.connect, self.end)
            else:
                print(self.end)


def if_print(connect, msg):
    if connect is None:
        print(msg)
    else:
        connect.sendall(msg.encode("utf-8"))
        time.sleep(0.02)


def if_input(connect, tips):
    if connect is None:
        i = input(tips)
    else:
        connect.sendall("input {}".format(tips).encode("utf-8"))
        i = connect.recv(1024).decode("utf-8")
        if i == "mm":  # 约定
            i = ""
        print(i)
    return i


def get_items(di):
    """
    :return:  条子 dics
    """
    keys = list(di.keys())
    itls = ['{}.{}'.format(k + 1, keys[k]) for k in range(len(keys))]
    dics = {str(k + 1): keys[k] for k in range(len(keys))}
    # print(itls)
    itme = get_chs(itls)
    return itme, dics


def dedupe(items):
    """
    列表有序驱虫
    :param items:
    :return:
    """
    seen = set()
    for item in items:
        if item not in seen:
            yield item
            seen.add(item)  ##序列值可哈希


def get_chs(iter,sep = '   ',n=5):
    """
    分割输出效果
    :return:
    """
    chs = ''
    nn = 1
    for i in iter:
        if nn % n == 0:
            chs += str(i) + '\n'
        else:
            chs += str(i) + sep
        nn += 1
    # print('当前选择{}\n'.format(chs))
    return chs


def chose_items(di, mum=-1, notify=Notify(), outbreak="q", connect=None):
    """

    :param di: 字典类型  要选的  后续支持其他可迭代对象
    :param mum: 要选数目的最大值
    :param notify: 选择前的引导类
    :param outbreak: 强制结束选择关键词
    :param connect: 套接字否
    :return:
    """

    def add_a():
        """
        正选应用
        :return:
        """
        # nonlocal mum
        # fw = []
        # for i in kls:
        #     if i in dics.values():
        #         if mum:
        #             bag[i] = di.pop(i)
        #             mum += -1
        #         else:
        #             fw.append(i)
        #     else:
        #         fw.append(i)
        # if fw:
        #     fw_dic = {i: i for i in fw}
        #     if_print(connect, '这些元素为无效选择：\n{}\n'.format(get_items(fw_dic)[0]))


# version 2
        nonlocal now_n
        fw = []
        for i in kls:
            if i in dics.values():
                bag[i] = di.pop(i)
                now_n += 1
            else:
                fw.append(i)
        if fw:
            fw_dic = {i: i for i in fw}
            if_print(connect, '这些元素为无效选择：\n{}\n'.format(get_items(fw_dic)[0]))

    def add_b():
        """
        反选应用
        :return:
        """
        # fw = []
        # nonlocal mum
        # for i in dics.values():
        #     if i not in kls:
        #         if mum:
        #             bag[i] = di.pop(i)
        #             mum += -1
        #         else:
        #             fw.append(i)
        # if fw:
        #     fw_dic = {i: i for i in fw}
        #     if_print(connect, '这些元素为无效选择：\n{}\n'.format(get_items(fw_dic)[0]))

        # version2
        nonlocal now_n
        for i in dics.values():
            if i not in kls:
                bag[i] = di.pop(i)
                now_n += 1


    def num_cho(contend):
        """
        matched_num  匹配内容
        基于数字的选择
        支持 1，2，3  | 1 2 3 | 1,2,3 | 1-3
        :return: 键列表
        """

        # 正则转换 1-9 to 1,2,3,4,5,6,7,8,9  辅助函数
        def range_n(x):
            a = int(x.group(1))
            b = int(x.group(2))
            rangee = [str(i) for i in range(a, b + 1)]
            fi = ','.join(rangee)
            return fi

        nonlocal mum, di, bag  # 声明非本地变量
        fw = []
        single_str = re.sub(pattern='(\d+)-(\d+)', repl=range_n, string=contend)  # 格式化字符串
        ns = re.split('[, ，]', single_str)
        ns = list(dedupe(ns))  # 每次选择后的选项列表
        kls = []
        for i in ns:
            if i in dics.keys():
                kls.append(dics[i])
            else:
                fw.append(i)
        if fw:
            fw_dic = {i: i for i in fw}
            if_print(connect, '这些元素为无效输入：\n{}\n'.format(get_items(fw_dic)[0]))
        return kls



    def re_choice(content):
        """
        正则处理
        :param content: 匹配内容
        基于增则表达选择 ，选择匹配到的条目，
        :return:  匹配到的键列表
        """
        nonlocal mum, di, bag
        r = re.compile(content, re.IGNORECASE)
        kls = [i for i in dics.values() if re.search(r, i)]  # 搜索结果 为键列表
        return kls
    if not isinstance(di,dict):
        try:
            di = {i:i for i in di}
        except TypeError:
            pass
    notify.nots()
    bag = {}
    now_n = 0
    r = re.compile('([+\-]?) ?([rR]?) ?(.+)', re.S)  # 宽匹配 检查是否符合输入格式
    while now_n!= mum:
        if now_n > mum and mum!=-1:
            di = {k:v for k,v in bag.items()}
            bag.clear()
            now_n = 0
        elif not di:
            break

        item, dics = get_items(di)
        if_print(connect, item)
        q = if_input(connect, "请选择")
        if q:
            if q == outbreak:
                break
            o = re.match(r, q)
            if o:
                zf = o.group(1)
                mode = o.group(2)
                content = o.group(3)
                if mode:
                    kls = re_choice(content)
                else:
                    kls = num_cho(content)
                if zf != "-":
                    add_a()
                else:
                    add_b()
            else:
                if_print(connect, print('输入错误'))
        else:
            if bag:
                kkk = list(bag.keys())
                bag = {i: bag[i] for i in kkk}
                if_print(connect, '当前有效选择: {}\n'.format(get_chs(bag.keys())))
                break
            else:
                if_print(connect, '你还没选')

        if_print(connect, '当前选择: {}\n'.format(get_chs(bag.keys())))

    notify.note()
    return bag




class Birthd(object):
    @property
    def birthday(self):
        r_birth = re.compile(r"([阴阳]?)(\d{4}[ -]?\d{2}[ -]?\d{2})")
        mc = r_birth.match(self.btd)
        sd = date.fromisoformat(mc.group(2))
        if mc.group(1) == "阴":
            sd = Lunar(sd.year, sd.month, sd.day).to_date()
        return sd

    def __init__(self, birth):
        ls = re.split(r"[ ,，]", birth)
        # print(ls)
        self.name = ls[0]
        self.btd = ls[1]
        self.cstm = ls[2]

    def _next_brith(self):
        if self.cstm == "阴":
            lu = Lunar.from_date(self.birthday)
            m = lu.month
            d = lu.day
            lu = Lunar(n.year, m, d)
            so = lu.to_date()
            if (so - n).days < 0:
                so = Lunar(n.year + 1, m, d).to_date()
        else:
            so = date(n.year, self.birthday.month, self.birthday.day)
            if (so - n).days < 0:
                so = date(n.year + 1, self.birthday.month, self.birthday.day)
        return so


def next_bird(Birthd):
    next_b = f"{Birthd.name}的下一个生日：{Birthd._next_brith()}"
    deadline = (Birthd._next_brith() - date.today()).days
    return next_b + "   " f"倒计时: {deadline}天"



def load(path):
    try:
        with open(path, encoding="utf-8") as f:
            f = f.readlines()
    except FileNotFoundError:
        raise FileNotFoundError("为创建储存文件，或指定错误路径")
    f = (i.rstrip() for i in f if not i.startswith('#'))
    item = (Birthd(i) for i in f)
    dic = {i.name: i for i in item}
    return dic

path = "a.txt"
n = date.today()
dic = load(path)
if dic:
    csd = chose_items(dic) 
    if csd:
        for k, v in csd.items():
            print(next_bird(v))
    else:
        print("无有效选择")
else:
    raise ValueError("文件内无有效数据，请检查文件")
    
