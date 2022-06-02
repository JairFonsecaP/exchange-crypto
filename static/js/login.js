const login = async () => {
  const data = {
    username: document.getElementById("username").value,
    password: document.getElementById("password").value,
  };

  const response = await fetch("/api/user/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });

  if (response.ok) {
    window.location.pathname = "/dashboard";
  } else {
    document.getElementById(
      "alert"
    ).innerHTML = `<div class="alert alert-danger alert-dismissible fade show" role="alert">
                                                    <strong>STOP!</strong> Check your username and/or password.
                                                    <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
                                                  </div>`;
  }
};

const validate = (e) => {
  const button = document.getElementById("button-login");
  if (
    document.getElementById("username").value === "" ||
    document.getElementById("password").value === ""
  ) {
    button.classList.add("disabled");
  } else {
    button.classList.remove("disabled");
    if (e.code === "Enter") {
      login();
    }
  }
};

document.getElementById("username").addEventListener("keyup", validate);
document.getElementById("password").addEventListener("keyup", validate);
