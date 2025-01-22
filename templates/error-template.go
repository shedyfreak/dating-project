package templates

const errTemplate = ` <div class="max-w-md w-full p-4">
    <!-- Notification Container -->
    <div class="flex items-start bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded-lg shadow-md" role="alert">
      <!-- Icon -->
      <svg class="w-6 h-6 flex-shrink-0 mr-3" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.364 5.636l-1.414-1.414M5.636 5.636L4.222 4.222M12 9v3m0 4h.01M21 12c0-4.97-4.03-9-9-9S3 7.03 3 12s4.03 9 9 9 9-4.03 9-9z" />
      </svg>
      <!-- Message -->
      <div>
        <strong class="font-semibold">Erreur:</strong>
        <span class="block sm:inline">{{ErrorMessage}} - {{Contact}} </span>
      </div>
      <!-- Close Button -->
      <button class="ml-auto text-red-500 hover:text-red-700 focus:outline-none" aria-label="Close">
        <svg class="w-5 h-5" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12" />
        </svg>
      </button>
    </div>
  </div>`
