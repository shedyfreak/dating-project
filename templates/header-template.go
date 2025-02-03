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

<header class="bg-gradient-to-r from-[#fce4ec] via-[#f8bbd0] to-[#ffffff] text-white shadow-md flex items-center" style="height: 3cm; overflow: hidden;">
  <div class="container mx-auto flex items-center justify-center h-full px-6">
    <!-- Logo in the center, fully adapting to header height -->
    <div class="h-full flex justify-center">
      <img src="assets/logo4.svg" alt="Affinitys Logo" class="h-full w-auto">
    </div>
  </div>
</header>

`
