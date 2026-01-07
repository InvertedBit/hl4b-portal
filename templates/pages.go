package templates

import (
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// LoginPage creates the login page
func LoginPage(errorMsg string) g.Node {
	return Layout("Login - HL4B Portal",
		Div(
			Class("flex items-center justify-center min-h-screen"),
			Card(
				Div(
					Class("w-96"),
					H1(Class("text-2xl font-bold mb-6 text-center"), g.Text("HL4B Portal Login")),
					g.If(errorMsg != "",
						Alert(errorMsg, "error"),
					),
					Form(
						Method("post"),
						Action("/login"),
						Class("space-y-4"),
						Div(
							Label(
								Class("block text-sm font-medium text-gray-700 mb-1"),
								For("email"),
								g.Text("Email"),
							),
							Input(
								Type("email"),
								Name("email"),
								ID("email"),
								Required(),
								Class("w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"),
								Placeholder("your@email.com"),
							),
						),
						Div(
							Label(
								Class("block text-sm font-medium text-gray-700 mb-1"),
								For("password"),
								g.Text("Password"),
							),
							Input(
								Type("password"),
								Name("password"),
								ID("password"),
								Required(),
								Class("w-full px-3 py-2 border border-gray-300 rounded focus:outline-none focus:ring-2 focus:ring-blue-500"),
								Placeholder("••••••••"),
							),
						),
						Button(
							Type("submit"),
							Class("w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition"),
							g.Text("Login"),
						),
					),
				),
			),
		),
	)
}

// HomePage creates the home page
func HomePage(username string, isAdmin bool, isAuthenticated bool) g.Node {
	var content []g.Node

	if isAuthenticated {
		content = []g.Node{
			NavBar(username, isAdmin),
			Container(
				H1(Class("text-3xl font-bold mb-4"), g.Text("Welcome to HL4B Portal")),
				Card(
					P(Class("text-gray-700 mb-4"),
						g.Text("This is your homelab portal. Use the navigation above to access different features."),
					),
					Div(
						Class("space-y-2"),
						A(
							Href("/user/dashboard"),
							Class("block text-blue-600 hover:underline"),
							g.Text("→ Go to Dashboard"),
						),
						A(
							Href("/user/uploads"),
							Class("block text-blue-600 hover:underline"),
							g.Text("→ Manage Uploads"),
						),
						g.If(isAdmin,
							A(
								Href("/admin/dashboard"),
								Class("block text-blue-600 hover:underline font-semibold"),
								g.Text("→ Admin Dashboard"),
							),
						),
					),
				),
			),
		}
	} else {
		content = []g.Node{
			Div(
				Class("flex items-center justify-center min-h-screen"),
				Card(
					H1(Class("text-3xl font-bold mb-4"), g.Text("Welcome to HL4B Portal")),
					P(Class("text-gray-700 mb-4"), g.Text("Please login to continue.")),
					A(
						Href("/login"),
						Class("inline-block bg-blue-600 text-white px-6 py-2 rounded hover:bg-blue-700 transition"),
						g.Text("Login"),
					),
				),
			),
		}
	}

	return Layout("Home - HL4B Portal", content...)
}
