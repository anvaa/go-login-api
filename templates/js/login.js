document.getElementById("login-form").addEventListener("submit", async function (event) {
  event.preventDefault();

  const email = document.getElementById("email").value;
  const password = document.getElementById("password").value;
  const messageElement = document.getElementById("message");

  if (!validateEmail(email)) {
    messageElement.innerHTML = "Please enter a valid email address";
    return;
  }

  if (!validatePassword(password)) {
    messageElement.innerHTML = "Password must be at least 8 characters long";
    return;
  }

  try {
    const userData = { email, password };
    const response = await fetch("/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(userData),
    });

    if (!response.ok) {
      throw new Error('Server error');
    }

    window.location.href = "/v/home";
  } catch (error) {
    messageElement.innerHTML = "Login failed: " + error.message;
  }
});

function validateEmail(email) {
  return /\S+@\S+\.\S+/.test(email);
}

function validatePassword(password) {
  return password.length >= 8;
}