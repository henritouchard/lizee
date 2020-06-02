import React from "react";
import styled from "styled-components";

import Grid from "@material-ui/core/Grid";
import Button from "@material-ui/core/Button";

import SelectInput from "./SelectInput";
import DatePicker from "./DatePicker";

const GridContainer = styled.div`
  margin: auto;
  position: relative;
  top: 80%;
  width: 90%;
  @media (max-width: 800px) {
    top: 70%;
  }
`;

const StyledGrid = styled(Grid)`
  background-color: white;
  align-content: space-around;
  justify-content: space-around;
  box-shadow: 5px 5px 5px rgb(0, 0, 0, 0.5);
`;

function HeroForm({ categories, setChoice, choice, validate }) {
  return (
    <GridContainer>
      <StyledGrid container justify="space-between">
        <SelectInput categories={categories} setChoice={setChoice} choice={choice} />
        <DatePicker
          label="Start Date"
          innerLabel="fromDate"
          setChoice={setChoice}
          choice={choice}
        />
        <DatePicker
          label="End Date"
          innerLabel="toDate"
          setChoice={setChoice}
          choice={choice}
        />
        <Button
          onClick={() => validate()}
          style={{
            flex: 1,
            width: "25%",
            minWidth: "250px",
            padding: "10px 10px 0",
            borderRadius: "0px",
            backgroundColor: "#d09237",
            color: "white",
            fontWeight: "900",
            fontSize: "20px",
          }}
        >
          Start
        </Button>
      </StyledGrid>
    </GridContainer>
  );
}

export default HeroForm;
