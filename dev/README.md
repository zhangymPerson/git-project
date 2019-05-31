# git使用

- 分支


```sh
# 查看所有的分支和远程分支
$ git branch -a

# 创建一个新的分支
$ git branch myNewBranch

# 删除一个分支
$ git branch -d myBranch

# 重命名分支
# git branch -m <旧名称> <新名称>
$ git branch -m myBranchName myNewBranchName

# 编辑分支的介绍
$ git branch myBranchName --edit-description

```

- 检出

将当前工作空间更新到索引所标识的或者某一特定的工作空间
```sh
# 检出一个版本库，默认将更新到master分支
$ git checkout
# 检出到一个特定的分支
$ git checkout branchName
# 新建一个分支，并且切换过去，相当于"git branch <名字>; git checkout <名字>"
$ git checkout -b newBranch
```

- clone

    这个命令就是将一个版本库拷贝到另一个目录中，同时也将 分支都拷贝到新的版本库中。这样就可以在新的版本库中提交到远程分支
```sh
# clone learnxinyminutes-docs
$ git clone https://github.com/adambard/learnxinyminutes-docs.git
```
- commit

    将当前索引的更改保存为一个新的提交，这个提交包括用户做出的更改与信息
```sh
# 提交时附带提交信息
$ git commit -m "Added multiplyNumbers() function to HelloWorld.c"
```