package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"troyan/pkg/dtypes"
	"troyan/pkg/generation"
)

func NormalizeHandler(w http.ResponseWriter, r *http.Request) {
	dto := dtypes.NormalizeDTO{}
	input, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Println("Can't read /input body", err)
		return
	}
	err = json.Unmarshal(input, &dto)
	if err != nil {
		log.Println("Can't unmarshall /normalize body", err)
		return
	}
	res := generation.Normalize(dto.ToNormalize)
	outVal := dtypes.OutputCorrDTO{
		OutputSignal: dtypes.OutputSignal{
			Values:     dto.ToNormalize,
			CorrValues: res,
		},
	}
	out, err := json.Marshal(&outVal)
	log.Println(string(out))
	if err != nil {
		log.Println("error marshalling /normalized", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
