const setTitles = () => {
  const titles = ["#", "Coin", "Price $", "1 h", "24 h", "7 d", "1 y"];
  const tTitles = document.getElementById("table-titles");
  titles.forEach((title) => {
    tTitles.innerHTML += `<td>${title}</td>`;
  });
};

const getData = async () => {
  return await fetch(`/api/coins/list`).then((res) => res.json());
};

const getDataFilter = async (search) => {
  return await fetch(`/api/coins/filter?search=${search}`, {
    method: "GET",
  }).then((res) => res.json());
};

const printData = (data) => {
  const tBody = document.getElementById("table-body");
  tBody.innerHTML = ``;
  if (data === null) {
    tBody.innerHTML = `<div class="alert alert-danger" role="alert">
                           Couldn't find coins!
                       </div>`;
    return;
  }

  data.forEach((coin, i) => {
    tBody.innerHTML += `<tr>
                            <td>${i + 1}</td>
                            <td>
                                <img src="${coin.image}" 
                                    alt=${coin.name} 
                                    style="width:3%" 
                                    class="me-4 img-fluid"/>
                                <span>${coin.name}</span>
                                <span 
                                class="ms-3 text-muted">
                                ${coin.symbol.toUpperCase()}
                                </span>  
                            </td>                   
                            <td>${coin.current_price.toFixed(2)}</td>
                            <td class="${
                              coin.price_change_percentage_1h_in_currency > 0
                                ? "text-success"
                                : "text-danger"
                            }">${coin.price_change_percentage_1h_in_currency.toFixed(
      1
    )} %
                            </td>
                            <td class="${
                              coin.price_change_percentage_24h > 0
                                ? "text-success"
                                : "text-danger"
                            }">${coin.price_change_percentage_24h.toFixed(1)} %
                            </td>
                            <td class="${
                              coin.price_change_percentage_7d_in_currency > 0
                                ? "text-success"
                                : "text-danger"
                            }">${coin.price_change_percentage_7d_in_currency.toFixed(
      1
    )} %
                            </td>
                            <td class="${
                              coin.price_change_percentage_1y_in_currency > 0
                                ? "text-success"
                                : "text-danger"
                            }">${coin.price_change_percentage_1y_in_currency.toFixed(
      1
    )} %
                            </td>
                        </tr>`;
  });
};

const init = async () => {
  setTitles();
  const data = await getData();
  printData(data);
};

window.addEventListener("load", init);

const apiSearch = async (e) => {
  const { value } = document.getElementById("search");
  let data = [];
  if (value === "") {
    data = await getData();
  } else {
    data = await getDataFilter(value);
  }
  printData(data);
};

document.getElementById("search").addEventListener("keyup", apiSearch);
