
function mngusersClick() {
    window.location.href = "/v/users";
}

function userhomeClick() {
    window.location.href = "/v/home";
}

async function logoutClick() {
    var messageElement = document.getElementById("_message");
    
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