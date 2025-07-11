// Blueprint: Auto-generated by HTTP Plugin
package http

import (
	"github.com/gorilla/mux"
	"github.com/blueprint-uservices/tutorial/examples/helloworld/workflow/servicea"
	"context"
	"encoding/json"
	"net/http"
)

type ServiceA_HTTPServerHandler struct {
	Service servicea.ServiceA
	Address string
}

func New_ServiceA_HTTPServerHandler(ctx context.Context, service servicea.ServiceA, serverAddress string) (*ServiceA_HTTPServerHandler, error) {
	handler := &ServiceA_HTTPServerHandler{}
	handler.Service = service
	handler.Address = serverAddress
	return handler, nil
}

// Blueprint: Run is called automatically in a separate goroutine by runtime/plugins/golang/di.go
func (handler *ServiceA_HTTPServerHandler) Run(ctx context.Context) error {
	router := mux.NewRouter()
	// Add paths for the mux router
	
	router.Path("/Hello").HandlerFunc(handler.Hello)
	
	router.Path("/World").HandlerFunc(handler.World)
	
	srv := &http.Server {
		Addr: handler.Address,
		Handler: router,
	}

	go func() {
		select {
		case <-ctx.Done():
			srv.Shutdown(ctx)
		}
	}()

	return srv.ListenAndServe()
}


func (handler *ServiceA_HTTPServerHandler) Hello(w http.ResponseWriter, r *http.Request) {
	var err error
	defer r.Body.Close()
	
	ctx := context.Background()
	err = handler.Service.Hello(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	response := struct {
		
	}{}
	
	json.NewEncoder(w).Encode(response)
}

func (handler *ServiceA_HTTPServerHandler) World(w http.ResponseWriter, r *http.Request) {
	var err error
	defer r.Body.Close()
	
	ctx := context.Background()
	err = handler.Service.World(ctx)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	response := struct {
		
	}{}
	
	json.NewEncoder(w).Encode(response)
}

