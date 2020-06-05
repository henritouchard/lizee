import styled from "styled-components";

import Button from "@material-ui/core/Button";

import { Colors } from "./Theme";

export const CounterButton = styled(Button)`
  border-radius: 0px !important;
  border-color: ${Colors.primary} !important;
`;

export const ButtonPrimary = styled(Button)`
  flex: 1;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  width: 25%;
  line-height: 22px;
  text-align: center;
  min-width: 250px !important;
  padding: 10px 10px 10px 0 !important;
  border-radius: 0px !important;
  background-color: ${({ disabled }) =>
    disabled ? Colors.grey : Colors.primary} !important;
  color: ${Colors.secondary} !important;
  font-weight: 900 !important;
  font-size: 20px !important;
  /* cursor: default; */
`;
