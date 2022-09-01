# go-version

[https://go.dev/ref/mod#workspaces](https://go.dev/ref/mod#workspaces)指出go模块使用版本号的规范

为了享受规范给开发带来的便利，因此尝试实践：
- [ ] 什么情况下会出现pseudo-version
- [ ] 在go.mod中如何规范使用版本号来解决多人协同开发可能出现的问题


## 试验结果

前提条件：**拉取时main分支只有add、dev分支有add、divide**
1. `go get -u github.com/tstgo/calculator@dev`
执行完上述命令后，本地代码可找到divide函数，证明的确拉取了dev分支，但是版本号为**v0.0.0-20220901075114-61173e2eb89e**，说明是pseudo-version

2. `go get -u github.com/tstgo/calculator`
执行完上述命令后，本地代码找不到divide函数，证明拉取的是默认分支main，版本号为**v0.0.0-20220901075041-137910462925**，可以看到时间戳较dev旧，最后面是commitid的前几位