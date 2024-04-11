
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
    
    messageElement.style.border = "1px solid red";
  
    if (!validateEmail(email)) {
      messageElement.innerHTML = "Not valid username";
      return;
    }
  
    if (!validatePassword(password)) {
      messageElement.innerHTML = "Password must be at least 8 characters long";
      return;
    }
  
    const data = {
      email: email,
      password: password
    };
    // const data = { email, password };
  
    const response = await fetch("/login", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data),
    });
  
    const responseData = await response.json();
  
    if (response.status === 200) {
      window.location.href = responseData.message;
    } else {
      messageElement.innerHTML = responseData.message;
    }
    
  };
  
  function validatePassword(psw) {
    return psw.length >= 8;
  }
  
  function validateEmail(email) {
    return email.length >= 8;
  }