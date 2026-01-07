package templates

import (
	"fmt"

	"github.com/InvertedBit/hl4b-portal/models"
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// AdminDashboard creates the admin dashboard page
func AdminDashboard(username string, userCount, uploadCount int) g.Node {
	return Layout("Admin Dashboard - HL4B Portal",
		NavBar(username, true),
		Container(
			H1(Class("text-3xl font-bold mb-6"), g.Text("Admin Dashboard")),
			Div(
				Class("grid grid-cols-1 md:grid-cols-3 gap-6 mb-6"),
				Card(
					H2(Class("text-xl font-semibold mb-2"), g.Text("Total Users")),
					P(Class("text-4xl font-bold text-blue-600"), g.Text(fmt.Sprintf("%d", userCount))),
				),
				Card(
					H2(Class("text-xl font-semibold mb-2"), g.Text("Total Uploads")),
					P(Class("text-4xl font-bold text-green-600"), g.Text(fmt.Sprintf("%d", uploadCount))),
				),
				Card(
					H2(Class("text-xl font-semibold mb-2"), g.Text("Quick Actions")),
					Div(
						Class("space-y-2 mt-2"),
						A(
							Href("/admin/users"),
							Class("block text-blue-600 hover:underline"),
							g.Text("→ Manage Users"),
						),
						A(
							Href("/admin/uploads"),
							Class("block text-blue-600 hover:underline"),
							g.Text("→ Manage Uploads"),
						),
					),
				),
			),
		),
	)
}

// AdminUsers creates the admin users management page
func AdminUsers(username string, users []models.PortalUser) g.Node {
	return Layout("Manage Users - HL4B Portal",
		NavBar(username, true),
		Container(
			H1(Class("text-3xl font-bold mb-6"), g.Text("Manage Users")),
			Card(
				Table(
					Class("w-full"),
					THead(
						Tr(
							Class("border-b"),
							Th(Class("text-left p-2"), g.Text("Username")),
							Th(Class("text-left p-2"), g.Text("User ID")),
							Th(Class("text-left p-2"), g.Text("Role")),
							Th(Class("text-left p-2"), g.Text("Created")),
							Th(Class("text-left p-2"), g.Text("Actions")),
						),
					),
					TBody(
						g.Group(g.Map(users, func(user models.PortalUser) g.Node {
							return AdminUserRow(user)
						})),
					),
				),
			),
		),
	)
}

// AdminUserRow creates a single user row in the admin table
func AdminUserRow(user models.PortalUser) g.Node {
	role := "User"
	if user.IsAdmin {
		role = "Admin"
	}

	adminValue := "true"
	adminText := "Make Admin"
	if user.IsAdmin {
		adminValue = "false"
		adminText = "Remove Admin"
	}

	return Tr(
		ID(fmt.Sprintf("user-%d", user.ID)),
		Class("border-b hover:bg-gray-50"),
		Td(Class("p-2"), g.Text(user.Username)),
		Td(Class("p-2 text-sm text-gray-600"), g.Text(user.UserID.String()[:8]+"...")),
		Td(Class("p-2"), g.Text(role)),
		Td(Class("p-2 text-sm"), g.Text(user.CreatedAt.Format("Jan 02, 2006"))),
		Td(
			Class("p-2"),
			Div(
				Class("flex space-x-2"),
				Form(
					g.Attr("hx-post", fmt.Sprintf("/admin/users/%d/role", user.ID)),
					g.Attr("hx-target", fmt.Sprintf("#user-%d", user.ID)),
					g.Attr("hx-swap", "outerHTML"),
					Input(
						Type("hidden"),
						Name("is_admin"),
						Value(adminValue),
					),
					Button(
						Type("submit"),
						Class("text-blue-600 hover:underline text-sm"),
						g.Text(adminText),
					),
				),
				Button(
					Type("button"),
					g.Attr("hx-delete", fmt.Sprintf("/admin/users/%d", user.ID)),
					g.Attr("hx-target", fmt.Sprintf("#user-%d", user.ID)),
					g.Attr("hx-swap", "outerHTML"),
					g.Attr("hx-confirm", "Are you sure you want to delete this user?"),
					Class("text-red-600 hover:underline text-sm"),
					g.Text("Delete"),
				),
			),
		),
	)
}

// AdminUploads creates the admin uploads management page
func AdminUploads(username string, uploads []models.Upload, users map[string]string) g.Node {
	return Layout("Manage Uploads - HL4B Portal",
		NavBar(username, true),
		Container(
			H1(Class("text-3xl font-bold mb-6"), g.Text("Manage All Uploads")),
			Card(
				Table(
					Class("w-full"),
					THead(
						Tr(
							Class("border-b"),
							Th(Class("text-left p-2"), g.Text("Filename")),
							Th(Class("text-left p-2"), g.Text("User")),
							Th(Class("text-left p-2"), g.Text("Size")),
							Th(Class("text-left p-2"), g.Text("Uploaded")),
							Th(Class("text-left p-2"), g.Text("Actions")),
						),
					),
					TBody(
						g.Group(g.Map(uploads, func(upload models.Upload) g.Node {
							return AdminUploadRow(upload, users)
						})),
					),
				),
			),
		),
	)
}

// AdminUploadRow creates a single upload row in the admin table
func AdminUploadRow(upload models.Upload, users map[string]string) g.Node {
	fileSize := fmt.Sprintf("%.2f KB", float64(upload.FileSize)/1024.0)
	if upload.FileSize > 1024*1024 {
		fileSize = fmt.Sprintf("%.2f MB", float64(upload.FileSize)/(1024.0*1024.0))
	}

	userName := users[upload.UserID.String()]
	if userName == "" {
		userName = "Unknown"
	}

	return Tr(
		ID(fmt.Sprintf("upload-%d", upload.ID)),
		Class("border-b hover:bg-gray-50"),
		Td(Class("p-2"), g.Text(upload.OriginalName)),
		Td(Class("p-2"), g.Text(userName)),
		Td(Class("p-2 text-sm"), g.Text(fileSize)),
		Td(Class("p-2 text-sm"), g.Text(upload.CreatedAt.Format("Jan 02, 2006"))),
		Td(
			Class("p-2"),
			Div(
				Class("flex space-x-2"),
				A(
					Href(fmt.Sprintf("/uploads/%s", upload.Filename)),
					Target("_blank"),
					Class("text-blue-600 hover:underline text-sm"),
					g.Text("Download"),
				),
				Button(
					Type("button"),
					g.Attr("hx-delete", fmt.Sprintf("/admin/uploads/%d", upload.ID)),
					g.Attr("hx-target", fmt.Sprintf("#upload-%d", upload.ID)),
					g.Attr("hx-swap", "outerHTML"),
					g.Attr("hx-confirm", "Are you sure you want to delete this file?"),
					Class("text-red-600 hover:underline text-sm"),
					g.Text("Delete"),
				),
			),
		),
	)
}
