package service

import (
	"testing"

	"github.com/SawitProRecruitment/JuniorBackendEngineering/model"
	gomock "go.uber.org/mock/gomock"
)

func TestGetFieldDimensions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock scanner
	mockScanner := NewMockScanner(ctrl)
	mockScanner.EXPECT().Scan().Return(true).Times(1)
	mockScanner.EXPECT().Text().Return("5 5 3").Times(1)

	configService := &FileConfigService{Scanner: mockScanner}
	field, err := configService.GetFieldDimensions()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expected := model.Field{Width: 5, Length: 5, NumberOfTrees: 3}
	if field != expected {
		t.Errorf("expected %v, got %v", expected, field)
	}
}

func TestGetFieldDimensionsInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock scanner
	mockScanner := NewMockScanner(ctrl)
	mockScanner.EXPECT().Scan().Return(true).Times(1)
	mockScanner.EXPECT().Text().Return("invalid input").Times(1)

	configService := &FileConfigService{Scanner: mockScanner}
	_, err := configService.GetFieldDimensions()
	if err == nil {
		t.Fatal("expected error, got none")
	}
}

func TestGetTrees(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock scanner
	mockScanner := NewMockScanner(ctrl)
	mockScanner.EXPECT().Scan().Return(true).Times(4)
	mockScanner.EXPECT().Text().Return("5 5 3").Times(1)
	mockScanner.EXPECT().Text().Return("1 1 2").Times(1)
	mockScanner.EXPECT().Text().Return("2 2 3").Times(1)
	mockScanner.EXPECT().Text().Return("3 3 4").Times(1)

	configService := &FileConfigService{Scanner: mockScanner}
	field, err := configService.GetFieldDimensions()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	trees, err := configService.GetTrees(field.NumberOfTrees)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	expectedTrees := []model.Tree{
		{X: 1, Y: 1, Height: 2},
		{X: 2, Y: 2, Height: 3},
		{X: 3, Y: 3, Height: 4},
	}
	if len(trees) != len(expectedTrees) {
		t.Fatalf("expected %d trees, got %d", len(expectedTrees), len(trees))
	}
	for i, tree := range trees {
		if tree != expectedTrees[i] {
			t.Errorf("expected tree %v, got %v", expectedTrees[i], tree)
		}
	}
}

func TestGetTreesInvalid(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock scanner
	mockScanner := NewMockScanner(ctrl)
	mockScanner.EXPECT().Scan().Return(true).Times(1)
	mockScanner.EXPECT().Text().Return("5 5 1").Times(1)
	mockScanner.EXPECT().Scan().Return(true).Times(1)
	mockScanner.EXPECT().Text().Return("invalid tree data").Times(1)

	configService := &FileConfigService{Scanner: mockScanner}
	field, err := configService.GetFieldDimensions()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = configService.GetTrees(field.NumberOfTrees)
	if err == nil {
		t.Fatal("expected error, got none")
	}
}

func TestInvalidDimensions(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock scanner
	mockScanner := NewMockScanner(ctrl)
	mockScanner.EXPECT().Scan().Return(true).Times(1)
	mockScanner.EXPECT().Text().Return("0 0 0").Times(1)

	configService := &FileConfigService{Scanner: mockScanner}
	_, err := configService.GetFieldDimensions()
	if err == nil {
		t.Fatal("expected error, got none")
	}
}

func TestInvalidTreeHeight(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock scanner
	mockScanner := NewMockScanner(ctrl)
	mockScanner.EXPECT().Scan().Return(true).Times(1)
	mockScanner.EXPECT().Text().Return("5 5 1").Times(1)
	mockScanner.EXPECT().Scan().Return(true).Times(1)
	mockScanner.EXPECT().Text().Return("1 1 31").Times(1)

	configService := &FileConfigService{Scanner: mockScanner}
	field, err := configService.GetFieldDimensions()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = configService.GetTrees(field.NumberOfTrees)
	if err == nil {
		t.Fatal("expected error, got none")
	}

}

func TestInvalidFieldInput(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock scanner
	mockScanner := NewMockScanner(ctrl)
	mockScanner.EXPECT().Scan().Return(false).Times(1)

	configService := &FileConfigService{Scanner: mockScanner}
	_, err := configService.GetFieldDimensions()
	if err == nil {
		t.Fatal("expected error, got none")
	}
}

func TestWidthTooLow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock scanner
	mockScanner := NewMockScanner(ctrl)
	mockScanner.EXPECT().Scan().Return(true).Times(1)
	mockScanner.EXPECT().Text().Return("0 5 1").Times(1)

	configService := &FileConfigService{Scanner: mockScanner}
	_, err := configService.GetFieldDimensions()
	if err == nil {
		t.Fatal("expected error, got none")
	}
}

func TestWidthTooHigh(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock scanner
	mockScanner := NewMockScanner(ctrl)
	mockScanner.EXPECT().Scan().Return(true).Times(1)
	mockScanner.EXPECT().Text().Return("50001 5 1").Times(1)

	configService := &FileConfigService{Scanner: mockScanner}
	_, err := configService.GetFieldDimensions()
	if err == nil {
		t.Fatal("expected error, got none")
	}
}

func TestLengthTooLow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock scanner
	mockScanner := NewMockScanner(ctrl)
	mockScanner.EXPECT().Scan().Return(true).Times(1)
	mockScanner.EXPECT().Text().Return("5 0 1").Times(1)

	configService := &FileConfigService{Scanner: mockScanner}
	_, err := configService.GetFieldDimensions()
	if err == nil {
		t.Fatal("expected error, got none")
	}
}

func TestLengthTooHigh(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock scanner
	mockScanner := NewMockScanner(ctrl)
	mockScanner.EXPECT().Scan().Return(true).Times(1)
	mockScanner.EXPECT().Text().Return("5 50001 1").Times(1)

	configService := &FileConfigService{Scanner: mockScanner}
	_, err := configService.GetFieldDimensions()
	if err == nil {
		t.Fatal("expected error, got none")
	}
}

func TestNumberOfTreesTooLow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock scanner
	mockScanner := NewMockScanner(ctrl)
	mockScanner.EXPECT().Scan().Return(true).Times(1)
	mockScanner.EXPECT().Text().Return("5 5 0").Times(1)

	configService := &FileConfigService{Scanner: mockScanner}
	_, err := configService.GetFieldDimensions()
	if err == nil {
		t.Fatal("expected error, got none")
	}
}

func TestNumberOfTreesTooHigh(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create mock scanner
	mockScanner := NewMockScanner(ctrl)
	mockScanner.EXPECT().Scan().Return(true).Times(1)
	mockScanner.EXPECT().Text().Return("5 5 50001").Times(1)

	configService := &FileConfigService{Scanner: mockScanner}
	_, err := configService.GetFieldDimensions()
	if err == nil {
		t.Fatal("expected error, got none")
	}
}
