mapboxgl.accessToken = 'pk.eyJ1IjoibGVyYXBob3UiLCJhIjoiY202a3R2MXNzMDFlYjJrcjV1NTh4N2l2ayJ9.x0B8KMcr28PJjV03zuK5Iw';
const map = new mapboxgl.Map({
    container: 'map',
    style: 'mapbox://styles/mapbox/streets-v9',
    projection: 'globe', // Display the map as a globe, since satellite-v9 defaults to Mercator
    zoom: 1,
    center: [30, 15]
});

map.addControl(new mapboxgl.NavigationControl());
map.scrollZoom.disable();

map.on('style.load', () => {
    map.setFog({}); // Set the default atmosphere style
});

// The following values can be changed to control rotation speed:

// At low zooms, complete a revolution every two minutes.
const secondsPerRevolution = 240;
// Above zoom level 5, do not rotate.
const maxSpinZoom = 5;
// Rotate at intermediate speeds between zoom levels 3 and 5.
const slowSpinZoom = 3;

let userInteracting = false;
const spinEnabled = true;

function spinGlobe() {
    const zoom = map.getZoom();
    if (spinEnabled && !userInteracting && zoom < maxSpinZoom) {
        let distancePerSecond = 360 / secondsPerRevolution;
        if (zoom > slowSpinZoom) {
            // Slow spinning at higher zooms
            const zoomDif =
                (maxSpinZoom - zoom) / (maxSpinZoom - slowSpinZoom);
            distancePerSecond *= zoomDif;
        }
        const center = map.getCenter();
        center.lng -= distancePerSecond;
        // Smoothly animate the map over one second.
        // When this animation is complete, it calls a 'moveend' event.
        map.easeTo({ center, duration: 1000, easing: (n) => n });
    }
}

// Pause spinning on interaction
map.on('mousedown', () => {
    userInteracting = true;
});
map.on('dragstart', () => {
    userInteracting = true;
});

// When animation is complete, start spinning if there is no ongoing interaction
map.on('moveend', () => {
    spinGlobe();
});

spinGlobe();

const mapboxClient = mapboxSdk({ accessToken: mapboxgl.accessToken });
async function afficherLocations() {
    console.log(" afficherLocations() appelée !"); // Debug
    let index = 0;
    for (const location of locations) {
        let formattedLocation = "";
        for (const char of location) {
            if (char === "_") {
                formattedLocation += " ";
            } else if (char === "-") {
                formattedLocation += ", ";
            } else {
                formattedLocation += char;
            }
        }

        try {
            const response = await mapboxClient.geocoding
                .forwardGeocode({
                    query: formattedLocation,
                    autocomplete: false,
                    limit: 1
                })
                .send();

            if (!response || !response.body || !response.body.features || !response.body.features.length) {
                console.error(`Géocodage échoué pour : ${formattedLocation}`);
                continue;
            }

            const feature = response.body.features[0];


            new mapboxgl.Marker()
                .setLngLat(feature.center)
                .setPopup(new mapboxgl.Popup().setHTML(
                    `<div style="color: blue; font-size: 16px;">${formattedLocation}</div>`
                ))
                .addTo(map);


            console.log(`Marqueur ajouté pour ${formattedLocation} à ${feature.center}`);
        } catch (error) {
            console.error(`Erreur Mapbox pour ${formattedLocation}:`, error);
        }
        index++;
    }
}
map.on('load', afficherLocations);