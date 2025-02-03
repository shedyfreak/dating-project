package templates

const registerTmpl = `
<!-- Reserve Spot Section -->
<section id="register-section" class="px-4 py-16">
  <div class="container mx-auto max-w-xl">
    <div class="border rounded-lg p-6 shadow-md">
      <h2 class="text-2xl font-bold text-center mb-6">Reservez Votre Place</h2>
      <form 
      class="space-y-6" 
      hx-post="/api/register"
      hx-trigger="submit[validateForm()]"
      hx-swap="none">
        <div class="space-y-2">
          <label for="event" class="block text-sm font-medium text-gray-700">Choisis un évenement</label>
          <select id="event" name="event" required 
            class="block w-full border border-gray-300 rounded-lg py-3 px-4 focus:outline-none focus:ring-2 focus:ring-pink-500 focus:border-pink-500 shadow-sm transition duration-150 ease-in-out"
            hx-get="/api/event-opts" 
            hx-trigger="load" 
            hx-target="#event" 
            hx-swap="innerHTML">
            <!-- Placeholder option -->
            <option value="" disabled selected>Select an event...</option>
          </select>
          <p id="event-error" class="text-red-600 text-sm hidden">Please select an event.</p>
        </div>
        
        <div class="grid grid-cols-2 gap-4">
          <div class="space-y-2">
            <label for="first-name" class="block text-sm font-medium text-gray-700">Prénom</label>
            <input id="first-name" name="first_name" type="text" placeholder="Votre Prénom" required maxlength="20" 
              class="block w-full border border-gray-300 rounded-lg py-3 px-4 focus:outline-none focus:ring-2 focus:ring-pink-500 focus:border-pink-500 shadow-sm transition duration-150 ease-in-out" />
            <p id="first-name-error" class="text-red-600 text-sm hidden">Le prénom ne doit pas dépasser 20 caractères.</p>
          </div>
          <div class="space-y-2">
            <label for="last-name" class="block text-sm font-medium text-gray-700">Nom de famille</label>
            <input id="last-name" name="last_name" type="text" placeholder="Votre Nom de famille" required maxlength="20" 
              class="block w-full border border-gray-300 rounded-lg py-3 px-4 focus:outline-none focus:ring-2 focus:ring-pink-500 focus:border-pink-500 shadow-sm transition duration-150 ease-in-out" />
            <p id="last-name-error" class="text-red-600 text-sm hidden">Le nom de famille ne doit pas dépasser 20 caractères.</p>
          </div>
        </div>
        
        <div class="space-y-2">
          <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
          <input id="email" name="email" type="email" placeholder="john@example.com" required 
            class="block w-full border border-gray-300 rounded-lg py-3 px-4 focus:outline-none focus:ring-2 focus:ring-pink-500 focus:border-pink-500 shadow-sm transition duration-150 ease-in-out" />
          <p id="email-error" class="text-red-600 text-sm hidden">Veuillez entrer une adresse email valide.</p>
        </div>
        
        <div class="space-y-2">
          <label for="phone" class="block text-sm font-medium text-gray-700">Numéro de téléphone</label>
          <input id="phone" name="phone" type="tel" placeholder="+41 12 345 4567" required 
            class="block w-full border border-gray-300 rounded-lg py-3 px-4 focus:outline-none focus:ring-2 focus:ring-pink-500 focus:border-pink-500 shadow-sm transition duration-150 ease-in-out" />
          <p id="phone-error" class="text-red-600 text-sm hidden">Veuillez entrer un numéro de téléphone suisse ou français valide.</p>
        </div>

        <div class="space-y-2">
          <label for="birthdate" class="block text-sm font-medium text-gray-700">Date de naissance</label>
          <input 
            id="birthdate" 
            name="birthdate" 
            type="date" 
            placeholder="YYYY-MM-DD" 
            required 
            class="block w-full border border-gray-300 rounded-lg py-3 px-4 focus:outline-none focus:ring-2 focus:ring-pink-500 focus:border-pink-500 shadow-sm transition duration-150 ease-in-out" />
          <p id="birthdate-error" class="text-red-600 text-sm hidden">Vous devez avoir plus de 20 ans pour vous inscrire.</p>
        </div>
         <div class="space-y-2">
          <label for="event" class="block text-sm font-medium text-gray-700">Sexe</label>
          <select id="sex" name="sex" required 
            class="block w-full border border-gray-300 rounded-lg py-3 px-4 focus:outline-none focus:ring-2 focus:ring-pink-500 focus:border-pink-500 shadow-sm transition duration-150 ease-in-out">
            <option value="" disabled selected>Choisis ton sexe</option>
            <option value="male">Homme</option>
            <option value="female">Femme</option>
          </select>
          <p id="event-error" class="text-red-600 text-sm hidden">Veuillez choisir ton sexe</p>
        </div>

        <button type="submit" class="w-full bg-pink-600 hover:bg-pink-700 text-white font-bold py-3 px-4 rounded-lg focus:outline-none focus:ring-2 focus:ring-pink-500 focus:ring-opacity-50 transition duration-150 ease-in-out">
          Réserver ma place
        </button>
      </form>
    </div>
  </div>
</section>

<script>
function validateForm() {
  // Clear previous errors
  const errorElements = document.querySelectorAll('.text-red-600');
  errorElements.forEach(element => element.classList.add('hidden'));

  // Get form values
  const firstName = document.getElementById('first-name').value.trim();
  const lastName = document.getElementById('last-name').value.trim();
  const email = document.getElementById('email').value.trim();
  const phone = document.getElementById('phone').value.trim();
  const birthdate = document.getElementById('birthdate').value.trim();
  const event = document.getElementById('event').value;
  const sex = document.getElementById('sex').value;

  let isValid = true;

  // Validate first name
  if (!firstName || firstName.length > 20) {
    document.getElementById('first-name-error').classList.remove('hidden');
    isValid = false;
  }

  // Validate last name
  if (!lastName || lastName.length > 20) {
    document.getElementById('last-name-error').classList.remove('hidden');
    isValid = false;
  }

  // Validate email
  const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  if (!email || !emailPattern.test(email)) {
    document.getElementById('email-error').classList.remove('hidden');
    isValid = false;
  }

  // Validate phone number
  const phonePattern = /^(\+41|\+33)\s?\d{2}\s?\d{3}\s?\d{4}$/;
  if (!phone || !phonePattern.test(phone)) {
    document.getElementById('phone-error').classList.remove('hidden');
    isValid = false;
  }

  // Validate birthdate (ensure user is over 20)
  if (birthdate) {
    const birthdateObj = new Date(birthdate);
    const today = new Date();
    let age = today.getFullYear() - birthdateObj.getFullYear();
    const monthDiff = today.getMonth() - birthdateObj.getMonth();
    const dayDiff = today.getDate() - birthdateObj.getDate();

    if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
      age--; // Adjust the age if the current date is before the birthdate
    }

    if (age < 20) {
      document.getElementById('birthdate-error').classList.remove('hidden');
      isValid = false;
    }
  } else {
    document.getElementById('birthdate-error').classList.remove('hidden');
    isValid = false;
  }

  // Validate event selection
  if (!event) {
    document.getElementById('event-error').classList.remove('hidden');
    isValid = false;
  }

  // Validate sex selection
  if (!sex) {
    document.getElementById('event-error').classList.remove('hidden');
    isValid = false;
  }

  // Prevent form submission if any validation fails
  if (!isValid) {
    return false;
  }

  // If everything is valid, allow submission
  return true;
}
</script>
`
