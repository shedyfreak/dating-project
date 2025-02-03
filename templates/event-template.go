package templates

const allEvsTempl = `<!-- EVENTS SECTION -->
<script>
    htmx.defineExtension('show', {
        onEvent: function (el) {
            const targetSelector = el.getAttribute('hx-target');
            const modal = document.querySelector(targetSelector);
            if (modal) {
                modal.classList.remove('hidden');
            } else {
                console.error("No element found for selector: ${targetSelector}");
            }
        }
    });
    htmx.defineExtension('hide', {
        onEvent: function (el) {
            const targetSelector = el.getAttribute('hx-target');
            const modal = document.querySelector(targetSelector);
            if (modal) {
                modal.classList.add('hidden');
            } else {
                console.error("No element found for selector: ${targetSelector}");
            }
        }
    });
</script>


<section class="px-4 py-16  max-w-screen-lg mx-auto">
  <h2 class="text-3xl font-bold text-center mb-12">Événements à venir</h2>
  
  <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
    <div class="text-center col-span-3">
      <!-- BEGIN upcoming events Placeholder -->
      <div 
        hx-get="/api/events" 
        hx-trigger="load" 
        hx-target="this" 
        hx-swap="innerHTML" >
      </div>
      <!-- END upcoming events Placeholder -->
    </div>
  </div>
</section>`

const EventTempl = `
<div class="w-full p-6 bg-white rounded-xl shadow-lg hover:shadow-xl transition duration-300 transform hover:scale-105 mb-8">
  <div class="flex items-center justify-between">
    <!-- Event Content Section -->
    <div class="w-2/3 pr-6">
      <h3 class="font-semibold text-xl mb-4">{{.Name}}</h3>
      <div class="space-y-3 text-gray-600 mb-6">
        <div class="flex items-center gap-2">
          <svg class="h-4 w-4 text-pink-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3M16 7V3M5 11h14M5 19h14m-7-8h.01" />
          </svg>
          <span>{{.Date}}</span>
        </div>
        <div class="flex items-center gap-2">
          <svg class="h-4 w-4 text-pink-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-9a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <span>{{.Time}}</span>
        </div>
        <div class="flex items-center gap-2">
          <svg class="h-4 w-4 text-pink-600" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 10.5a8.38 8.38 0 01-.9 3.8l-3.9 7a2 2 0 01-3.6 0l-3.9-7A8.38 8.38 0 013 10.5C3 5.81 7.03 2 12 2s9 3.81 9 8.5z" />
          </svg>
          <span>{{.Location}}</span>
        </div>
      </div>
      <p class="text-sm text-gray-600 mb-4">
        {{.Description}}
      </p>
      <div class="text-pink-600 font-semibold">$49.99</div>
      
      <!-- Action Button -->
      <div class="mt-6 flex justify-center">
        <button  hx-on="click: document.querySelector('#modal-{{.ID}}').classList.remove('hidden')" class="px-6 py-2 bg-pink-600 text-white rounded-lg font-semibold hover:bg-pink-700 transition duration-200">
         Plus d'informations
        </button>
      </div>
    </div>

    <!-- Event Image Section -->
    <div class="w-1/3 rounded-lg overflow-hidden shadow-xl hover:scale-105 transition duration-300">
      <img src="https://dynamic-media-cdn.tripadvisor.com/media/photo-o/2d/72/a4/3f/caption.jpg" alt="Event Image" class="w-full h-full object-cover rounded-lg transform hover:scale-110 transition duration-300" />
    </div>
  </div>
</div>
` + FullDescModal

const FullDescModal = `
<!-- Modal -->
<div class="fixed inset-0 bg-black bg-opacity-50 z-50 flex items-center justify-center hidden" id="modal-{{.ID}}">
    <div class="bg-white text-left rounded-2xl shadow-lg p-6 w-full max-w-lg">
        <div class="mb-4">
            <h2 class="text-2xl font-bold text-pink-600" id="event-name">{{.Name}}</h2>
        </div>
        <div class="space-y-4">
            <div>
                <h3 class="text-lg font-semibold text-pink-500">Location:</h3>
                <p class="text-gray-700" id="event-location">{{.Location}}</p>
            </div>
            <div>
                <h3 class="text-lg font-semibold text-pink-500">Date and Time:</h3>
                <p class="text-gray-700" id="event-date-time">{{.Date}} - {{.Time}}</p>
            </div>
            <div>
                <h3 class="text-lg font-semibold text-pink-500">Description:</h3>
                <p class="text-gray-700" id="event-description">{{.LongDescription}}</p>
            </div>
            <div>
                <h3 class="text-lg font-semibold text-pink-500">Age Range:</h3>
                <p class="text-gray-700" id="event-age-range">{{.MinAge}} - {{.MaxAge}}</p>
            </div>
        </div>
        <div class="mt-6 flex justify-between">
            <button id="register-button" class="bg-pink-500 text-white py-2 px-4 rounded-lg hover:bg-pink-600"
                    hx-on="click: document.querySelector('#modal-{{.ID}}').classList.add('hidden'); 
                    document.querySelector('#register-section').scrollIntoView({ behavior: 'smooth' })">Register</button>
            <button class="bg-gray-300 text-gray-700 py-2 px-4 rounded-lg hover:bg-gray-400"
                     hx-on="click: document.querySelector('#modal-{{.ID}}').classList.add('hidden')">Close</button>
        </div>
    </div>
</div>
`
