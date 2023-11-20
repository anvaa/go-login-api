async function loginClick() {

  const email = document.getElementById("_email").value;
  const password = document.getElementById("_password").value;
  const messageElement = document.getElementById("_message");

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
};

function validatePassword(psw) {
  return psw.length >= 8;
}