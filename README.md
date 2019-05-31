# git-project
git-各种命令测试
# 主要测试创建分支 合并分支 创建版本号

- 合并分支

```sh

# 切换到要合并入的分支
$ git checkout branchName
# 将其他分支合并到当前分支
$ git merge branchName

# 在合并时创建一个新的合并后的提交
$ git merge --no-ff branchName

```

log
显示这个版本库的所有提交

# 显示所有提交
$ git log

# 显示某几条提交信息
$ git log -n 10

# 仅显示合并提交
$ git log --merges
merge
合并就是将外部的提交合并到自己的分支中
