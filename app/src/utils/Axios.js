import axios from "axios";

export const serverAddr = "http://localhost:5000/";
export const availableProductsURL = "categories/products?categoryID=";
export const categoryListURL = "categories/";
export const orderURL = "products/order/";

export async function fetchAPI(query, callback) {
  await axios
    .get(query)
    .then((r) => callback(r.data))
    .catch((error) => {
      callback({ error: "An Error occured please reload the page", details: error });
    });
}

export async function postAPI(query, params, successCallback, errorCallback) {
  await axios
    .post(query, params)
    .then((r) => successCallback([]))
    .catch((error) => {
      if (error.response && error.response.data) errorCallback(error.response.data);
      else alert(error);
    });
}
