package main

import (
	"context"
	"fmt"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func usage() {
	fmt.Fprintf(os.Stderr, "用法:\n")
	fmt.Fprintf(os.Stderr, "  上传: %s upload <本地图片路径>\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  删除: %s delete <对象名>\n", os.Args[0])
	os.Exit(2)
}

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("获取用户目录失败: %v", err)
	}
	envPath := filepath.Join(homeDir, ".config", "minio-upload", ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("加载 %s 失败: %v", envPath, err)
	}

	if len(os.Args) < 3 {
		usage()
	}

	action := os.Args[1]
	target := os.Args[2]

	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL, _ := strconv.ParseBool(os.Getenv("MINIO_USE_SSL"))
	bucketName := os.Getenv("MINIO_BUCKET")
	expireHoursStr := os.Getenv("MINIO_PRESIGNED_EXPIRE_HOURS")
	expireHours, _ := strconv.Atoi(expireHoursStr)
	if expireHours <= 0 {
		expireHours = 168 // 默认7天
	}

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		die("创建客户端失败: %v", err)
	}

	ctx := context.Background()

	switch action {
	case "upload":
		doUpload(ctx, minioClient, bucketName, target, time.Duration(expireHours)*time.Hour)
	case "delete":
		doDelete(ctx, minioClient, bucketName, target)
	default:
		usage()
	}
}

func doUpload(ctx context.Context, client *minio.Client, bucket, localFile string, expire time.Duration) {
	objectName := filepath.Base(localFile)

	// 确保桶存在
	exists, err := client.BucketExists(ctx, bucket)
	if err != nil {
		die("检查桶失败: %v", err)
	}
	if !exists {
		err = client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
		if err != nil {
			die("创建桶失败: %v", err)
		}
	}

	file, err := os.Open(localFile)
	if err != nil {
		die("打开文件失败: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		die("获取文件信息失败: %v", err)
	}

	contentType := mime.TypeByExtension(filepath.Ext(localFile))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err = client.PutObject(ctx, bucket, objectName, file, fileInfo.Size(), minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		die("上传失败: %v", err)
	}

	scheme := "http"
	if os.Getenv("MINIO_USE_SSL") == "true" {
		scheme = "https"
	}
	fmt.Printf("%s://%s/%s/%s", scheme, os.Getenv("MINIO_ENDPOINT"), bucket, objectName)
}

func doDelete(ctx context.Context, client *minio.Client, bucket, objectName string) {
	err := client.RemoveObject(ctx, bucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		die("删除失败: %v", err)
	}
	fmt.Print("删除成功")
}

func die(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}
