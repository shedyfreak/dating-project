package templates

const SuccessPopUp = `<div id="successModal" class="fixed inset-0 bg-gray-500 bg-opacity-50 flex justify-center items-center transition-opacity duration-300">
    <div class="bg-white rounded-2xl shadow-2xl max-w-2xl w-full transform transition-all duration-500 scale-95 hover:scale-100">
        <!-- Modal Header -->
        <div class="p-8 border-b-2 border-gray-200 rounded-t-2xl">
            <h3 class="text-3xl font-bold text-gray-800">Inscription à l'Événement Confirmée!</h3>
        </div>
        <!-- Modal Body -->
        <div class="p-8">
            <p class="text-xl text-gray-700 mb-6">
                Vous êtes inscrit à l'événement le <span class="font-semibold text-pink-600">15.01.2025</span> à 
                <a href="https://maps.app.goo.gl/7U8YEJgAyJoB9XTd9" target="_blank" 
                class="font-semibold text-blue-600 hover:text-blue-800 transition duration-200">Pink Crocodilo</a>, à 18h.
            </p>
            <p class="text-lg text-gray-600">
                Nous sommes impatients de vous y retrouver! Cliquez sur le lien pour obtenir l'itinéraire ou plus d'infos.
            </p>
            <div class="flex justify-center mt-6">
                <button class="px-6 py-3 bg-gray-300 text-gray-700 font-semibold rounded-lg hover:bg-gray-400 transition duration-200" 
                        hx-on="click: this.closest('#successModal').style.display='none'">
                    Fermer
                </button>
            </div>
        </div>
    </div>
</div>
`
