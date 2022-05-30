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
  if (response.status == 200) {
    const data = response.json();
    sessionStorage.setItem("token", token);
    fetch("/dahsboard");
    // window.location.pathname = "/dashboard";
  }
  alert("Bad credentials");
};
