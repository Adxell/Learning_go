package handler 
import "Learning_go/claseCrud/model"
type person struct {
	storage Storage
}

func newPerson(storage Storage) person {
	return person{storage}
}

func (p *person) create(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		w.Header().Set("Content-Type", "appplication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":"error", "message": "Method no default"}`))
		return
	}

	data:=model.Person{}
	err:=json.newDecoder(r.Body).Decode(&data)
	if err !=nil{
		w.Header().Set("Content-Type", "appplication/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"message_type":"error", "message": "Data  does not correct"}`))
		return
	}

	err = p.storage.Create(&data)
	if err !=nil{
		w.Header().Set("Content-Type", "appplication/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message_type":"error", "message": "Problen with server"}`))
		return
	}

	w.Header().Set("Content-Type", "appplication/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message_type":"successful", "message": "Person Created"}`))
}