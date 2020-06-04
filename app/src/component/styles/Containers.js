import React from "react";

import styled from "styled-components";
import Tab from "@material-ui/core/Tab";
import Typography from "@material-ui/core/Typography";
import Box from "@material-ui/core/Box";

import { Colors } from "./Theme";

export const ProductCard = styled.div`
  border-radius: 2px;
  height: auto;
  width: 100%;
  text-align: center;
  border: 1px solid rgba(0, 0, 0, 0.05);
  border-image: initial;
  box-shadow: rgba(0, 0, 0, 0.12) 0px 5px 10px;
  transition: box-shadow 0.2s ease 0s;
  padding: 20px;
  &:hover {
    box-shadow: rgba(0, 0, 0, 0.12) 0px 30px 60px;
  }
  h3 {
    color: black;
  }
  @media (max-width: 1201px) {
    margin-left: -25px;
  }
`;

export const FlexContainer = styled.ul`
  display: flex;
  flex-flow: wrap;
  justify-content: space-between;
  align-content: space-around;
  padding: 30px 10%;
  &::after {
    content: "";
    flex-basis: 33.3%;
    max-width: 30%;
    @media (min-width: 800px) and (max-width: 1200px) {
      flex-basis: 45%;
      max-width: 45%;
    }
  }
  li {
    flex-basis: 100%;
    max-width: 100%;
    margin-bottom: 30px;
    @media (min-width: 800px) and (max-width: 1200px) {
      flex-basis: 45%;
      max-width: 45%;
    }
    @media (min-width: 1201px) {
      flex-basis: 33.3%;
      max-width: 30%;
    }
  }
`;

export function TabPanel(props) {
  const { children, value, index, ...other } = props;

  return (
    <div
      role="tabpanel"
      hidden={value !== index}
      id={`full-width-tabpanel-${index}`}
      aria-labelledby={`full-width-tab-${index}`}
      {...other}
    >
      {value === index && (
        <Box p={3}>
          <Typography component={"div"}>{children}</Typography>
        </Box>
      )}
    </div>
  );
}

export const StyledTab = styled(Tab)`
  background-color: ${({ active }) =>
    active ? Colors.primary : Colors.secondary}!important;
  color: ${({ active, disabled }) =>
    active ? Colors.secondary : !disabled && Colors.primary} !important;
  font-weight: 700 !important;
`;
