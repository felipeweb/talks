func main() {
	// Reorder middlewares
	middlewares.MiddlewareStack = []negroni.Handler{
		negroni.Handler(negroni.NewLogger()),
		negroni.Handler(negroni.NewRecovery()),
		negroni.Handler(negroni.HandlerFunc(CustomMiddleware)),
	}
	// Get pREST app
	n := middlewares.GetApp()
	// Get pPREST router
	r := router.Get()
	// Register custom routes
	r.HandleFunc("/ping", Pong).Methods("GET")
	// Call pREST cmd
	cmd.Execute()
}