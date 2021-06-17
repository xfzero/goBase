##### golang基础学习笔记
参照网易云课堂，Golang小白到大神你戏班免费课程

##### 笔记和coded对应
1. 笔记中的0x对应code中的0xchap
如果笔记中的某一节有多个code文件，可以以0x_0y命名

2. index.md 文件和index文件夹对应

3. 笔记文件和book文件夹对应

4. code中的test文件夹为测试demo

##### github
git@github.com:xfzero/goBase.git

##### 创建项目
1. 安装git

2. 在当前文件夹打开git

3. git init

4. git stauts
发现中文文件夹名字显示异常
git config --global core.quotepath false

5. git add .

git commit -m '项目添加'
提示：
Author identity unknown

*** Please tell me who you are.

Run

  git config --global user.email "you@example.com"
  git config --global user.name "Your Name"

to set your account's default identity.
Omit --global to set the identity only in this repository.

fatal: unable to auto-detect email address (got 'xfwang@DESKTOP-N2VQKM4.(none)')

6. 配置
git config --global user.name xfwang
git config --global user.email 1414572031@qq.com

继续commit:
git commit -m '项目初始化'

7. ssh-keygen -t rsa -C "1414572031@qq.com"
回车-回车 在users/xfwang 下生成 .ssh文件加

8. github上添加key
登录Github,找到右上角的图标，打开点进里面的Settings，再选中里面的SSH and GPG KEYS，点击右上角的New SSH key，然后Title里面随便填，再把刚才id_rsa.pub里面的内容复制到Title下面的Key内容框里面，最后点击Add SSH key，这样就完成了SSH Key的加密。

9. github上创建项目
你可以直接点New repository来创建goBase项目

Initialize this repository with a README 项如果勾选，则生成README

10. 关联 
git remote add origin git@github.com:xfzero/goBase.git

11. 推送到远程can仓库
git push -u origin master

此时要输入yes,不能直接回车