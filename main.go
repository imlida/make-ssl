package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// 定义命令行参数
	certPath := flag.String("cert-path", "./certs", "证书存放目录路径")
	cn := flag.String("cn", "localhost", "证书的Common Name (CN)")
	days := flag.Int("days", 3650, "证书有效期(天)")
	altNames := flag.String("alt-names", "", "额外的主题别名 (格式: DNS:example.com,IP:192.168.1.1)")

	// 添加文件名选项，但默认为空，表示使用域名作为文件名
	certFile := flag.String("cert-file", "", "证书文件名 (默认为 [域名].pem)")
	keyFile := flag.String("key-file", "", "私钥文件名 (默认为 [域名].key)")

	flag.Parse()

	// 设置文件名，如果用户未指定，则使用域名
	actualCertFile := *certFile
	actualKeyFile := *keyFile

	if actualCertFile == "" {
		actualCertFile = *cn + ".pem"
	}

	if actualKeyFile == "" {
		actualKeyFile = *cn + ".key"
	}

	// 确保证书目录存在
	if err := os.MkdirAll(*certPath, 0755); err != nil {
		log.Fatalf("创建证书目录失败: %v", err)
	}

	certFilePath := filepath.Join(*certPath, actualCertFile)
	keyFilePath := filepath.Join(*certPath, actualKeyFile)

	// 检查证书文件是否已存在
	certExists := fileExists(certFilePath)
	keyExists := fileExists(keyFilePath)

	if certExists && keyExists {
		fmt.Println("证书已存在，跳过生成步骤")
		return
	}

	fmt.Println("证书文件不存在，正在生成自签名证书...")

	// 构建 OpenSSL 命令
	args := []string{
		"req", "-x509",
		"-newkey", "rsa:4096",
		"-keyout", keyFilePath,
		"-out", certFilePath,
		"-days", fmt.Sprintf("%d", *days),
		"-nodes",
		"-subj", fmt.Sprintf("/CN=%s", *cn),
	}

	// 处理主题别名
	if *altNames != "" {
		// 用户提供了自定义的主题别名
		args = append(args, "-addext", fmt.Sprintf("subjectAltName=%s", *altNames))
	} else {
		// 如果没有指定alt-names，则使用CN作为DNS名称
		args = append(args, "-addext", fmt.Sprintf("subjectAltName=DNS:%s", *cn))
	}

	// 执行OpenSSL命令
	cmd := exec.Command("openssl", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("执行命令:", cmd.String())

	if err := cmd.Run(); err != nil {
		log.Fatalf("生成证书失败: %v", err)
	}

	// 设置文件权限
	if err := os.Chmod(keyFilePath, 0600); err != nil {
		log.Printf("警告: 无法设置密钥文件权限: %v", err)
	}

	if err := os.Chmod(certFilePath, 0644); err != nil {
		log.Printf("警告: 无法设置证书文件权限: %v", err)
	}

	fmt.Println("自签名证书已生成")
	fmt.Printf("证书路径: %s\n", certFilePath)
	fmt.Printf("密钥路径: %s\n", keyFilePath)
}

// 检查文件是否存在
func fileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
