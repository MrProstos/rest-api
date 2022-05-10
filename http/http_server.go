package http

import (
	"encoding/json"
	"net/http"

	"github.com/MrProstos/rest-api/db"
)

var oper db.Operator

func getOper(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(oper)
}

func main() {

}
