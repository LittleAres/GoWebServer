package views

import (
	"github.com/gorilla/mux"
)

var RouterList = mux.NewRouter()

func CollectRouter(){
	RouterList.HandleFunc("/", RenderView)
	RouterList.HandleFunc("/register_handler", RegisterView).Methods("POST")
	RouterList.HandleFunc("/login_handler", LoginView).Methods("POST")
	RouterList.HandleFunc("/logout_handler", LogoutView).Methods("POST")
}
