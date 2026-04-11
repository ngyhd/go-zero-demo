package upload

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

const (
	// MaxFileSize 最大文件大小 10MB
	MaxFileSize = 10 * 1024 * 1024
	// AllowedExtensions 允许的文件扩展名
	AllowedExtensions = ".jpg,.jpeg,.png,.gif,.webp,.pdf"
)

// UploadResult 上传结果
type UploadResult struct {
	URL      string `json:"url"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

// LocalStorage 本地存储
type LocalStorage struct {
	BasePath string
	BaseURL  string
}

// NewLocalStorage 创建本地存储
func NewLocalStorage(basePath, baseURL string) *LocalStorage {
	// 确保目录存在
	os.MkdirAll(basePath, 0755)
	return &LocalStorage{
		BasePath: basePath,
		BaseURL:  baseURL,
	}
}

// Upload 上传文件
func (s *LocalStorage) Upload(ctx context.Context, file *multipart.FileHeader) (*UploadResult, error) {
	// 检查文件大小
	if file.Size > MaxFileSize {
		return nil, fmt.Errorf("文件大小超过限制: %d bytes", file.Size)
	}

	// 检查扩展名
	ext := filepath.Ext(file.Filename)
	if !isAllowedExtension(ext) {
		return nil, fmt.Errorf("不支持的文件类型: %s", ext)
	}

	// 打开源文件进行 MIME 和 Magic 检查
	src, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("打开文件失败: %v", err)
	}
	defer src.Close()

	// 读取文件头部用于 MIME 类型检测
	buffer := make([]byte, 512)
	n, err := src.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("读取文件失败: %v", err)
	}

	// MIME 类型检测
	mimeType := http.DetectContentType(buffer[:n])
	if !isAllowedMIMEType(mimeType) {
		return nil, fmt.Errorf("不支持的 MIME 类型: %s", mimeType)
	}

	// Magic Number 检测 - 验证文件真实类型
	if !isAllowedMagicNumber(buffer[:n], ext) {
		return nil, fmt.Errorf("文件内容与扩展名不匹配")
	}

	// 重置文件指针到开头
	if _, err := src.Seek(0, io.SeekStart); err != nil {
		return nil, fmt.Errorf("重置文件指针失败: %v", err)
	}

	// 生成唯一文件名
	filename := fmt.Sprintf("%s_%d%s", uuid.New().String()[:8], time.Now().UnixNano(), ext)
	filepath := filepath.Join(s.BasePath, filename)

	// 创建目标文件
	dst, err := os.Create(filepath)
	if err != nil {
		return nil, fmt.Errorf("创建文件失败: %v", err)
	}
	defer dst.Close()

	// 复制内容
	if _, err := io.Copy(dst, src); err != nil {
		return nil, fmt.Errorf("保存文件失败: %v", err)
	}

	return &UploadResult{
		URL:      s.BaseURL + "/" + filename,
		Filename: filename,
		Size:     file.Size,
	}, nil
}

// isAllowedExtension 检查扩展名是否允许
func isAllowedExtension(ext string) bool {
	ext = filepath.Ext(ext)
	for _, allowed := range splitExtensions(AllowedExtensions) {
		if ext == allowed {
			return true
		}
	}
	return false
}

// AllowedMIMETypes 允许的 MIME 类型
var AllowedMIMETypes = map[string]bool{
	"image/jpeg":              true,
	"image/png":                true,
	"image/gif":                true,
	"image/webp":               true,
	"application/pdf":          true,
	"application/octet-stream": true, // 某些系统可能返回这个
}

// isAllowedMIMEType 检查 MIME 类型是否允许
func isAllowedMIMEType(mimeType string) bool {
	return AllowedMIMETypes[mimeType]
}

// MagicNumberMap 定义扩展名对应的 Magic Number 签名
var MagicNumberMap = map[string][][]byte{
	".jpg":  {{0xFF, 0xD8, 0xFF}},                       // JPEG
	".jpeg": {{0xFF, 0xD8, 0xFF}},                       // JPEG
	".png":  {{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}}, // PNG
	".gif":  {{0x47, 0x49, 0x46, 0x38, 0x37, 0x61}, {0x47, 0x49, 0x46, 0x38, 0x39, 0x61}}, // GIF87a, GIF89a
	".pdf":  {{0x25, 0x50, 0x44, 0x46}},               // %PDF
}

// isAllowedMagicNumber 检查文件 Magic Number 是否与扩展名匹配
func isAllowedMagicNumber(data []byte, ext string) bool {
	signatures, ok := MagicNumberMap[ext]
	if !ok {
		// 如果没有定义 Magic Number，检查 MIME 类型作为后备
		mimeType := http.DetectContentType(data)
		return isAllowedMIMEType(mimeType)
	}

	for _, sig := range signatures {
		if len(data) < len(sig) {
			continue
		}
		// 检查前几个字节是否匹配
		match := true
		for i := 0; i < len(sig); i++ {
			if data[i] != sig[i] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

func splitExtensions(s string) []string {
	var result []string
	for _, ext := range splitString(s, ",") {
		result = append(result, ext)
	}
	return result
}

func splitString(s, sep string) []string {
	var result []string
	start := 0
	for i := 0; i < len(s); i++ {
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			start = i + len(sep)
		}
	}
	result = append(result, s[start:])
	return result
}

// ParseFormFile 解析 multipart form 的文件
func ParseFormFile(r *http.Request, fieldName string) (*multipart.FileHeader, error) {
	if err := r.ParseMultipartForm(MaxFileSize); err != nil {
		return nil, fmt.Errorf("解析表单失败: %v", err)
	}

	file, header, err := r.FormFile(fieldName)
	if err != nil {
		return nil, fmt.Errorf("获取文件失败: %v", err)
	}
	file.Close()

	return header, nil
}
