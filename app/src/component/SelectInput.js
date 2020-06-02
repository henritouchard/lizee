import React from "react";
import styled from "styled-components";
import InputLabel from "@material-ui/core/InputLabel";
import MenuItem from "@material-ui/core/MenuItem";
import FormControl from "@material-ui/core/FormControl";
import Select from "@material-ui/core/Select";

const renderMenuItem = (categories) => {
  if (categories !== undefined || categories.error !== undefined) {
    return categories.map((category, key) => (
      <MenuItem key={key} value={category.id}>
        {category.name}
      </MenuItem>
    ));
  }
};

const SelectContainer = styled.div`
  flex: 1;
  width: 25%;
  padding: 10px 10px !important;
`;

function SelectInput({ categories, setChoice, choice }) {
  const handleChange = (event) => {
    let c = choice;
    c.category = event.target.value;
    setChoice(c);
  };

  return (
    <SelectContainer>
      <FormControl style={{ width: "100%", minWidth: "250px" }}>
        <InputLabel>Item category</InputLabel>
        <Select onChange={handleChange}>{renderMenuItem(categories)}</Select>
      </FormControl>
    </SelectContainer>
  );
}

export default SelectInput;
