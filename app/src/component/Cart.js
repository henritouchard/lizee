import React, { useState, useEffect } from "react";

import { postAPI, serverAddr, orderURL } from "../utils/Axios";
import CartProduct from "./CartProduct";
import { ButtonPrimary } from "./styles/Buttons";

function Cart({ addToCart, cartProducts }) {
  const [order, setOrder] = useState(false);

  // Delete cart item
  const onDelete = (itemID) => {
    const newCart = cartProducts.filter((productInfo) => {
      const { product } = productInfo;
      return product.id !== itemID;
    });
    addToCart(newCart);
  };

  useEffect(() => {
    if (order === true) {
      const emptyCart = () => addToCart([]);
      const displayProductsError = (error) => {
        if (error && error.product) {
          // Find corresponding products and insert errors into
          let cartCopy = cartProducts;
          cartCopy.forEach((p) => {
            error.product.forEach((errorProduct) => {
              const { product_id, availability } = errorProduct;
              if (p.product.id === product_id) {
                // Insert error object in product
                p.product.availability = availability;
              }
            });
          });
          console.log(cartCopy);
          // Reinitialize cart with error data
          addToCart(cartCopy);
        }
      };
      postAPI(
        serverAddr + orderURL,
        JSON.stringify(cartProducts),
        emptyCart,
        displayProductsError
      );
      setOrder(false);
    }
  }, [order, cartProducts, addToCart]);

  let error = false;
  console.log(cartProducts);
  cartProducts.forEach((element) => {
    if (element.quantity > element.product.availability) error = true;
  });

  return (
    <>
      <ul>
        <hr style={{ color: "lightgrey", height: "0.1px" }} />
        <CartProduct products={cartProducts} onDelete={onDelete} />
      </ul>
      <div style={{ width: "100%", textAlign: "center" }}>
        <ButtonPrimary disabled={error} onClick={() => setOrder(true)}>
          Rent
        </ButtonPrimary>
      </div>
    </>
  );
}

export default Cart;
