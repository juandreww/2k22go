package staffs

import "net/http"

func List(w http.Response, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	list, err := AllStaffs()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	//
}
