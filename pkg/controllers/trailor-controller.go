package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/HammadShakeel463/go-crud/pkg/models"
	"github.com/HammadShakeel463/go-crud/pkg/utils"
	"github.com/gorilla/mux"
)

var NewTailor models.Tailor

func GetTailor(w http.ResponseWriter, r *http.Request) {
	newTailor := models.GetTailor()
	res, _ := json.Marshal(&newTailor)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTailorById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tailorid := vars["tailorid"]
	Id, err := strconv.ParseInt(tailorid, 0, 0)
	if err != nil {
		fmt.Print("Error in parsing int ")
	}
	tailorDetail, _ := models.GetTailorById(Id)
	res, _ := json.Marshal(tailorDetail)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateTailor(w http.ResponseWriter, r *http.Request) {
	tailor := &models.Tailor{}
	utils.ParseBody(r, tailor)
	b := tailor.CreateTailor()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTailor(w http.ResponseWriter, x *http.Request) {
	vars := mux.Vars(x)
	Id := vars["tailoeid"]
	ID, err := strconv.ParseInt(Id, 0, 0)
	if err != nil {
		fmt.Print("Error in parsing int while delete")
	}
	tailor := models.DeleteTailor(ID)
	res, _ := json.Marshal(tailor)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTailor(w http.ResponseWriter, r *http.Request) {
	var updatetailor = &models.Tailor{}
	utils.ParseBody(r, updatetailor)
	vars := mux.Vars(r)
	Id := vars["tailorid"]
	id, err := strconv.ParseInt(Id, 0, 0)
	tailorDeatail, db := models.GetTailorById(id)
	if err != nil {
		fmt.Print("Update : error while parsing int")
	}
	if updatetailor.FirstName != "" {
		tailorDeatail.FirstName = updatetailor.FirstName
	}
	if updatetailor.LastName != "" {
		tailorDeatail.LastName = updatetailor.LastName
	}
	if updatetailor.Cnic != "" {
		tailorDeatail.Cnic = updatetailor.Cnic
	}
	if updatetailor.PhoneNo != "" {
		tailorDeatail.PhoneNo = updatetailor.PhoneNo
	}
	if updatetailor.Email != "" {
		tailorDeatail.Email = updatetailor.Email
	}
	db.Save(tailorDeatail)
	res, _ := json.Marshal(tailorDeatail)
	w.Header().Set("Conent-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
