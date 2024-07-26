package clipboard

import (
	"fmt"
	"github.com/jhue/misgo/internal/util"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"strconv"
)

func (b *ClipBoard) SaveFile(rootPath string, p []byte) error {

	if b.Type != FileType {
		return fmt.Errorf("clipboard type is not file")
	}
	if b.Content == "" {
		return fmt.Errorf("clipboard content(filename) is empty")
	}
	if b.Hash == "" {
		b.Hash = util.NewHashBuilder(p).MD5().String()
	}
	fileName := b.Content
	dirPath := filepath.Join(rootPath, strconv.Itoa(int(b.UserID)), b.Hash)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	filePath := filepath.Join(dirPath, fileName)
	return os.WriteFile(filePath, p, os.ModePerm)
}

func (b *ClipBoard) DeleteFile(rootPath string) error {
	if b.Type != FileType {
		return fmt.Errorf("clipboard type is not file")
	}

	if b.Hash == "" {
		return fmt.Errorf("clipboard hash is empty")
	}

	dirPath := filepath.Join(rootPath, strconv.Itoa(int(b.UserID)), b.Hash)

	return os.RemoveAll(dirPath)
}

func (b *ClipBoard) FilePath(rootPath string) (string, error) {
	if b.Type != FileType {
		return "", fmt.Errorf("clipboard type is not file")
	}
	if b.Content == "" {
		return "", fmt.Errorf("clipboard content(filename) is empty")
	}
	if b.Hash == "" {
		return "", fmt.Errorf("clipboard hash is empty")
	}
	fileName := b.Content
	dirPath := filepath.Join(rootPath, strconv.Itoa(int(b.UserID)), b.Hash, fileName)
	return dirPath, nil
}

func (b *ClipBoard) Store(d *gorm.DB, rootPath string, maxRecords int64, p []byte) (deleteCount int, boardCount int, err error) {
	err = d.Transaction(func(tx *gorm.DB) error {

		if b.Type == FileType {
			if err := b.SaveFile(rootPath, p); err != nil {
				return err
			}
		}

		// 插入新记录
		if err := tx.Create(b).Error; err != nil {
			return err
		}

		// 获取当前用户的记录数量
		var count int64
		if err := tx.Model(b).Where("user_id = ?", b.UserID).Count(&count).Error; err != nil {
			return err
		}

		// 如果记录数超过最大值，删除最旧的记录
		if count > maxRecords {
			var recordsToDelete []ClipBoard
			// 查询需要删除的记录
			if err := tx.Where("user_id = ?", b.UserID).
				Order("time ASC").
				Limit(int(count - maxRecords)).
				Find(&recordsToDelete).Error; err != nil {
				return err
			}

			// 删除记录
			for _, record := range recordsToDelete {
				if err := tx.Delete(&record).Error; err != nil {
					return err
				}
				if record.Type == FileType {
					if err := record.DeleteFile(rootPath); err != nil {
						return err
					}
				}
				deleteCount++
			}

		}
		boardCount = int(count) - deleteCount
		return nil

	})
	return
}
