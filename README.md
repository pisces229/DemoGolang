## [Golang](https://marketplace.visualstudio.com/items?itemName=golang.Go)

### [設定 Go 開發的Visual Studio Code](https://docs.microsoft.com/zh-tw/azure/developer/go/configure-visual-studio-code)

### Get Start

`go mod init [package]`

`go mod init demo.golang`

`go run .`

`go run . first "second"`

`go build .`

`gofmt -w .`

`go run -race .`

### [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)

`go test -v`

`go test -v -run <method>`

### 性能測試(Benchmark)

`go test -v -bench=. -benchmem <name>.go`

`go test -v -bench=<method> -benchmem <name>.go`

### Module

`go mod init [module]`

`go get [library]`

`go mod tidy`

`go list -m all`

`go mod download`

`go build -o .`

### go.mod

module
定義模組路徑

go
定義go語言 version

require
指定依賴的套件，預設是最新版，可以指定版本號

exclude
排除該套件和其版本

replace
使用不同的套件版本並替換原有的套件版本

註解
// 單行註解
/* 多行註解 */
indirect 代表被間接導入的依賴包

### pprof

[runtime](https://pkg.go.dev/runtime/pprof)

[http](https://pkg.go.dev/net/http/pprof)

### Package

`go mod tidy`

`go get <package>`

[Go Module](https://proxy.golang.org/)

[uuid](https://github.com/google/uuid)

`go get github.com/google/uuid`

[test](https://github.com/stretchr/testify/assert)

`go get github.com/stretchr/testify/assert`

[pprof]([github.com/gin-contrib/pprof](https://github.com/gin-contrib/pprof))

`go get github.com/gin-contrib/pprof`

[viper](https://pkg.go.dev/github.com/spf13/viper)

`go get github.com/spf13/viper`

[log](https://github.com/sirupsen/logrus)

`go get github.com/sirupsen/logrus`

[gorm](https://gorm.io/)

`go get -u gorm.io/gorm`

`go get gorm.io/driver/sqlserver`

`go get gorm.io/driver/postgres`

[gin](https://github.com/gin-gonic/gin/blob/master/README.md)

`go get github.com/gin-contrib/cors`

`go get github.com/gin-contrib/sessions`

[jwt](https://github.com/golang-jwt/jwt)

`go get -u github.com/golang-jwt/jwt`

[Gin with Logger](https://ithelp.ithome.com.tw/articles/10280560)

[excelize](https://github.com/qax-os/excelize)
