import React, { useState, useEffect } from "react";

import { postAPI, serverAddr, orderURL } from "../utils/Axios";
import IconButton from "@material-ui/core/IconButton";
import DeleteForeverIcon from "@material-ui/icons/DeleteForever";
import { MainButton } from "./styles/Buttons";

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
    console.log(JSON.stringify(cartProducts));
    if (order === true) {
      postAPI(serverAddr + orderURL, JSON.stringify(cartProducts), addToCart);
      setOrder(false);
    }
  }, [order, cartProducts, addToCart]);

  console.log(cartProducts);
  return (
    <>
      <ul>
        <hr style={{ color: "lightgrey", height: "0.1px" }} />
        {cartProducts.error ? (
          <p>{cartProducts.error}</p>
        ) : (
          cartProducts.map((productInfo) => {
            const { product, quantity } = productInfo;
            const { picture, availability, name, id } = product;
            return (
              <li key={id}>
                <div
                  style={{
                    display: "flex",
                    width: "100%",
                    justifyContent: "space-around",
                    alignContent: "space-around",
                  }}
                >
                  <img src={picture} style={{ width: "100px" }} alt={name}></img>
                  <h3 style={{ color: "black", margin: "auto", width: "300px" }}>
                    {name}
                  </h3>
                  <p style={{ color: "black", margin: "auto" }}>Quantity: {quantity}</p>
                  <div style={{ textAlign: "center", margin: "auto" }}>
                    <IconButton aria-label="delete" onClick={() => onDelete(id)}>
                      <DeleteForeverIcon fontSize="large" style={{ color: "red" }} />
                    </IconButton>
                  </div>
                </div>
                <hr style={{ color: "lightgrey", height: "0.1px" }} />
              </li>
            );
          })
        )}
      </ul>
      <div style={{ width: "100%", textAlign: "center" }}>
        <MainButton onClick={() => setOrder(true)}>Rent</MainButton>
      </div>
    </>
  );
}

export default Cart;
