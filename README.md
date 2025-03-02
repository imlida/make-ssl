# SSL证书生成工具

这是一个简单易用的命令行工具，用于生成自签名SSL证书。该工具基于OpenSSL，可以快速生成用于开发和测试环境的证书。

## 功能特点

- 一键生成RSA 4096位自签名SSL证书
- 自动设置证书的Common Name (CN)
- 支持配置Subject Alternative Name (SAN)
- **始终将DNS:localhost添加为主题别名**，确保本地开发环境可用
- 自动检测IP地址格式的CN，并添加为IP主题别名
- 默认使用CN值作为证书文件名
- 支持自定义证书目录路径和有效期

## 安装要求

- Go 1.15或更高版本
- OpenSSL命令行工具

## 安装方法

```bash
# 克隆仓库
git clone https://github.com/yourusername/make-ssl.git
cd make-ssl

# 构建
go build -o make-ssl
```

## 使用方法

### 基本用法

```bash
# 使用默认设置生成证书（CN=localhost）
./make-ssl

# 指定CN为IP地址
./make-ssl -cn 192.168.1.9

# 指定额外的主题别名
./make-ssl -cn example.com -alt-names "IP:192.168.1.9"
```

### 命令行参数

| 参数 | 描述 | 默认值 |
|------|------|--------|
| `-cert-path` | 证书存放目录路径 | `./certs` |
| `-cn` | 证书的Common Name (CN) | `localhost` |
| `-days` | 证书有效期(天) | `3650` (约10年) |
| `-alt-names` | 额外的主题别名 | `""` (空) |
| `-cert-file` | 证书文件名 | `[CN].pem` |
| `-key-file` | 私钥文件名 | `[CN].key` |

## 使用示例

### 为本地开发生成证书

```bash
./make-ssl
```

这将在`./certs`目录下生成`localhost.pem`和`localhost.key`文件，有效期为10年。

### 为特定IP地址生成证书

```bash
./make-ssl -cn 192.168.1.9
```

这将生成CN为`192.168.1.9`的证书，并自动添加`IP:192.168.1.9`和`DNS:localhost`作为主题别名。

### 为域名生成证书并添加多个主题别名

```bash
./make-ssl -cn example.com -alt-names "IP:192.168.1.9,IP:10.0.0.1"
```

这将生成CN为`example.com`的证书，并添加指定的IP地址和`DNS:localhost`作为主题别名。

### 自定义证书路径和文件名

```bash
./make-ssl -cn 192.168.1.9 -cert-path "/etc/ssl" -cert-file "mycert.pem" -key-file "mykey.key"
```

## 注意事项

- 此工具生成的是**自签名证书**，浏览器会显示安全警告。若要避免警告，需要将证书添加到系统或浏览器的信任存储中。
- 始终确保您的私钥文件（.key）受到保护，不要将其共享或提交到版本控制系统。
- 本工具适用于开发和测试环境，生产环境建议使用权威证书颁发机构(CA)签发的证书。

## 命令示例

以下是工具实际执行的命令示例：

```
openssl req -x509 -newkey rsa:4096 -keyout ./certs/192.168.1.9.key -out ./certs/192.168.1.9.pem -days 3650 -nodes -subj "/CN=192.168.1.9" -addext "subjectAltName=IP:192.168.1.9,DNS:localhost"
```

## 许可证

[MIT License](LICENSE)