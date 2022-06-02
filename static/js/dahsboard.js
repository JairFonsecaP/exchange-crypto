const logout = () => {
  document.cookie = "token=x;expires=1 Mar 1990 00:00:00 GMT";
  window.location.pathname = "/";
};

const sellSpot = async (id) => {
  const response = await fetch(`/api/sell/spot/`, {
    method: "PATCH",
    headers: { "Content-Type": "application/json" },
    body: { id: id },
  });
  if (response.ok) {
    init();
  }
};

const setTitlesSpot = () => {
  const titles = [
    "Coin",
    "Quantity",
    "Current",
    "Variation",
    "Purchase to",
    "Sell",
  ];
  const tTitles = document.getElementById("table-titles-spot");
  tTitles.innerHTML = "";
  titles.forEach((title) => {
    tTitles.innerHTML += `<td>${title}</td>`;
  });
};
const setTitlesEarn = () => {
  const titles = [
    "Coin",
    "Quantity",
    "Time of delivery",
    "Quantity earn",
    "Sell",
  ];
  const tTitles = document.getElementById("table-titles-earn");
  tTitles.innerHTML = "";
  titles.forEach((title) => {
    tTitles.innerHTML += `<td>${title}</td>`;
  });
};
const getData = async () => {
  return await fetch(`/api/wallet/get`, { method: "GET" }).then((res) =>
    res.json()
  );
};

const printDataSpot = (data) => {
  const tBody = document.getElementById("table-body-spot");
  tBody.innerHTML = ``;
  if (data === null) {
    tBody.innerHTML = `<div class="alert alert-danger" role="alert">
                           Couldn't find coins!
                       </div>`;
    return;
  }

  data.spotPocket.coins.forEach((coin) => {
    const Variation = (coin.current_price * 100) / coin.buyprice - 100;
    tBody.innerHTML += `<tr>                            
                            <td>
                                <span 
                                class="text-muted">
                                ${coin.symbol.toUpperCase()}
                                </span>  
                            </td>        
                            <td>${coin.quantity}</td>           
                            <td>${coin.current_price.toFixed(2)}</td>
                            <td class="${
                              Variation > 0 ? "text-success" : "text-danger"
                            }">${Variation.toFixed(2)} %
                            <td>${coin.buyprice.toFixed(2)}
                            <td><button class="btn btn-success" onclick="sellSpot('${
                              coin.id
                            }')">Sell</button></td>
                        </tr>`;
  });
};

const printDataEarn = (data) => {
  const tBody = document.getElementById("table-body-earn");
  tBody.innerHTML = ``;
  if (data === null) {
    tBody.innerHTML = `<div class="alert alert-danger" role="alert">
                           Couldn't find coins!
                       </div>`;
    return;
  }

  data.spotPocket.coins.forEach((coin) => {
    const Variation = (coin.current_price * 100) / coin.buyprice - 100;
    tBody.innerHTML += `<tr>                            
                            <td>
                                <span 
                                class="text-muted">
                                ${coin.symbol.toUpperCase()}
                                </span>  
                            </td>        
                            <td>${coin.quantity}</td>           
                            <td>${coin.current_price.toFixed(2)}</td>
                            <td class="${
                              Variation > 0 ? "text-success" : "text-danger"
                            }">${Variation.toFixed(2)} %
                            <td>${coin.buyprice.toFixed(2)}
                            <td><button class="btn btn-success" onclick="sellSpot('${
                              coin.id
                            }')">Sell</button></td>
                        </tr>`;
  });
};

const init = async () => {
  setTitlesSpot();
  setTitlesEarn();
  const data = await getData();
  printDataSpot(data);
  printDataEarn(data);
};

window.addEventListener("load", init);
