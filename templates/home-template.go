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
<body class="min-h-screen bg-gradient-to-b from-pink-50 to-white">
  <!-- Hero Section -->
  <section class="px-4 py-20 md:py-32">
    <div class="container mx-auto text-center">
      <div class="inline-flex items-center justify-center rounded-full bg-pink-100 px-3 py-1 text-sm text-pink-700 mb-8">
        <svg class="mr-2 h-4 w-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
        Join our next event
      </div>
      <h1 class="text-4xl md:text-6xl font-bold mb-6 bg-gradient-to-r from-pink-600 to-rose-600 bg-clip-text text-transparent">
        Find Your Perfect Match
      </h1>
      <p class="text-lg md:text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
        Join us for an evening of meaningful connections, laughter, and the possibility of finding that special someone.
        Our curated dating events bring together like-minded individuals in a comfortable setting.
      </p>
      <div class="flex flex-wrap justify-center gap-4 mb-12">
        <div class="flex items-center gap-2 text-gray-600">
          <svg class="h-5 w-5 text-pink-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3M16 7V3M5 11h14M5 19h14m-7-8h.01" />
          </svg>
          <span>February 14, 2024</span>
        </div>
        <div class="flex items-center gap-2 text-gray-600">
          <svg class="h-5 w-5 text-pink-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-9a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <span>7:00 PM - 10:00 PM</span>
        </div>
        <div class="flex items-center gap-2 text-gray-600">
          <svg class="h-5 w-5 text-pink-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 10.5a8.38 8.38 0 01-.9 3.8l-3.9 7a2 2 0 01-3.6 0l-3.9-7A8.38 8.38 0 013 10.5C3 5.81 7.03 2 12 2s9 3.81 9 8.5z" />
          </svg>
          <span>The Grand Hotel, Downtown</span>
        </div>
      </div>
    </div>
  </section>

  <!-- How It Works Section -->
  <section class="px-4 py-16 bg-gray-50">
    <div class="container mx-auto">
      <h2 class="text-3xl font-bold text-center mb-12">How It Works</h2>
      <div class="grid md:grid-cols-3 gap-8">
        <div class="border rounded-lg p-6 shadow-md">
          <div class="rounded-full bg-pink-100 w-12 h-12 flex items-center justify-center mb-4">
            <svg class="h-6 w-6 text-pink-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
            </svg>
          </div>
          <h3 class="font-semibold text-xl mb-2">Curated Matches</h3>
          <p class="text-gray-600">
            We carefully select participants to ensure a balanced and compatible group of singles.
          </p>
        </div>
        <div class="border rounded-lg p-6 shadow-md">
          <div class="rounded-full bg-pink-100 w-12 h-12 flex items-center justify-center mb-4">
            <svg class="h-6 w-6 text-pink-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-9a9 9 0 11-18 0 9 9 0 0118 0z" />
            </svg>
          </div>
          <h3 class="font-semibold text-xl mb-2">Quality Time</h3>
          <p class="text-gray-600">
            Enjoy meaningful conversations in a structured format designed to help you connect.
          </p>
        </div>
        <div class="border rounded-lg p-6 shadow-md">
          <div class="rounded-full bg-pink-100 w-12 h-12 flex items-center justify-center mb-4">
            <svg class="h-6 w-6 text-pink-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 10.5a8.38 8.38 0 01-.9 3.8l-3.9 7a2 2 0 01-3.6 0l-3.9-7A8.38 8.38 0 013 10.5C3 5.81 7.03 2 12 2s9 3.81 9 8.5z" />
            </svg>
          </div>
          <h3 class="font-semibold text-xl mb-2">Fun Activities</h3>
          <p class="text-gray-600">
            Participate in engaging ice-breakers and activities that make connecting natural and enjoyable.
          </p>
        </div>
      </div>
    </div>
  </section>
    <!-- Upcoming Events Section -->
    <section class="px-4 py-16 bg-white">
      <div class="container mx-auto">
        <h2 class="text-3xl font-bold text-center mb-12">Upcoming Events</h2>
        <div 
          class="grid md:grid-cols-3 gap-8"
        >
         
          <div class="text-center col-span-3">
           <!-- Loading Placeholder -->
            <div class="animate-pulse"
             hx-get="/api/events" 
            hx-trigger="load" 
            hx-target="this" 
            hx-swap="innerHTML" >
            </div>
          </div>
        </div>
      </div>
    </section>
    

  <!-- Reserve Spot Section -->
  <section class="px-4 py-16">
    <div class="container mx-auto max-w-xl">
      <div class="border rounded-lg p-6 shadow-md">
        <h2 class="text-2xl font-bold text-center mb-6">Reserve Your Spot</h2>
        <form class="space-y-6">
          <div class="grid grid-cols-2 gap-4">
            <div class="space-y-2">
              <label for="first-name" class="block text-sm font-medium text-gray-700">First name</label>
              <input id="first-name" name="first-name" type="text" placeholder="John" required class="block w-full border-gray-300 rounded-md shadow-sm focus:ring-pink-500 focus:border-pink-500" />
            </div>
            <div class="space-y-2">
              <label for="last-name" class="block text-sm font-medium text-gray-700">Last name</label>
              <input id="last-name" name="last-name" type="text" placeholder="Doe" required class="block w-full border-gray-300 rounded-md shadow-sm focus:ring-pink-500 focus:border-pink-500" />
            </div>
          </div>
          <div class="space-y-2">
            <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
            <input id="email" name="email" type="email" placeholder="john@example.com" required class="block w-full border-gray-300 rounded-md shadow-sm focus:ring-pink-500 focus:border-pink-500" />
          </div>
          <div class="space-y-2">
            <label for="phone" class="block text-sm font-medium text-gray-700">Phone number</label>
            <input id="phone" name="phone" type="tel" placeholder="+1 (555) 000-0000" required class="block w-full border-gray-300 rounded-md shadow-sm focus:ring-pink-500 focus:border-pink-500" />
          </div>
          <div class="space-y-2">
            <label for="about" class="block text-sm font-medium text-gray-700">Tell us about yourself</label>
            <textarea id="about" name="about" placeholder="Share a bit about your interests and what you're looking for..." rows="3" class="block w-full border-gray-300 rounded-md shadow-sm focus:ring-pink-500 focus:border-pink-500"></textarea>
          </div>
          <div class="space-y-4">
            <label class="block text-sm font-medium text-gray-700">Payment Method</label>
            <div class="grid grid-cols-2 gap-4">
              <div class="flex items-center space-x-2">
                <input type="radio" id="card" name="payment-method" value="card" class="focus:ring-pink-500 focus:border-pink-500" />
                <label for="card" class="text-sm font-medium text-gray-700">Credit Card</label>
              </div>
              <div class="flex items-center space-x-2">
                <input type="radio" id="wallet" name="payment-method" value="wallet" class="focus:ring-pink-500 focus:border-pink-500" />
                <label for="wallet" class="text-sm font-medium text-gray-700">Digital Wallet</label>
              </div>
            </div>
          </div>
          <button type="submit" class="w-full bg-pink-600 hover:bg-pink-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:ring-2 focus:ring-pink-500 focus:ring-opacity-50">
            Register Now
          </button>
        </form>
      </div>
    </div>
  </section>

  <!-- Footer -->
  <footer class="bg-gray-50 py-8">
    <div class="container mx-auto px-4 text-center text-gray-600">
      <p>&copy; 2024 Dating Event. All rights reserved.</p>
    </div>
  </footer>
</body>
</html>
`
