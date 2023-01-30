# SimpleTikTok_MSA
go-zero从单体向微服务改进的尝试


## 1. goctl生成
什么都不填，两个都会生成
```shell
useage: ./GoZeroUse create api
        ./GoZeroUse create proto
        ./GoZeroUse create 
```
### 2. 编译
现在编译会输出参数
```shell
useage: ./GoZeroUse create api
        ./GoZeroUse create proto
        ./GoZeroUse create proto -a # 重新编译
        ./GoZeroUse create  # 全部编译
```
### 3. 启动
发现运行失败，尝试使用ps查看进程是否已经启动了
如果有启动的进程，使用```./GoZeroUse.sh kill all```清除
```shell
useage: ./GoZeroUse run api
        ./GoZeroUse run proto
```

### 4. 停止进程
```shell
useage: ./GoZeroUse kill # 显示帮助信息，一定要填写参数
        ./GoZeroUse kill api
        ./GoZeroUse kill etcd
        ./GoZeroUse kill all
```

### 5. 清理日志文件
```shell
useage: ./GoZeroUse clear 
```

# commit类型
用于说明 commit 的类别，只允许使用下面7个标识。
feat：新功能（feature）</br>
fix/to：修补bug </br>
  - fix：产生 diff 并自动修复此问题。适合于一次提交直接修复问题 </br>
  - to：只产生 diff不 自动修复此问题。适合于多次提交。最终修复问题提交时使用 fix </br>
docs：仅仅修改了文档（documentation） </br>
style： 仅仅修改了空格、格式缩进、逗号等等，不改变代码逻辑 </br>
refactor：代码重构，没有加新功能或者修复 bug（即不是新增功能，也不是修改bug的代码变动） </br>
test：增加测试 </br>
chore：改变构建流程、或者增加依赖库、工具等 </br>
revert：回滚到上一个版本 </br>
merge：代码合并 </br>
sync：同步主线或分支的Bug </br>

# 参考资料

https://go-zero.dev/cn/docs/quick-start/monolithic-service
https://go-zero.dev/cn/docs/goctl/goctl/
https://go-zero.dev/cn/docs/goctl/goctl/