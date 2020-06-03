import axios from "axios";

export const serverAddr = "http://localhost:3000/";
export const availableProductsURL = "categories/products?categoryID=";
export const categoryListURL = "categories/";
export const orderURL = "products/order/";

export async function fetchAPI(query, callback) {
  await axios
    .get(query)
    .then((r) => callback(r.data))
    .catch((error) => {
      callback({ error: "An Error occured please reload the page" });
    });
}

export async function postAPI(query, params, callback) {
  await axios
    .post(query, params)
    .then((r) => callback([]))
    .catch((error) => {
      callback({ error: "An Error occured please reload the page" });
    });
}
