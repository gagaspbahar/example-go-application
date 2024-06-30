package main

import (
	"bytes"
	"testing"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/model"
	"go.uber.org/mock/gomock"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock config service
	mockConfigService := NewMockConfigService(ctrl)
	mockConfigService.EXPECT().GetFieldDimensions().Return(model.Field{Width: 5, Length: 1, NumberOfTrees: 3}, nil).Times(1)
	mockConfigService.EXPECT().GetTrees(3).Return([]model.Tree{
		{X: 2, Y: 1, Height: 10},
		{X: 3, Y: 1, Height: 10},
		{X: 4, Y: 1, Height: 10},
	}, nil).Times(1)

	// Create mock writer
	mockOutput := &bytes.Buffer{}
	mockErrorOutput := &bytes.Buffer{}

	err := run(mockConfigService, mockOutput, mockErrorOutput)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	expectedOutput := "62\n"
	output := mockOutput.String()
	if output != expectedOutput {
		t.Errorf("expected %s, got %s", expectedOutput, output)
	}
}

func TestCalculateTotalDistance_NoTrees(t *testing.T) {
	field := model.Field{Width: 5, Length: 5, NumberOfTrees: 0}
	trees := make([][]int, 6)
	for i := range trees {
		trees[i] = make([]int, 6)
	}

	expected := 242 // 24 * 10 + 2 * 1
	result := calculateTotalDistance(field, trees)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestCalculateTotalDistance_TreesAtEdges(t *testing.T) {
	field := model.Field{Width: 5, Length: 5, NumberOfTrees: 2}
	trees := make([][]int, 6)
	for i := range trees {
		trees[i] = make([]int, 6)
	}
	trees[1][1] = 2
	trees[5][5] = 3

	expected := 252 // Calculated expected result manually
	result := calculateTotalDistance(field, trees)
	if result != expected {
		t.Errorf("expected %d, got %d", expected, result)
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"positive", 5, 5},
		{"negative", -5, 5},
		{"zero", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := abs(tt.input)
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}
