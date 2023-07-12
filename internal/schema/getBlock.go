package schema

import "fmt"

func GetBlockFromName(name string, files []*FileData) (*BlockData, error) {
	for _, file := range files {
		for _, block := range file.Blocks {
			if block.Name == name {
				return &block, nil
			}
		}
	}
	return nil, fmt.Errorf("No block found with name %s", name)
}
