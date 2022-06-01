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
    const res = await response.json();
    document.cookie = `token=${res.token};expires=${new Date().setTime(
      new Date().getTime() * 24 * 60 * 60 * 1000
    )};SameSite=None;Secure`;
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
