const logout = () => {
  document.cookie = "token=x;expires=1 Mar 1990 00:00:00 GMT";
  window.location.pathname = "/";
};
