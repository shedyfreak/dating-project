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
		` <header class="bg-gradient-to-r from-pink-500 to-rose-500 text-white shadow-md">
    <div class="container mx-auto flex items-center justify-between py-4 px-6">
      <!-- Logo Section -->
      <div class="flex items-center space-x-3">
        <!-- Icon -->
        <svg class="w-10 h-10 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M17 16v-2a2 2 0 00-2-2H9a2 2 0 00-2 2v2m5 6h.01M12 8v.01" />
        </svg>
        <!-- Logo Text -->
        <span class="text-2xl font-bold tracking-wide">Dating Is Good</span>
      </div>
      <!-- Navigation -->
      <nav>
        <a href="/contact" class="text-white text-lg font-medium hover:text-pink-300 transition duration-300">
          Contact Us
        </a>
      </nav>
    </div>
  </header>
` +
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
