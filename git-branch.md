# git branch 

## 主分支

- master分支

    通常，master分支只能从其它分支合并，不能在master分支直接修改。master分支上存放的是随时可供在生产环境中部署的代码（Production Ready state）。当开发活动到一定阶段，产生一份新的可供部署的代码时，master分支上的代码会被更新。同时，每一次更新，最好添加对应的版本号标签（TAG）。

    所有在Master分支上的Commit应该打Tag。

- develop分支

    develop分支是保持当前开发最新成果的分支，一般会在此分支上进行晚间构建(Nightly Build)并执行自动化测试。 

    develop分支产生于master分支, 并长期存在。

    当一个版本功能开发完毕且通过测试功能稳定时，就会合并到master分支上，并打好带有相应版本号的tag。

    develop分支一般命名为develop

    develop分支是主开发分支，包含所有要发布到下一个Release的代码，主要合并其它分支，比如Feature分支。


## 辅组分支

- feature分支

    **feature分支使用规范**：

    可以从develop分支发起feature分支。

    代码必须合并回develop分支。

    feature分支的命名可以使用除master，develop，release-*，hotfix-*之外的任何名称。

    feature分支（topic分支）通常在开发一项新的软件功能的时候使用，分支上的代码变更最终合并回develop分支或者干脆被抛弃掉（例如实验性且效果不好的代码变更）。

    一般而言，feature分支代码可以保存在开发者自己的代码库中而不强制提交到主代码库里。

    Feature分支开发完成后，必须合并回Develop分支，合并完分支后一般会删Feature分支，但也可以保留。

- release分支

    **release分支使用规范：**

    可以从develop分支派生；

    必须合并回develop分支和master分支；

    分支命名惯例：release-*；

    release分支是为发布新的产品版本而设计的。在release分支上的代码允许做测试、bug修改、准备发布版本所需的各项说明信息（版本号、发布时间、编译时间等）。通过在release分支上进行发布相关工作可以让develop分支空闲出来以接受新的feature分支上的代码提交，进入新的软件开发迭代周期。

    当develop分支上的代码已经包含了所有即将发布的版本中所计划包含的软件功能，并且已通过所有测试时，可以考虑准备创建release分支。而所有在当前即将发布的版本外的业务需求一定要确保不能混到release分支内（避免由此引入一些不可控的系统缺陷）。

    成功的派生release分支并被赋予版本号后，develop分支就可以为下一个版本服务。版本号的命名可以依据项目定义的版本号命名规则进行。

    发布Release分支时，合并Release到Master和Develop， 同时在Master分支上打个Tag记住Release版本号，然后就可以删除Release分支。
   
- git 分支结构图

![git分支结构图](https://github.com/zhangymPerson/learning-notes/blob/master/Picture/git-version.png)