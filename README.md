# gostudy
### 一、项目拉取、提交
#### 1、添加SSH and GPG keys
```
ssh-keygen -t rsa -C "邮箱"
# 在C:\Users\***\.ssh中找到id_rsa.pub, 用notepad++打开，复制里面的内容
# settings添加SSH and GPG keys
```
#### 2、$GOPATH/src/gostudy目录下：
```
git init
git remote add [起的远程分支的名字] [github上的仓库的ssh地址]
# 把远程origin分支上的内容同步到本地master上
git pull master master --allow-unrelated-histories

```
#### 3、本地与仓库关联后
```
# 本地的内容已经利用git add 和git commit -m 提交到本地仓库。
# 使用git push 来把本地内容同步到远程main， 
--set-upstream orgin master 是设置master的上游即远程为origin。
```
### 
