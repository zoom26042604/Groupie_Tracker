document.addEventListener('DOMContentLoaded', function() {
    // Slider de localisations
    const locationSlider = document.getElementById('locationSlider');
    const locationSlides = locationSlider.children;
    const locationSlideCount = locationSlides.length;
    let currentLocationSlide = 0;

    // Slider de dates
    const dateSlider = document.getElementById('dateSlider');
    const dateSlides = dateSlider.children;
    const dateSlideCount = dateSlides.length;
    let currentDateSlide = 0;

    // Fonction pour passer à la slide suivante pour les localisations
    function nextLocationSlide() {
        currentLocationSlide = (currentLocationSlide + 1) % locationSlideCount;
        locationSlider.style.transform = `translateX(-${currentLocationSlide * 100}%)`;
    }

    // Fonction pour passer à la slide suivante pour les dates
    function nextDateSlide() {
        currentDateSlide = (currentDateSlide + 1) % dateSlideCount;
        dateSlider.style.transform = `translateX(-${currentDateSlide * 100}%)`;
    }

    // Défilement automatique toutes les 3 secondes pour les deux sliders
    setInterval(nextLocationSlide, 3000);
    setInterval(nextDateSlide, 3000);
});