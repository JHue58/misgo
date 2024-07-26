package util

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func ReadYaml(obj any, yamlPath ...string) (err error) {
	var data []byte
	// 读取 YAML 文件
	for _, path := range yamlPath {
		data, err = os.ReadFile(path)
		if err == nil {
			break
		}
	}
	if len(data) <= 0 {
		return fmt.Errorf("尝试读取yaml失败 yaml文件不存在或为空: %v ", yamlPath)
	}

	return yaml.Unmarshal(data, obj)
}
