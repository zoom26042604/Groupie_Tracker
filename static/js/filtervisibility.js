document.addEventListener('DOMContentLoaded', function() {
    const filterSection = document.querySelector('.filters-section');
    const toggleButton = document.querySelector('.toggle-filters');
    const yearToggle = document.getElementById('yearToggle');
    const checkboxGroup = yearToggle.nextElementSibling.nextElementSibling;

    const isFilterVisible = localStorage.getItem('filterMenuVisible') === 'true';
    if (isFilterVisible) {
        filterSection.classList.add('show');
    }

    toggleButton.addEventListener('click', function() {
        filterSection.classList.toggle('show');
        localStorage.setItem('filterMenuVisible', filterSection.classList.contains('show'));
    });

    const isYearRangeVisible = localStorage.getItem('yearRangeVisible') === 'true';
    yearToggle.checked = isYearRangeVisible;
    checkboxGroup.style.display = isYearRangeVisible ? 'block' : 'none';

    yearToggle.addEventListener('change', function() {
        localStorage.setItem('yearRangeVisible', this.checked);
        checkboxGroup.style.display = this.checked ? 'block' : 'none';
    });
});
document.addEventListener('DOMContentLoaded', function() {
    const locationInput = document.getElementById('location');
    const form = locationInput.closest('form');

    locationInput.addEventListener('input', function(e) {
        const activeField = form.querySelector('[name=activeField]');
        activeField.value = 'location';

        const currentValue = this.value;
        this.dataset.lastValue = currentValue;

        form.submit();
    });

    if (locationInput.hasAttribute('autofocus')) {
        locationInput.focus();
        if (locationInput.value) {
            locationInput.setSelectionRange(locationInput.value.length, locationInput.value.length);
        }
    }
});