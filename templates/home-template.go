package templates

func GetHomeWithAddOns(templates ...string) string {
	var addOns string
	for _, tmpl := range templates {
		addOns += tmpl
	}
	var homeTempl = `
<!DOCTYPE html>
<html lang="en">

<body class="min-h-screen bg-gradient-to-b from-pink-50 to-white">` +
		headertmpl +
		descriptionTempl +
		about +
		allEvsTempl +
		registerTmpl +
		addOns +
		footerTempl +
		`</body></html>`
	return homeTempl
}
