package entity

type RouteIdentify int

const (
	RouteIdentifyIndex RouteIdentify = iota + 1
	RouteIdentifyPlay
	RouteIdentifyTimer
)

func GetBreadcrumbsMap() map[RouteIdentify][]RouteIdentify {
	return map[RouteIdentify][]RouteIdentify{
		RouteIdentifyIndex: {
			RouteIdentifyIndex,
		},
		RouteIdentifyTimer: {
			RouteIdentifyIndex,
			RouteIdentifyTimer,
		},
		RouteIdentifyPlay: {
			RouteIdentifyIndex,
			RouteIdentifyPlay,
		},
	}
}

func GetRouteIdentifyMap() map[RouteIdentify]Route {
	return map[RouteIdentify]Route{
		RouteIdentifyIndex: {
			Template:  "index.tpl",
			Path:      "index.html",
			ClassIcon: "fa-solid fa-chart-line",
			Title:     "Dashboard",
		},
		RouteIdentifyTimer: {
			Template:  "timer.tpl",
			Path:      "timer.html",
			ClassIcon: "fa-solid fa-clock",
			Title:     "Timer",
		},
		RouteIdentifyPlay: {
			Template:  "play.tpl",
			Path:      "play.html",
			ClassIcon: "fa-solid fa-circle-play",
			Title:     "Play",
		},
	}
}

func GetRouteIdentifyMapToRoute(key RouteIdentify) (route Route, routes map[RouteIdentify]Route) {
	routes = GetRouteIdentifyMap()

	route = routes[key]
	route.Active = true

	for _, breadcrumb := range GetBreadcrumbsMap()[key] {
		route.Breadcrumbs = append(route.Breadcrumbs, routes[breadcrumb])
	}

	routes[key] = route

	return
}

type Route struct {
	Template    string
	Path        string
	ClassIcon   string
	Active      bool
	Title       string
	Breadcrumbs []Route
}
