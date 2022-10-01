package staffs

import (
	"net/http"

	"github.com/juandreww/2k22go/18x0_gomongopostgres/config"
)

func List(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	list, err := AllStaffs()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	config.TPL.ExecuteTemplate(w, "staffs.gohtml", list)
}
