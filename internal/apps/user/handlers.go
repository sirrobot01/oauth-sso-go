package user

import (
	"github.com/sirrobot01/oauth-sso/internal/common"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("register"))
	if err != nil {
		return
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var data LoginInSchema
	err := common.ReaderToStruct(r, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	common.RenderJSON(w, data, http.StatusOK)

}
