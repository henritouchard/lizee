import React from "react";
import IconButton from "@material-ui/core/IconButton";
import DeleteForeverIcon from "@material-ui/icons/DeleteForever";

import { Colors } from "./styles/Theme";

function CartProduct({ products, onDelete }) {
  return products.map((productInfo) => {
    const { product, quantity } = productInfo;
    const { picture, name, id, availability } = product;
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
          <h3 style={{ color: "black", margin: "auto", width: "300px" }}>{name}</h3>
          <div style={{ margin: "auto" }}>
            {availability < quantity ? (
              <p style={{ color: Colors.warning }}>
                You asked {quantity} but only {availability} are available.
              </p>
            ) : (
              <p style={{ color: "black" }}>Quantity: {quantity}</p>
            )}
          </div>
          <div style={{ textAlign: "center", margin: "auto" }}>
            <IconButton aria-label="delete" onClick={() => onDelete(id)}>
              <DeleteForeverIcon fontSize="large" style={{ color: Colors.warning }} />
            </IconButton>
          </div>
        </div>
        <hr style={{ color: "lightgrey", height: "0.1px" }} />
      </li>
    );
  });
}

export default CartProduct;
