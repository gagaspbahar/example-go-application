package service

import (
	"fmt"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/model"
)

// ConfigService interface to abstract field and tree configuration
type ConfigService interface {
	GetFieldDimensions() (model.Field, error)
	GetTrees(NumberOfTrees int) ([]model.Tree, error)
}

type Scanner interface {
	Scan() bool
	Text() string
}

const (
	// FieldWidthMin minimum field Width
	FieldMin      = 1
	FieldMax      = 50000
	TreeHeightMin = 1
	TreeHeightMax = 30
)

// FileConfigService struct to implement ConfigService
type FileConfigService struct {
	Scanner
}

func (f *FileConfigService) GetFieldDimensions() (model.Field, error) {
	if !f.Scanner.Scan() {
		return model.Field{}, fmt.Errorf("failed to read field dimensions")
	}
	line := f.Scanner.Text()

	var field model.Field
	if _, err := fmt.Sscanf(line, "%d %d %d", &field.Width, &field.Length, &field.NumberOfTrees); err != nil {
		return model.Field{}, fmt.Errorf("error reading field dimensions: %v", err)
	}

	if field.Width < FieldMin {
		return model.Field{}, fmt.Errorf("field Width is too low: %d", field.Width)
	}
	if field.Width > FieldMax {
		return model.Field{}, fmt.Errorf("field Width is too high: %d", field.Width)
	}
	if field.Length < FieldMin {
		return model.Field{}, fmt.Errorf("field Length is too low: %d", field.Length)
	}
	if field.Length > FieldMax {
		return model.Field{}, fmt.Errorf("field Length is too high: %d", field.Length)
	}
	if field.NumberOfTrees < FieldMin {
		return model.Field{}, fmt.Errorf("number of trees is too low: %d", field.NumberOfTrees)
	}
	if field.NumberOfTrees > FieldMax {
		return model.Field{}, fmt.Errorf("number of trees is too high: %d", field.NumberOfTrees)
	}

	return field, nil
}

func (f *FileConfigService) GetTrees(NumberOfTrees int) ([]model.Tree, error) {
	var trees []model.Tree
	for i := 0; i < NumberOfTrees; i++ {
		if !f.Scanner.Scan() {
			return nil, fmt.Errorf("failed to read tree data")
		}
		line := f.Scanner.Text()

		var tree model.Tree
		if _, err := fmt.Sscanf(line, "%d %d %d", &tree.X, &tree.Y, &tree.Height); err != nil {
			return nil, fmt.Errorf("error reading tree data: %v", err)
		}
		if tree.Height > TreeHeightMax {
			return nil, fmt.Errorf("tree height is too high: %d", tree.Height)
		}
		if tree.Height < TreeHeightMin {
			return nil, fmt.Errorf("tree height is too low: %d", tree.Height)
		}
		trees = append(trees, tree)
	}
	return trees, nil
}
