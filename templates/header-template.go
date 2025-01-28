package templates

const headertmpl = `
<head class="bg-gradient-to-r from-pink-500 to-rose-500 text-white shadow-md">
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Dating Event</title>
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=Courgette&family=Lobster&display=swap" rel="stylesheet">
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/tailwindcss@2.2.19/dist/tailwind.min.css">
  <script src="https://unpkg.com/htmx.org"></script>
</head>
	<header class="bg-gradient-to-r from-[#fce4ec] via-[#f8bbd0] to-[#ffffff] text-white shadow-md">
    <div class="container mx-auto flex items-center justify-between py-4 px-6">
      <!-- Logo Section -->
      <div class="flex items-center space-x-3">
      <!-- Logo Image -->
      <img src="assets/logo2.svg" alt="Affinitys Logo" class="w-15 h-15" />
      <!-- Logo Text -->
      <span class="text-2xl tracking-wide"  style="font-family: 'Courgette', cursive; color: #f567ca;">AFFINITYS</span>
    </div>
      <!-- Navigation -->
      <nav>
        <a href="/contact-us" class="text-pink-300 hover:text-pink-700 text-lg font-medium transition duration-150 ease-in-out">
          Contact Us
        </a>    
      </nav>
    </div>
  </header>
`
