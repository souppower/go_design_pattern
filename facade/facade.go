package facade

var db = map[string]string{
	"a@a.com": "a",
	"b@b.com": "b",
}

type database struct {
}

func (d *database) getNameByMail(mail string) string {
	return db[mail]
}

type mdWriter struct {
}

func (mw *mdWriter) title(title string) string {
	return "# Welcome to " + title + "'s page!"
}

// PageMaker struct
type PageMaker struct {
}

// MakeWelcomePage returns page content
func (pm *PageMaker) MakeWelcomePage(mail string) string {
	database := database{}
	writer := mdWriter{}

	name := database.getNameByMail(mail)
	page := writer.title(name)

	return page
}
