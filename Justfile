# 跨平台 Justfile (Linux / macOS / Windows PowerShell)
# 设置 shell

set shell := ["bash", "-c"]
set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]

# 安装必要工具
init:
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
    go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
    go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
    go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
    go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
    go install github.com/bufbuild/buf/cmd/buf@latest
    go install github.com/google/wire/cmd/wire@latest
    go install github.com/lhlyu/kratos-easy/cmd/ke@latest


# 构建二进制
build:
    @echo "Building project..."
    mkdir -p bin
    go build -ldflags="-s -w" -o ./bin/server ./cmd/...

# 生成 API proto，例如：just api demo [v1]
api name version="v1":
    ke api {{name}} {{version}}

# 生成代码 & tidy
generate:
    @echo "Running go generate..."
    buf dep update
    buf generate
    wire ./cmd/...
    go generate ./...
    go mod tidy

# 更新依赖
update:
    buf dep update
    go get -u ./...
    go mod tidy

# 格式化代码
format:
    buf format -w
    go fmt ./...

# 检查代码
check:
    buf lint
    go vet ./...

# 执行所有生成任务
all: generate check

# 运行项目
run:
    kratos run

# 显示帮助
help:
    @echo ""
    @echo "Usage:"
    @echo "  just <recipe>"
    @echo ""
    @echo "Recipes:"
    just --list
