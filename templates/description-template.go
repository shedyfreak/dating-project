package templates

const descriptionTempl = ` 
<style>
  /* Background Image Styling */
  .hero-section {
    background-image: url('assets/background2.jpeg'); 
    background-size: cover;
    background-position: center;
    background-attachment: fixed;
    transition: opacity 0.5s ease-out;
  }
</style>

<section class="hero-section relative px-4 py-20 md:py-32">
  <div class="absolute inset-0 bg-gradient-to-b from-white/50 via-white/30 to-transparent"></div> <!-- Gradient Overlay -->
  
  <div class="container mx-auto text-center relative">
    <div class="inline-flex items-center justify-center rounded-full bg-pink-100 px-3 py-1 text-sm text-pink-700 mb-8">
      <svg class="mr-2 h-4 w-4" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
      </svg>
      Rejoins notre prochain événement
    </div>
    
    <h1 class="text-4xl md:text-6xl font-bold mb-6 bg-gradient-to-r from-pink-600 to-rose-600 bg-clip-text text-transparent">
      Trouve Ton Match Parfait
    </h1>
    
    <p class="text-lg md:text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
      Rejoignez-nous pour une soirée de connexions authentiques, de rires et peut-être la chance de trouver ta personne spéciale.
      Nos événements de dating soigneusement organisés rassemblent des personnes partageant les mêmes valeurs dans une ambiance conviviale.
    </p>
  </div>
</section>

<script>
  document.addEventListener("scroll", function () {
    let scrollPosition = window.scrollY;
    let heroSection = document.querySelector(".hero-section");

    /* Adjust opacity based on scroll */
    heroSection.style.opacity = 1 - scrollPosition / 600; 
  });
</script>
`
