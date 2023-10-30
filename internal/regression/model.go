package regression

import (
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/linear_models"
)

const (
	FILE_PATH = "data/boston.csv"
	SPLIT     = 0.7
	MV_COL    = 13
)

type Model struct {
	Data         base.FixedDataGrid
	Columns      map[string]bool
	TrainingData base.FixedDataGrid
	TestingData  base.FixedDataGrid
	LRModel      *linear_models.LinearRegression
}

func New(columns []string) *Model {
	data, err := base.ParseCSVToInstances(FILE_PATH, true)
	if err != nil {
		panic(err)
	}

	columnsSet := make(map[string]bool)
	for _, col := range columns {
		columnsSet[col] = true
	}

	trainingData, testingData := base.InstancesTrainTestSplit(data, SPLIT)

	model := &Model{
		Data:         data,
		Columns:      columnsSet,
		TrainingData: trainingData,
		TestingData:  testingData,
		LRModel:      linear_models.NewLinearRegression(),
	}

	return model
}

func (m *Model) TrainAndPredict() float64 {
	err := m.LRModel.Fit(m.TrainingData)
	if err != nil {
		panic(err)
	}

	predictions, err := m.LRModel.Predict(m.TestingData)
	if err != nil {
		panic(err)
	}

	mse := meanSquaredError(m.TestingData, predictions)
	return mse
}

func meanSquaredError(actualData, predictions base.FixedDataGrid) float64 {
	_, numRows := actualData.Size()

	actualAttribute := actualData.AllAttributes()[0]
	predAttribute := predictions.AllAttributes()[0]

	actualSpec, _ := actualData.GetAttribute(actualAttribute)
	predSpec, _ := predictions.GetAttribute(predAttribute)

	var sum float64

	for i := 0; i < numRows; i++ {
		actualValBytes := actualData.Get(actualSpec, i)
		actualVal := base.UnpackBytesToFloat(actualValBytes)

		predValBytes := predictions.Get(predSpec, i)
		predVal := base.UnpackBytesToFloat(predValBytes)

		diff := actualVal - predVal
		sum += diff * diff
	}

	mse := sum / float64(numRows)
	return mse
}
