package entity

type RouteIdentify int

const (
	RouteIdentifyIndex RouteIdentify = iota + 1
	RouteIdentifyChat
	RouteIdentifyPlay
)

func GetBreadcrumbsMap() map[RouteIdentify][]RouteIdentify {
	return map[RouteIdentify][]RouteIdentify{
		RouteIdentifyIndex: {
			RouteIdentifyIndex,
		},
		RouteIdentifyChat: {
			RouteIdentifyIndex,
			RouteIdentifyChat,
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
			Index:     true,
			Template:  "index.tpl",
			Path:      "index.html",
			ClassIcon: "fa-solid fa-chart-line",
			Title:     "Dashboard",
		},
		RouteIdentifyChat: {
			Template:  "chat.tpl",
			Path:      "chat.html",
			ClassIcon: "fa-solid fa-comment",
			Title:     "Chat",
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
	Index       bool
	Template    string
	Path        string
	ClassIcon   string
	Active      bool
	Title       string
	Breadcrumbs []Route
}
