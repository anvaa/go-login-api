
async function delClick() {
        
    var messageElement = document.getElementById("_message");
    var uid = document.getElementById("_uid").value;
    var email = document.getElementById("_email").value;
    
    if (uid == "1") {
        messageElement.innerHTML = "Can´t delete superadmin!";
        return;
    }

    const verify = confirm("Are you sure you want to delete user " + email + 
                            "? \n\nYou can just remove auth to deny access.", "Delete user");
    if (!verify) {
        return;
    }

    
    var userData = {
        id: uid,
    };

    try {
        const response = await fetch("/user/delete/" + uid, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(userData),
        });
    
        if (!response.ok) {
        //   throw new Error('Server error');
            messageElement.innerHTML = error.message;
        }
    
        window.location.href = "/v/users";
        } catch (error) {
        messageElement.innerHTML = "Delete failed: " + error.message;

    }
}

async function setPswClick() {
        
    var messageElement = document.getElementById("_message");
    var uid = document.getElementById("_uid").value;
    var psw1 = document.getElementById("_password1").value;
    var psw2 = document.getElementById("_password2").value;

    if (!validatePasswords(psw1, psw2)) {
        return; // Message is set inside the validatePasswords function
    }

    var userData = {
        id: uid,
        password: psw1,
    };

    try {
        const response = await fetch("/user/psw", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(userData),
        });
    
        if (!response.ok) {
        //   throw new Error('Server error');
            messageElement.innerHTML = error.message;
        }
    
        window.location.href = "/v/users";
        } catch (error) {
        messageElement.innerHTML = "Change password failed: " + error.message;

    }
}

function validatePasswords(password, password2) {
    var messageElement = document.getElementById("_message");

    if (password === "" || password2 === "") {
        messageElement.innerHTML = "Passwords cannot be empty";
    return false;
    }

    if (password !== password2) {
        messageElement.innerHTML = "Passwords do not match";
    return false;
    }

    if (password.length < 8) {
    messageElement.innerHTML = "Passwords must be at least 8 characters long";
    return false;
    }

    return true;
}






