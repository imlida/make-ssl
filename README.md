# Make-SSL

一个用Go语言编写的命令行工具，用于生成自签名SSL证书。

## 功能特点

- 检查证书文件是否存在，避免重复生成
- 支持指定自定义域名和IP地址
- 支持自定义证书有效期
- 默认使用[域名].key和[域名].pem作为文件名
- 自动创建证书目录
- 自动设置适当的文件权限
- certs目录已加入.gitignore，不会被提交到Git仓库

## 依赖项

- OpenSSL（需要在系统上预先安装）
- Go 1.13+（如果要从源码编译）

## 安装

### 从源码编译

```bash
git clone https://github.com/imlida/make-ssl.git
cd make-ssl
go build -o make-ssl
```

### 使用预编译二进制文件

从 [Releases](https://github.com/imlida/make-ssl/releases) 页面下载适用于您操作系统的二进制文件。

## 使用方法

### 基本用法

```bash
# 使用默认参数，在当前目录的certs子目录下生成证书
# 证书将保存为 certs/localhost.pem 和 certs/localhost.key
./make-ssl

# 指定Common Name
# 证书将保存为 certs/example.com.pem 和 certs/example.com.key
./make-ssl -cn example.com

# 指定IP地址和域名作为SAN（主题备用名称）
./make-ssl -cn example.com -alt-names "DNS:example.com,IP:192.168.1.8,DNS:localhost"
```

### 可用选项

```
  -alt-names string
        额外的主题别名 (格式: DNS:example.com,IP:192.168.1.1)
  -cert-file string
        证书文件名 (默认为 [域名].pem)
  -cert-path string
        证书存放目录路径 (默认 "./certs")
  -cn string
        证书的Common Name (CN) (默认 "localhost")
  -days int
        证书有效期(天) (默认 3650)
  -key-file string
        私钥文件名 (默认为 [域名].key)
```

## 示例

### 为本地开发生成证书

```bash
./make-ssl -cn localhost -alt-names "DNS:localhost,IP:127.0.0.1"
```

### 为Web服务器生成证书

```bash
./make-ssl -cn example.com -alt-names "DNS:example.com,DNS:www.example.com" -cert-path /etc/ssl/certs
```

### 生成长期有效的证书

```bash
./make-ssl -cn example.com -days 7300
```

### 指定自定义文件名

```bash
# 如果您不想使用域名作为文件名
./make-ssl -cn example.com -cert-file custom.pem -key-file custom.key
```

## 许可证

MIT