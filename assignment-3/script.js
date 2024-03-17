const waterValueElement = document.getElementById("water-value");
const windValueElement = document.getElementById("wind-value");
const waterStatusElement = document.getElementById("water-status");
const windStatusElement = document.getElementById("wind-status");

function updateStatus(data) {
    waterValueElement.textContent = data.water;
    windValueElement.textContent = data.wind;

    let waterStatus = "aman";
    if (data.water >= 9) {
        waterStatus = "bahaya";
    } else if (data.water >= 5) {
        waterStatus = "siaga";
    }

    let windStatus = "aman";
    if (data.wind >= 16) {
        windStatus = "bahaya";
    } else if (data.wind >= 6) {
        windStatus = "siaga";
    }

    waterStatusElement.textContent = waterStatus;
    windStatusElement.textContent = windStatus;
}

// Fetch data from the microservice
fetch("/status")
    .then((response) => response.json())
    .then((data) => updateStatus(data));
