import React from "react";
import styled from "styled-components";

import SelectInput from "./SelectInput";
import DatePicker from "./DatePicker";

import Grid from "@material-ui/core/Grid";
import { ButtonPrimary } from "./styles/Buttons";
import { Colors } from "./styles/Theme";

const Div = styled.div`
  margin: auto;
  position: relative;
  top: 80%;
  width: 90%;
  @media (max-width: 1200px) {
    top: 55%;
  }
`;

const StyledGrid = styled(Grid)`
  background-color: ${Colors.secondary};
  align-content: space-around;
  justify-content: space-around;
  box-shadow: 5px 5px 5px rgb(0, 0, 0, 0.5);
`;

function HeroForm({ categories, setChoice, choice, validate }) {
  return (
    <Div>
      <StyledGrid container justify="space-between">
        <SelectInput categories={categories} setChoice={setChoice} choice={choice} />
        <DatePicker
          label="Start Date"
          innerLabel="from"
          setChoice={setChoice}
          choice={choice}
        />
        <DatePicker
          label="End Date"
          innerLabel="to"
          setChoice={setChoice}
          choice={choice}
        />
        <ButtonPrimary onClick={() => validate()} as="a" href="#bottomSection">
          START
        </ButtonPrimary>
      </StyledGrid>
    </Div>
  );
}

export default HeroForm;
