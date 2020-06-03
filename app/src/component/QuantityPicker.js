import React, { useState } from "react";
import styled from "styled-components";

import ButtonGroup from "@material-ui/core/ButtonGroup";
import { CounterButton } from "./styles/Buttons";

const CounterValue = styled.div`
  border-radius: 0px;
  border: 1px solid #d09237;
  padding-top: 5px;
`;

function QuantityPicker({ product, addToCart, defaultQuantity = 0 }) {
  const [quantity, setQuantity] = useState(defaultQuantity);

  const maxQuantity = product.availability;

  const handleIncrement = () => {
    if (quantity < maxQuantity) setQuantity(quantity + 1);
  };
  const handleDecrement = () => {
    if (quantity > 0) setQuantity(quantity - 1);
  };

  const ValidateProduct = () => {
    if (quantity > 0) addToCart(product, quantity);
    setQuantity(0);
  };

  return (
    <div style={{ height: "35px" }}>
      <ButtonGroup size="small" aria-label="small outlined button group">
        <CounterButton onClick={() => handleDecrement()}>-</CounterButton>
        <CounterValue>{quantity}</CounterValue>
        <CounterButton onClick={() => handleIncrement()}>+</CounterButton>
        <CounterButton
          style={{
            backgroundColor: "#d09237",
            color: "white",
            fontSize: "18px",
            height: "100%",
          }}
          onClick={() => ValidateProduct()}
        >
          Validate
        </CounterButton>
      </ButtonGroup>
    </div>
  );
}

export default QuantityPicker;
