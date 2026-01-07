package templates

import (
	"fmt"

	"github.com/InvertedBit/hl4b-portal/models"
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

// UserDashboard creates the user dashboard page
func UserDashboard(username string, isAdmin bool, uploadCount int) g.Node {
	accountType := "User"
	if isAdmin {
		accountType = "Administrator"
	}

	return Layout("Dashboard - HL4B Portal",
		NavBar(username, isAdmin),
		Container(
			H1(Class("text-3xl font-bold mb-6"), g.Text("User Dashboard")),
			Div(
				Class("grid grid-cols-1 md:grid-cols-2 gap-6"),
				Card(
					H2(Class("text-xl font-semibold mb-4"), g.Text("Your Statistics")),
					Div(
						Class("space-y-2"),
						P(g.Text(fmt.Sprintf("Username: %s", username))),
						P(g.Text(fmt.Sprintf("Total Uploads: %d", uploadCount))),
						P(g.Text(fmt.Sprintf("Account Type: %s", accountType))),
					),
				),
				Card(
					H2(Class("text-xl font-semibold mb-4"), g.Text("Quick Actions")),
					Div(
						Class("space-y-2"),
						A(
							Href("/user/uploads"),
							Class("block text-blue-600 hover:underline"),
							g.Text("→ Manage Your Uploads"),
						),
						g.If(isAdmin,
							A(
								Href("/admin/dashboard"),
								Class("block text-blue-600 hover:underline"),
								g.Text("→ Admin Dashboard"),
							),
						),
					),
				),
			),
		),
	)
}

// UserUploads creates the user uploads page
func UserUploads(username string, isAdmin bool, uploads []models.Upload) g.Node {
	return Layout("My Uploads - HL4B Portal",
		NavBar(username, isAdmin),
		Container(
			H1(Class("text-3xl font-bold mb-6"), g.Text("My Uploads")),
			Card(
				H2(Class("text-xl font-semibold mb-4"), g.Text("Upload New File")),
				Form(
					Method("post"),
					Action("/user/upload"),
					g.Attr("enctype", "multipart/form-data"),
					g.Attr("hx-post", "/user/upload"),
					g.Attr("hx-target", "#upload-list"),
					g.Attr("hx-swap", "beforeend"),
					Class("space-y-4"),
					Div(
						Input(
							Type("file"),
							Name("file"),
							Required(),
							Class("w-full px-3 py-2 border border-gray-300 rounded"),
						),
					),
					Button(
						Type("submit"),
						Class("bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 transition"),
						g.Text("Upload"),
					),
				),
			),
			Div(Class("mt-6")),
			Card(
				H2(Class("text-xl font-semibold mb-4"), g.Text("Your Files")),
				Div(
					ID("upload-list"),
					Class("space-y-2"),
					g.Group(g.Map(uploads, func(upload models.Upload) g.Node {
						return UploadItem(upload)
					})),
					g.If(len(uploads) == 0,
						P(Class("text-gray-500 text-center py-4"), g.Text("No uploads yet")),
					),
				),
			),
		),
	)
}

// UploadItem creates a single upload item component
func UploadItem(upload models.Upload) g.Node {
	fileSize := fmt.Sprintf("%.2f KB", float64(upload.FileSize)/1024.0)
	if upload.FileSize > 1024*1024 {
		fileSize = fmt.Sprintf("%.2f MB", float64(upload.FileSize)/(1024.0*1024.0))
	}

	return Div(
		ID(fmt.Sprintf("upload-%d", upload.ID)),
		Class("flex items-center justify-between p-4 border border-gray-200 rounded hover:bg-gray-50"),
		Div(
			Class("flex-1"),
			P(Class("font-medium"), g.Text(upload.OriginalName)),
			P(Class("text-sm text-gray-500"), g.Text(fileSize+" • "+upload.CreatedAt.Format("Jan 02, 2006"))),
		),
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
				g.Attr("hx-delete", fmt.Sprintf("/user/upload/%d", upload.ID)),
				g.Attr("hx-target", fmt.Sprintf("#upload-%d", upload.ID)),
				g.Attr("hx-swap", "outerHTML"),
				g.Attr("hx-confirm", "Are you sure you want to delete this file?"),
				Class("text-red-600 hover:underline text-sm"),
				g.Text("Delete"),
			),
		),
	)
}
