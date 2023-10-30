package regression

import (
	"github.com/sjwhitworth/golearn/base"
	"github.com/sjwhitworth/golearn/linear_models"
)

const (
	FILE_PATH = "data/boston.csv"
	SPLIT     = 0.7
)

type Model struct {
	Data         base.FixedDataGrid
	TrainingData base.FixedDataGrid
	TestingData  base.FixedDataGrid
	LRModel      *linear_models.LinearRegression
}

func New(columnsToOmit map[string]bool) *Model {
	data, err := base.ParseCSVToInstances(FILE_PATH, true)
	if err != nil {
		panic(err)
	}

	//remove attributes
	for _, attr := range data.AllAttributes() {
		if columnsToOmit[attr.GetName()] {
			err := data.RemoveClassAttribute(attr)
			if err != nil {
				panic(err)
			}
		}
	}

	trainingData, testingData := base.InstancesTrainTestSplit(data, SPLIT)

	model := &Model{
		Data:         data,
		TrainingData: trainingData,
		TestingData:  testingData,
		LRModel:      linear_models.NewLinearRegression(),
	}

	return model
}

func (m *Model) TrainAndPredict() float64 {

	//train
	err := m.LRModel.Fit(m.TrainingData)
	if err != nil {
		panic(err)
	}

	// test
	predictions, err := m.LRModel.Predict(m.TestingData)
	if err != nil {
		panic(err)
	}

	return meanSquaredError(m.TestingData, predictions)
}

func meanSquaredError(actualData, predictions base.FixedDataGrid) float64 {
	_, numRows := actualData.Size()

	// get mv as attribute
	actualAttribute := actualData.AllClassAttributes()[0]
	predAttribute := predictions.AllClassAttributes()[0]

	// get mv as attribute spec
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

	return sum / float64(numRows)
}
