package templates

const ContactPageTempl = `
<!-- Contact Modal -->
  <div
    id="contact-modal"
    class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
  >
    <div class="bg-white rounded-lg shadow-lg max-w-xl w-full p-6 relative">

      <!-- Form Content -->
      <h2 class="text-2xl font-bold text-center mb-6 text-gray-800">Contact Us</h2>
      <form class="space-y-6" action="/submit-contact" method="POST">
        <!-- Name Field -->
        <div class="space-y-2">
          <label for="name" class="block text-sm font-medium text-gray-700">Your Name</label>
          <input
            type="text"
            id="name"
            name="name"
            placeholder="Enter your name"
            required
            class="block w-full border border-gray-300 rounded-lg py-3 px-4 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 shadow-sm"
          />
        </div>

        <!-- Email Field -->
        <div class="space-y-2">
          <label for="email" class="block text-sm font-medium text-gray-700">Email Address</label>
          <input
            type="email"
            id="email"
            name="email"
            placeholder="you@example.com"
            required
            class="block w-full border border-gray-300 rounded-lg py-3 px-4 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 shadow-sm"
          />
        </div>

        <!-- Subject Field -->
        <div class="space-y-2">
          <label for="subject" class="block text-sm font-medium text-gray-700">Subject</label>
          <input
            type="text"
            id="subject"
            name="subject"
            placeholder="What's your message about?"
            required
            class="block w-full border border-gray-300 rounded-lg py-3 px-4 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 shadow-sm"
          />
        </div>

        <!-- Message Field -->
        <div class="space-y-2">
          <label for="message" class="block text-sm font-medium text-gray-700">Your Message</label>
          <textarea
            id="message"
            name="message"
            placeholder="Write your message here"
            rows="5"
            required
            class="block w-full border border-gray-300 rounded-lg py-3 px-4 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500 resize-none shadow-sm"
          ></textarea>
        </div>

        <!-- Submit Button -->
        <button
          type="submit"
          class="w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-4 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-opacity-50 transition duration-150 ease-in-out"
        >
          Send Message
        </button>
      </form>
    </div>
  </div>
   <!-- JavaScript to Handle Modal -->
  <script>
    const modal = document.getElementById('contact-modal');
    // Close modal on clicking outside of the form
    window.addEventListener('click', (e) => {
      if (e.target === modal) {
        modal.classList.add('hidden');
      }
    });
  </script>
`
