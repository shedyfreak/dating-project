package templates

const HomeTempl = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Dating Event</title>
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css">
  <script src="https://unpkg.com/htmx.org"></script>
</head>
<body class="min-h-screen bg-gradient-to-b from-pink-50 to-white">` +
	descriptionTempl +
	About +
	allEvsTempl +
	RegisterTmpl +
	`<!-- Footer -->
  <footer class="bg-gray-50 py-8">
    <div class="container mx-auto px-4 text-center text-gray-600">
      <p>&copy; 2024 Dating Event. All rights reserved.</p>
    </div>
  </footer>
</body>
</html>
`
