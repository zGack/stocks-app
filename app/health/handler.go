package health

import (
	"net/http"

	"github.com/zgack/stocks/pkg/router"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
    data := map[string]string{
		"status": "ok",
	}
    // This is a simple health check handler that responds with a 200 OK status.
    // It can be used to check if the service is running and healthy.
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
    router.RespondWithJSON(w, http.StatusOK, data)
}


