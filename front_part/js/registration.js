const form = document.querySelector("form");
document.addEventListener("DOMContentLoaded", function() {
    const form = document.querySelector("form");
    form.addEventListener("submit", function(event) {
        event.preventDefault();
        const formData = {
            name: document.getElementById("name").value,
            role: document.getElementById("role").value
        };
        fetch("/submit", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body : JSON.stringify(formData)
        })
        .then(response => response.json())
        .then(data => {
            console.log("Data received:", data);
            if (data.redirect) {
                window.location.replace(data.redirect);
            } else {
                document.getElementById("message").innerText = `Ответ сервера: ${data.message}`;
            }
        })
        .catch(error => {
                console.error("Ошибка:", error);
        });
    });
});
