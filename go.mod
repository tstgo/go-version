module github.com/tstgo/go-version

go 1.19

require (
	github.com/tstgo/calculator/v2 v2.0.0-prev2
	github.com/tstgo/conv v0.0.0-20220901085829-d6fee8f9ee1d
)

// 可以不使用版本号来replace
replace github.com/tstgo/conv => /src/test/conv
