# git reset 版本回退测试

##  git reset 的三个参数

### soft
- 仅仅移动本地库HEAD指针
### mixed
- 移动本地库HEAD指针
- 重置暂存区
### hard
- 移动本地库HEAD指针
- 重置暂存区
- 重置工作区

### 执行命令

```shell script
#切换git分支进行测试
git checkout -b reset
# 随便修改文件内容 多次执行 记录 commit 日志
git commit ./
# 执行回退
# git reset --hard 
# reset --hard 会在重置 HEAD 和branch的同时，重置stage区和工作目录里的内容。
# 当你在 reset 后面加了 --hard 参数时，
# 你的stage区和工作目录里的内容会被完全重置为和HEAD的新位置相同的内容。
# 回退到指定版本，此后所有的commit和修改全部清除
git reset --hard e1ce2bc6d279596bb06690ef895be2b40d70d53f
```
