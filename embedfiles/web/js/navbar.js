
const messageElement = document.getElementById("_message");

function mngusersClick() {
    window.location.href = "/v/users";
}

function newUserClick() {
    window.location.href = "/v/newusers";
}

async function logoutClick() {
    
    try {
        const response = await fetch("/logout", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
        });

        if (!response.ok) {
            throw new Error("Server error");
        }

        window.location.href = "/";
    } catch (error) {
        messageElement.innerHTML = "Logout failed: " + error.message;
    }
}