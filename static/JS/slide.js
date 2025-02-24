document.addEventListener('DOMContentLoaded', function() {
    const slider = document.getElementById('locationSlider');
    const slides = slider.children;
    const slideCount = slides.length;
    let currentSlide = 0;

    if (slideCount <= 1) return;

    const slideDuration = 3000;

    function nextSlide() {
        currentSlide = (currentSlide + 1) % slideCount;
        slider.style.transform = `translateX(-${currentSlide * 100}%)`;
    }

    setInterval(nextSlide, slideDuration);
});