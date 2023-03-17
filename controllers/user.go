package cont

import (
	"encoding/json"
	"fmt"
	db "gym-api/Database"
	mod "gym-api/models"
	"net/http"
	"time"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var user mod.User
	json.NewDecoder(r.Body).Decode(&user)
	db.DB.Create(&user)
	json.NewEncoder(w).Encode(&user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("get users called")
	var output mod.Display
	query := "SELECT users.user_id,users.user_name,users.gender, payments.amount, payments.payment_type, payments.payment_id, subscriptions.subs_name, subscriptions.start_date, subscriptions.deleted_at,subscriptions.end_date, subscriptions.duration, subscriptions.emp_id FROM users JOIN payments ON users.user_id = payments.user_id JOIN subscriptions ON payments.payment_id = subscriptions.payment_id; "

	db.DB.Raw(query).Scan(&output)

	json.NewEncoder(w).Encode(&output)

}

func UserAttendence(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	now := time.Now()

	var temp mod.UAttendence
	temp.User_Id = id
	temp.Present = "Present"
	temp.Date = now.Format("02 Jan 2006")
	db.DB.Create(&temp)
}
