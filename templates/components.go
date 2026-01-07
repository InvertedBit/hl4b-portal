package templates

import (
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// Layout creates the base HTML layout
func Layout(title string, content ...g.Node) g.Node {
	return Doctype(
		HTML(
			Lang("en"),
			Head(
				Meta(Charset("utf-8")),
				Meta(Name("viewport"), Content("width=device-width, initial-scale=1")),
				TitleEl(g.Text(title)),
				Link(Rel("stylesheet"), Href("/static/css/output.css")),
				Script(Src("https://unpkg.com/htmx.org@1.9.10")),
			),
			Body(
				Class("bg-gray-100 min-h-screen"),
				g.Group(content),
			),
		),
	)
}

// NavBar creates a navigation bar
func NavBar(username string, isAdmin bool) g.Node {
	return Nav(
		Class("bg-blue-600 text-white p-4 shadow-lg"),
		Div(
			Class("container mx-auto flex justify-between items-center"),
			Div(
				Class("flex items-center space-x-4"),
				A(
					Href("/"),
					Class("text-xl font-bold hover:text-blue-200"),
					g.Text("HL4B Portal"),
				),
				A(
					Href("/user/dashboard"),
					Class("hover:text-blue-200"),
					g.Text("Dashboard"),
				),
				A(
					Href("/user/uploads"),
					Class("hover:text-blue-200"),
					g.Text("My Uploads"),
				),
				g.If(isAdmin,
					A(
						Href("/admin/dashboard"),
						Class("hover:text-blue-200 font-semibold"),
						g.Text("Admin"),
					),
				),
			),
			Div(
				Class("flex items-center space-x-4"),
				Span(Class("text-sm"), g.Text("Welcome, "+username)),
				Form(
					Method("post"),
					Action("/logout"),
					g.Attr("hx-post", "/logout"),
					g.Attr("hx-swap", "none"),
					Button(
						Type("submit"),
						Class("bg-red-500 hover:bg-red-600 px-4 py-2 rounded"),
						g.Text("Logout"),
					),
				),
			),
		),
	)
}

// Container creates a centered container
func Container(content ...g.Node) g.Node {
	return Div(
		Class("container mx-auto px-4 py-8"),
		g.Group(content),
	)
}

// Card creates a card component
func Card(content ...g.Node) g.Node {
	return Div(
		Class("bg-white rounded-lg shadow-md p-6"),
		g.Group(content),
	)
}

// ButtonComponent creates a button component
func ButtonComponent(text string, classes ...string) g.Node {
	baseClass := "px-4 py-2 rounded hover:opacity-90 transition"
	if len(classes) > 0 {
		baseClass += " " + classes[0]
	} else {
		baseClass += " bg-blue-600 text-white"
	}
	return Button(
		Type("button"),
		Class(baseClass),
		g.Text(text),
	)
}

// Alert creates an alert message
func Alert(message string, alertType string) g.Node {
	alertClass := "p-4 rounded mb-4 "
	switch alertType {
	case "error":
		alertClass += "bg-red-100 text-red-700 border border-red-400"
	case "success":
		alertClass += "bg-green-100 text-green-700 border border-green-400"
	case "info":
		alertClass += "bg-blue-100 text-blue-700 border border-blue-400"
	default:
		alertClass += "bg-gray-100 text-gray-700 border border-gray-400"
	}
	return Div(
		Class(alertClass),
		g.Text(message),
	)
}
