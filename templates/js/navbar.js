
document.getElementById("mngusers").addEventListener("click", function (event) {
    window.location.href = "/v/users";
});
document.getElementById("userhome").addEventListener("click", function (event) {
    window.location.href = "/v/home";
});

document.getElementById("logoutbtn").addEventListener("click", async (event) => {
    
    const confirmed = confirm("Are you sure you want to logout?");
    if (!confirmed) {
    return;
    }

    const response = await fetch("/u/logout", {
    method: "POST",
    headers: {
        "Content-Type": "application/json",
    },
    });

    if (response.status === 401) {
    window.location.href = "/";
    }

});
