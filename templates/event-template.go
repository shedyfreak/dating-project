package templates

const allEvsTempl = ` <!-- Upcoming Events Section -->
    <section class="px-4 py-16 bg-white">
      <div class="container mx-auto">
        <h2 class="text-3xl font-bold text-center mb-12">Upcoming Events</h2>
        <div 
          class="grid md:grid-cols-3 gap-8" >
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
      </div>
    </section>`

const EventTempl = `
<div class="border rounded-lg p-6 shadow-md">
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
</div>
`
