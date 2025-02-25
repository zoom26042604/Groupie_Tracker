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




document.addEventListener('DOMContentLoaded', function() {
    const slider2 = document.getElementById('dateSlider');
    const slides2 = slider2.children;
    const slideCount2 = slides2.length;
    let currentSlide2 = 0;

    if (slideCount2 <= 1) return;

    const slideDuration = 3000;

    function nextSlide() {
        currentSlide2 = (currentSlide2 + 1) % slideCount2;
        slider2.style.transform = `translateX(-${currentSlide2 * 100}%)`;
    }

    setInterval(nextSlide, slideDuration);
});