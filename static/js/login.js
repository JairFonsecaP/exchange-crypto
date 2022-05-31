const login = async () => {
  const data = {
    username: document.getElementById("username").value,
    password: document.getElementById("password").value,
  };

  fetch("/api/user/login", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(data),
  })
    .then((response) => {
      if (response.ok) {
        return response.json();
      } else {
        alert("Bad credentials");
      }
    })
    .then((res) => {
      window.location.pathname = "/dashboard";
      document.cookie = `token=${res.token};expires=${new Date().setTime(
        new Date().getTime() * 24 * 60 * 60 * 1000
      )};SameSite=None;Secure`;
    });
};
