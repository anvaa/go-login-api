async function signUp() {
  
  const email = document.getElementById("email").value;
  const password = document.getElementById("password").value;
  const password2 = document.getElementById("password2").value;
  const messageElement = document.getElementById("_message");
  
  if (!validateEmail(email)) {
    messageElement.innerHTML = "Please enter a valid email address";
    return;
  }

  if (!validatePasswords(password, password2)) {
    return; // Message is set inside the validatePasswords function
  }

  try {
    const userData = { email, password };
    
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

  
function validateEmail(email) {
  return /\S+@\S+\.\S+/.test(email);
}
  
function validatePasswords(password, password2) {
  if (password === "" || password2 === "") {
    document.getElementById("message").innerHTML = "Passwords cannot be empty";
    return false;
  }

  if (password !== password2) {
    document.getElementById("message").innerHTML = "Passwords do not match";
    return false;
  }

  if (password.length < 8) {
    document.getElementById("message").innerHTML = "Passwords must be at least 8 characters long";
    return false;
  }

  return true;
}
