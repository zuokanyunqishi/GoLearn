package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
	minio "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("加载 .env 文件失败: %v", err)
	}

	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL, _ := strconv.ParseBool(os.Getenv("MINIO_USE_SSL"))
	bucketName := os.Getenv("MINIO_BUCKET")
	localFile := os.Getenv("MINIO_LOCAL_FILE")
	objectName := filepath.Base(localFile)

	// 1. 初始化 MinIO 客户端
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("创建客户端失败: %v", err)
	}

	// 2. 检查桶是否存在，不存在则自动创建（可选）
	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		log.Fatalf("检查桶是否存在时出错: %v", err)
	}
	if !exists {
		log.Printf("桶 %s 不存在，正在创建...", bucketName)
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("创建桶失败: %v", err)
		}
		log.Printf("桶 %s 创建成功", bucketName)
	}

	// 3. 打开本地图片文件
	file, err := os.Open(localFile)
	if err != nil {
		log.Fatalf("打开文件失败: %v", err)
	}
	defer file.Close()

	// 获取文件信息（用于上传时传递文件大小）
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalf("获取文件信息失败: %v", err)
	}
	fileSize := fileInfo.Size()

	// 4. 执行上传
	uploadInfo, err := minioClient.PutObject(ctx, bucketName, objectName, file, fileSize, minio.PutObjectOptions{
		ContentType: "image/jpeg", // 如果上传其他类型请修改，也可以让MinIO自动检测
	})
	if err != nil {
		log.Fatalf("上传失败: %v", err)
	}

	fmt.Printf("✅ 上传成功！\n")
	fmt.Printf("   桶: %s\n", bucketName)
	fmt.Printf("   对象名: %s\n", objectName)
	fmt.Printf("   ETag: %s\n", uploadInfo.ETag)
	fmt.Printf("   大小: %d bytes\n", uploadInfo.Size)
}
