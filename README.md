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

3. `replace github.com/tstgo/conv => /src/test/conv`
replace后的版本号可以省略

4. `go get github/tstgo/calculator@v1.0.0`
在执行上述命令之前，在github上创建了v1.0.0的tag及release，执行后本地依赖版本被更新为v1.0.0

5. `go get github.com/tstgo/calculator@v2.0.0-pre`
在此之前从dev分支直接创建tagv2.0.0-pre并标记为pre-release，执行后报错
> go: github.com/tstgo/calculator@v2.0.0-pre: invalid version: module contains a go.mod file, so module path must match major version ("github.com/tstgo/calculator/v2")

6. `go get github.com/tstgo/calculator@v2.0.0-prev2`
> invalid version: go.mod has post-v2 module path "github.com/tstgo/calculator/v2" at revision v2.0.0-prev2`
原来是module已经变了，所以应该执行`go get github.com/tstgo/calculator/v2@v2.0.0-prev2`

1. 使用`go.work`后，所依赖的模块可以位于磁盘的任意位置，此时`go mod tidy`依然会去仓库拉取go.mod中对应版本的代码，但不会被使用，项目会使用`go.work`中所`use`的代码