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
    window.location.pathname = "/dashboard";
    document.cookie = `token=${res.token};expires=${new Date().setTime(
      new Date().getTime() * 24 * 60 * 60 * 1000
    )};SameSite=None;Secure`;
  } else {
    alert("Bad credentials");
  }
};
