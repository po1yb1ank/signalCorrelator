package dtypes

const (
	Dots       = 150
	WhiteNoise = "WHITE_NOISE"
	Example    = "EXAMPLE"
	Correlated = "CORRELATED"
	Regular    = "REGULAR"
)

type Object struct {
	Config struct {
		SignalType string  `json:"signalType"`
		ObjType    string  `json:"objType"`
		HyperParam float64 `json:"hyperParam"`
	} `json:"config"`
	InputSignal  InputSignal `json:"inputSignal"`
	OutputSignal struct {
		Values     []float64 `json:"values"`
		CorrValues []float64 `json:"corrValues"`
	} `json:"outputSignal"`
}
type Config struct {
	SignalType string `json:"signalType"`
}
type InputDTO struct {
	SignalType string `json:"signalType"`
}
type InputSignal struct {
	Values     []float64 `json:"values"`
	CorrValues []float64 `json:"corrValues"`
}
type OutputSignal struct {
	Values     []float64 `json:"values"`
	CorrValues []float64 `json:"corrValues"`
}
type OutputDTO struct {
	Config      Config      `json:"config"`
	InputSignal InputSignal `json:"inputSignal"`
}
type OutputCorrDTO struct {
	OutputSignal OutputSignal `json:"outputSignal"`
}
type NormalizeDTO struct {
	ToNormalize []float64 `json:"toNormalize"`
}
