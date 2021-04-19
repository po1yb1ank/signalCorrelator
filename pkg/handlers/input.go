package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"troyan/pkg/dtypes"
	"troyan/pkg/generation"
)

func InputHandler(w http.ResponseWriter, r *http.Request) {
	dto := dtypes.InputDTO{}
	input, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("Can't read /input body", err)
		return
	}
	err = json.Unmarshal(input, &dto)
	if err != nil {
		log.Println("Can't unmarshall /input body", err)
		return
	}
	switch dto.SignalType {
	case dtypes.Example:
		signal := generation.GetExampleSignal()
		rxx := generation.Normalize(signal)
		out := dtypes.OutputDTO{
			Config: dtypes.Config{
				SignalType: dtypes.Example,
			},
			InputSignal: dtypes.InputSignal{
				Values:     signal,
				CorrValues: rxx,
			},
		}
		res, err := json.Marshal(&out)
		if err != nil {
			log.Println("Can't marshall output", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	case dtypes.Correlated:
		signal := generation.GetCorrelatedSignal()
		rxx := generation.Normalize(signal)
		out := dtypes.OutputDTO{
			Config: dtypes.Config{
				SignalType: dtypes.Correlated,
			},
			InputSignal: dtypes.InputSignal{
				Values:     signal,
				CorrValues: rxx,
			},
		}
		res, err := json.Marshal(&out)
		if err != nil {
			log.Println("Can't marshall output", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	case dtypes.WhiteNoise:
		signal := generation.GetGaussianSignal()
		rxx := generation.Normalize(signal)
		out := dtypes.OutputDTO{
			Config: dtypes.Config{
				SignalType: dtypes.WhiteNoise,
			},
			InputSignal: dtypes.InputSignal{
				Values:     signal,
				CorrValues: rxx,
			},
		}
		res, err := json.Marshal(&out)
		if err != nil {
			log.Println("Can't marshall output", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	case dtypes.Regular:
		signal := generation.GetRegularSignal()
		rxx := generation.Normalize(signal)
		out := dtypes.OutputDTO{
			Config: dtypes.Config{
				SignalType: dtypes.Regular,
			},
			InputSignal: dtypes.InputSignal{
				Values:     signal,
				CorrValues: rxx,
			},
		}
		res, err := json.Marshal(&out)
		if err != nil {
			log.Println("Can't marshall output", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	}

}
