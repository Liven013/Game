let socket;
window.addEventListener("DOMContentLoaded", () => {
    const participantsList = document.getElementById("participants");
    const statusMessage = document.getElementById("statusMessage");

    // Подключение к WebSocket
    socket = new WebSocket("ws://" + window.location.host + "/ws");

    socket.onopen = function() {
        statusMessage.innerText = "Соединение установлено.";

        const data = {
            name: "Имя пользователя",
            role: "Роль пользователя"
        };
        socket.send(JSON.stringify(data));
    };

    socket.onmessage = function(event) {
        const newParticipant = document.createElement("li");
        newParticipant.textContent = event.data;
        participantsList.appendChild(newParticipant);
    };

    socket.onclose = function() {
        statusMessage.innerText = "Соединение закрыто.";
    };

    socket.onerror = function(error) {
        statusMessage.innerText = "Ошибка: " + error.message;
    };
});

function closeConnection() {
    if (socket) {
        socket.close();
    }
}

window.addEventListener("beforeunload", closeConnection);