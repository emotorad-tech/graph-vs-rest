package handlers

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"rest/middlewares"
	"time"
)

type response struct {
	Id    string `json:"id"`
	Phone string `json:"phone"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

func GetData(w http.ResponseWriter, r *http.Request) {
	ra := rand.New(rand.NewSource(time.Now().UnixNano()))
	query := fmt.Sprintf(`SELECT id, phone, name, age FROM dummy WHERE age = %d`, ra.Intn(10))
	db, _ := middlewares.GetDatabaseFromContext(r.Context())
	row, _ := db.Query(r.Context(), query)
	var data response
	var result []response
	for row.Next() {
		if row.Scan(&data.Id, &data.Phone, &data.Name, &data.Age) != nil {
			log.Panic("something went wrong")
		}
		result = append(result, data)
	}
	writeJson(w, http.StatusOK, result)
}
