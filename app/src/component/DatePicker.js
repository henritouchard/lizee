import React, { useState } from "react";

import DateFnsUtils from "@date-io/date-fns";
import { MuiPickersUtilsProvider, KeyboardDatePicker } from "@material-ui/pickers";

function DatePicker({ label, innerLabel, setChoice, choice }) {
  const [selectedDate, setSelectedDate] = useState(new Date());

  const handleDateChange = (date) => {
    let c = choice;
    c[innerLabel] = date.toISOString().substr(0, 10);
    setChoice(c);
    setSelectedDate(date);
  };

  return (
    <div style={{ width: "25%", flex: 1, padding: "10px 10px 0" }}>
      <MuiPickersUtilsProvider utils={DateFnsUtils}>
        <KeyboardDatePicker
          label={label}
          format="yyyy/MM/dd"
          value={selectedDate}
          autoOk
          onChange={handleDateChange}
          KeyboardButtonProps={{
            "aria-label": "change date",
          }}
          style={{ width: "100%", minWidth: "250px" }}
        />
      </MuiPickersUtilsProvider>
    </div>
  );
}

export default DatePicker;
