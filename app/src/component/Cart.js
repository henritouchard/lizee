import React, { useState, useEffect } from "react";

import { postAPI, serverAddr, orderURL } from "../utils/Axios";
import CartProduct from "./CartProduct";
import { ButtonPrimary } from "./styles/Buttons";

function Cart({ addToCart, cartProducts }) {
  const [order, setOrder] = useState(false);
  const [cartError, setCartError] = useState(null);

  // Delete cart item
  const onDelete = (itemID, errorProduct) => {
    const toFilter = errorProduct ? cartError : cartProducts;
    const newCart = toFilter.filter((productInfo) => {
      const { product } = productInfo;
      return product.id !== itemID;
    });
    if (errorProduct) setCartError(newCart);
    else addToCart(newCart);
  };

  useEffect(() => {
    if (order === true) {
      const emptyCart = () => addToCart([]);
      const displayProductsError = (error) => {
        if (error && error.product) {
          setCartError(error.product);
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

  return (
    <>
      <ul>
        <hr style={{ color: "lightgrey", height: "0.1px" }} />
        {cartError ? (
          <>
            <h2>{cartError.error}</h2>
            <CartProduct withError products={cartError} onDelete={onDelete} />
          </>
        ) : (
          <CartProduct products={cartProducts} onDelete={onDelete} />
        )}
      </ul>
      <div style={{ width: "100%", textAlign: "center" }}>
        <ButtonPrimary onClick={() => setOrder(true)}>Rent</ButtonPrimary>
      </div>
    </>
  );
}

export default Cart;
