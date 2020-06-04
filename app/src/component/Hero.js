import React, { useState, useEffect } from "react";
import styled from "styled-components";

import { fetchAPI, serverAddr, categoryListURL } from "../utils/Axios";
import HeroForm from "../component/HeroForm";

import CircularProgress from "@material-ui/core/CircularProgress";

import { Colors } from "../component/styles/Theme";
import HeroImg from "../assets/mount.jpg";

const PageTitle = styled.h1`
  font-family: sans-serif;
  position: absolute;
  top: 15%;
  left: 10%;
  font-size: 5vw;
  font-weight: 900;
`;

const Background = styled.div`
  width: 100vw;
  height: 95vh;
  background: url(${({ bgImage }) => bgImage}) center;
  background-size: cover;
  margin-bottom: 50px;
`;

const FormPlaceHolder = styled.div`
  margin: auto;
  position: relative;
  top: 80%;
  width: 90%;
  height: 90px;
  background-color: ${Colors.secondary};
  text-align: center;
  box-shadow: 5px 5px 5px ${Colors.shadow};
  @media (max-width: 800px) {
    top: 70%;
  }
`;

function renderForm(categories, setChoice, choice, onValidateChoice) {
  if (categories === null || categories.error !== undefined) {
    return (
      <FormPlaceHolder>
        <div style={{ paddingTop: "25px" }}>
          <CircularProgress />
        </div>
      </FormPlaceHolder>
    );
  } else if (categories.error !== undefined) {
    return (
      <FormPlaceHolder>
        <h3
          style={{ display: "block", color: "red", margin: "auto", paddingTop: "35px" }}
        >
          {categories.error}
        </h3>
      </FormPlaceHolder>
    );
  } else {
    return (
      <HeroForm
        categories={categories}
        setChoice={setChoice}
        choice={choice}
        validate={onValidateChoice}
      />
    );
  }
}

function Hero({ setTrek }) {
  const [categories, setCategories] = useState(null);
  const [choice, setChoice] = useState({
    category: null,
    from: new Date().toISOString().substr(0, 10),
    to: new Date().toISOString().substr(0, 10),
  });

  // Asynchronous call to api to get existing categories
  useEffect(() => {
    fetchAPI(serverAddr + categoryListURL, setCategories);
  }, []);

  const onValidateChoice = () => {
    if (
      choice != null &&
      choice.category !== null &&
      choice.from !== null &&
      choice.to !== null
    ) {
      setTrek((oldState) => [choice]);
    }
  };

  return (
    <Background bgImage={HeroImg}>
      <PageTitle>Adrenaline Heroes</PageTitle>
      {renderForm(categories, setChoice, choice, onValidateChoice)}
    </Background>
  );
}

export default Hero;
