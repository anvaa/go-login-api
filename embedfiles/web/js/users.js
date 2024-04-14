async function setAuth(uid) {
    
    const isAuth = document.getElementById("_auth" + uid).value;
    const messageElement = document.getElementById("message");

    var authData = {
        id: uid,
        isauth: isAuth,
    };

    try {
        const response = await fetch("/user/auth", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(authData),
        });

        if (!response.ok) {
            messageElement.innerHTML = error.message;
            messageElement.style.border = "1px solid red";
            return;
        }

        window.location.href = "/v/users";
    } catch (error) {
        messageElement.innerHTML = "Change auth failed: " + error.message;
        messageElement.style.border = "1px solid red";
    }
}