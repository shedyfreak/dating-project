package templates

const RegisterTmpl = ` <!-- Reserve Spot Section -->
 <section class="px-4 py-16">
  <div class="container mx-auto max-w-xl">
    <div class="border rounded-lg p-6 shadow-md">
      <h2 class="text-2xl font-bold text-center mb-6">Reserve Your Spot</h2>
      <form class="space-y-6">
         <div class="space-y-2">
            <label for="event" class="block text-sm font-medium text-gray-700">Pick an Event</label>
            <select id="event" name="event" required 
              class="block w-full border-gray-300 rounded-md shadow-sm focus:ring-pink-500 focus:border-pink-500" 
              hx-get="/api/event-opts" 
              hx-trigger="load" 
              hx-target="#event" 
              hx-swap="innerHTML">
              <!-- Placeholder option -->
              <option value="" disabled selected>Select an event...</option>
           </select>
        </div>
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
          <input id="phone" name="phone" type="tel" placeholder="+41 12 345 4567" required class="block w-full border-gray-300 rounded-md shadow-sm focus:ring-pink-500 focus:border-pink-500" />
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
`
