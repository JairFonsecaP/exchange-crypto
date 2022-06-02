const signUp = () => {
  const data = {
    name: document.getElementById("name").value,
    email: document.getElementById("email").value,
    username: document.getElementById("username").value,
    password: document.getElementById("password").value,
    repeatPassword: document.getElementById("repeatPassword").value,
  };
  if (!validation(data)) {
    return;
  }
  registration(data);
};

const registration = async (data) => {
  const response = await fetch("/api/user/signup", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  });
  if (response.ok) {
    showAlert(
      "success",
      "<strong>Congratulations!!</strong> User created, you can login now."
    );
    setTimeout(() => (window.location.pathname = "/login"), 3000);
  }
};

const validation = (data) => {
  if (
    data.name === "" ||
    data.username === "" ||
    data.username === "" ||
    data.password === ""
  ) {
    showAlert("danger", "All fields must be completed");
    return false;
  }

  if (data.password !== data.repeatPassword) {
    showAlert("danger", "Password and repeat password must be same");
    document.getElementById("repeatPassword").value = "";
    document.getElementById("password").value = "";
    return false;
  }

  if (data.password.length < 8) {
    showAlert("danger", "Password must be at least 8 characters long");
    return false;
  }
  return true;
};

const showAlert = (type, message) => {
  document.getElementById(
    "alert"
  ).innerHTML = `<div class="alert alert-${type} alert-dismissible fade show" role="alert">
                            <div>
                            ${message}
                            </div>
                            <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
                        </div>`;
};
