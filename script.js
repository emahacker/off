let socialStartTime, webStartTime;
let socialTime = 0;
let webTime = 0;
let isTracking = false;

document.getElementById('startBtn').addEventListener('click', () => {
    if (!isTracking) {
        socialStartTime = Date.now();
        webStartTime = Date.now();
        isTracking = true;
        updateStatus("Tracking started...");
    }
});

document.getElementById('stopBtn').addEventListener('click', () => {
    if (isTracking) {
        const currentTime = Date.now();
        socialTime += currentTime - socialStartTime;
        webTime += currentTime - webStartTime;
        isTracking = false;
        updateStatus("Tracking stopped.");
        updateDisplay();
    }
});

document.getElementById('reportBtn').addEventListener('click', () => {
    sendReport();
});

function updateStatus(message) {
    document.getElementById('status').innerHTML += `<p>${message}</p>`;
}

function updateDisplay() {
    document.getElementById('socialTime').textContent = formatTime(socialTime);
    document.getElementById('webTime').textContent = formatTime(webTime);
}

function formatTime(ms) {
    const totalSeconds = Math.floor(ms / 1000);
    const hours = Math.floor(totalSeconds / 3600);
    const minutes = Math.floor((totalSeconds % 3600) / 60);
    const seconds = totalSeconds % 60;
    return `${hours.toString().padStart(2, '0')}:${minutes.toString().padStart(2, '0')}:${seconds.toString().padStart(2, '0')}`;
}

function sendReport() {
    fetch('/sendReport', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            socialTime: socialTime,
            webTime: webTime
        })
    })
    .then(response => response.json())
    .then(data => {
        if (data.success) {
            updateStatus("Report sent via email!");
        } else {
            updateStatus("Failed to send report.");
        }
    });
}
