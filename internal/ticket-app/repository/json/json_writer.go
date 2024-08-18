package repository

import (
	"encoding/json"
	"gowit-task/internal/ticket-app/model"
	"os"
)

type jsonWriter struct {
	path string
}

func NewJsonWriter(path string) *jsonWriter {
	return &jsonWriter{path: path}
}

func (w *jsonWriter) Send(b []byte) error {
	return os.WriteFile(w.path, b, 0644)
}

func (w *jsonWriter) Write(fiestaInfo model.FiestaInfo) error {
	b, err := os.ReadFile(w.path)
	if err != nil {
		return err
	}
	var data JsonData
	err = json.Unmarshal(b, &data)
	if err != nil {
		return err
	}

	mp := convertDataToFiestaMap(data)
	mp[fiestaInfo.Id] = FiestaData(fiestaInfo)

	jsonData := convertFistaMapToData(mp)

	b, err = json.Marshal(jsonData)
	if err != nil {
		return err
	}

	return w.write(b)
}

func (w *jsonWriter) write(b []byte) error {
	return os.WriteFile(w.path, b, 0644)
}

func convertDataToFiestaMap(d JsonData) map[int]FiestaData {
	mp := map[int]FiestaData{}
	for _, info := range d.Data {
		mp[info.Id] = info
	}
	return mp
}

func convertFistaMapToData(mp map[int]FiestaData) JsonData {
	var jData JsonData
	jData.Data = make([]FiestaData, 0, len(mp))

	for _, info := range mp {
		jData.Data = append(jData.Data, info)
	}

	return jData
}
