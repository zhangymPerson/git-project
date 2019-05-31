- 创建dev-3分支 合并到dev-1 分支 并且删除掉dev-3分支


```sh
#创建并使用分支
$ git checkout -b dev-3

# git branch -a 

# 创建相关内容并且 add 和 commit

# 切换回要合并入的分支
git checkout dev-1
# 合并dev-3的内容到dev-1中
git merge dev-3

# 删除dev-3分支
git branch -d dev-3
