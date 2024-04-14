
const messageElement = document.getElementById("_message");

document.getElementById("_password").addEventListener("keydown", function (event) {
  if (event.key === "Enter") {
    loginClick();
  }
}
);

document.getElementById("_email").addEventListener("keydown", function (event) {
  if (event.key === "Enter") {
    loginClick();
  }
}
);


async function loginClick() {

  const email = document.getElementById("_email").value;
  const password = document.getElementById("_password").value;

  if (!validateEmail(email)) {
    messageElement.innerHTML = "Not a valid email";
    return;
  }

  if (!validatePassword(password)) {
    messageElement.innerHTML = "Password must be at least 8 characters long";
    messageElement.style.border = "1px solid red";
    return;
  }

  const data = {
    email: email,
    password: password
  };

  try {
    const response = await fetch("/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data),
    });

    const responseData = await response.json();

    if (!response.ok) {
      throw new Error('Server error');
    }

    window.location.href = responseData.url;
  } catch (error) {
    messageElement.innerHTML = "Login failed: " + error.message;
    messageElement.style.border = "1px solid red";
  }
  
};

function validatePassword(psw) {
  return psw.length >= 8;
}

function validateEmail(email) {
  return email.length >= 8;
}