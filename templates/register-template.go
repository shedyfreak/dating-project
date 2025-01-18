package templates

const registerTmpl = ` <!-- Reserve Spot Section -->
<section class="px-4 py-16">
  <div class="container mx-auto max-w-xl">
    <div class="border rounded-lg p-6 shadow-md">
      <h2 class="text-2xl font-bold text-center mb-6">Reservez Votre Place</h2>
      <form 
        class="space-y-6" 
        hx-post="/api/register"
        hx-trigger="submit"
		    hx-swap="none">
        <div class="space-y-2">
          <label for="event" class="block text-sm font-medium text-gray-700">Choisis un évenement</label>
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
            <label for="first-name" class="block text-sm font-medium text-gray-700">Prénom</label>
            <input id="first-name" name="first_name" type="text" placeholder="Votre Prénom" required class="block w-full border-gray-300 rounded-md shadow-sm focus:ring-pink-500 focus:border-pink-500" />
          </div>
          <div class="space-y-2">
            <label for="last-name" class="block text-sm font-medium text-gray-700">Nom de famille</label>
            <input id="last-name" name="last_name" type="text" placeholder="Votre Nom de famille" required class="block w-full border-gray-300 rounded-md shadow-sm focus:ring-pink-500 focus:border-pink-500" />
          </div>
        </div>
        
        <div class="space-y-2">
          <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
          <input id="email" name="email" type="email" placeholder="john@example.com" required class="block w-full border-gray-300 rounded-md shadow-sm focus:ring-pink-500 focus:border-pink-500" />
        </div>
        
        <div class="space-y-2">
          <label for="phone" class="block text-sm font-medium text-gray-700">Numéro de téléphone</label>
          <input id="phone" name="phone" type="tel" placeholder="+41 12 345 4567" required class="block w-full border-gray-300 rounded-md shadow-sm focus:ring-pink-500 focus:border-pink-500" />
        </div>
        <div class="space-y-2">
          <label for="birthdate" class="block text-sm font-medium text-gray-700">
    Birthdate
  </label>
  <input 
    id="birthdate" 
    name="Date de naissance" 
    type="date" 
    placeholder="YYYY-MM-DD" 
    required 
    class="block w-full border-gray-300 rounded-md shadow-sm focus:ring-pink-500 focus:border-pink-500 placeholder-gray-400 text-gray-700"
  />
</div>  
        <button type="submit" class="w-full bg-pink-600 hover:bg-pink-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:ring-2 focus:ring-pink-500 focus:ring-opacity-50">
          Réserver ma place
        </button>
      </form>
    </div>
  </div>
</section>
`
