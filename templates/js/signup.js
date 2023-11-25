
document.getElementById("_password1").addEventListener("keydown", function (event) {
  if (event.key === "Enter") {
    signupClick();
  }
}
);

document.getElementById("_password2").addEventListener("keydown", function (event) {
  if (event.key === "Enter") {
    signupClick();
  }
}
);

document.getElementById("_email").addEventListener("keydown", function (event) {
  if (event.key === "Enter") {
    signupClick();
  }
}
);

async function signupClick() {
  
  const email = document.getElementById("_email").value;
  const password = document.getElementById("_password1").value;
  const password2 = document.getElementById("_password2").value;
  const messageElement = document.getElementById("_message");

  if (!validatePasswords(password, password2)) {
    return; // Message is set inside the validatePasswords function
  }

  try {
    const userData = { email, password, password2 };
    
    const response = await fetch("/signup", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(userData),
    });

    if (!response.ok) {
      throw new Error('Server error');
    }

    window.location.href = "/login";
  } catch (error) {
    messageElement.innerHTML = "Signup failed: " + error.message;
  }
};
  
function validatePasswords(psw1, psw2) {

  const errmsg = document.getElementById("_message");

  if (psw1 === "" || psw2 === "") {
    errmsg.innerHTML = "Passwords cannot be empty";
    return false;
  }

  if (psw1 !== psw2) {
    errmsg.innerHTML = "Passwords do not match";
    return false;
  }

  if (psw1.length < 8) {
    errmsg.innerHTML = "Passwords must be at least 8 characters long";
    return false;
  }

  if (psw1.length > 50) {
    errmsg.innerHTML = "Passwords must be less than 50 characters long";
    return false;
  }

  return true;
}
