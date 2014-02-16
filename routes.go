package urlrouter

type Routes map[string]Route

type Router struct {
	Context Routes
}

type Match struct {
	Params Params
	Pattern string
}

func New (patterns...string) *Router {
	router := &Router{make(Routes)}

	for _, pattern := range patterns {
		router.Context[pattern] = NewRoute(pattern)
	}

	return router
}

func (router *Router) Match (url string) *Match {
	var result *Match

	for _, route := range router.Context {
		if !route.Regex.MatchString(url) {
			continue
		}

		result = &Match{route.Params(url), route.Pattern}
	}

	return result
}
