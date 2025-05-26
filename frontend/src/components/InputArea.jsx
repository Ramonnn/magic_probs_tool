import React from "react";

const InputArea = ({ value, onChange, disabled }) => {
  return (
    <textarea
      className="w-full p-3 border rounded h-40 resize-y"
      placeholder="Enter card names, one per line"
      value={value}
      onChange={onChange}
      disabled={disabled}
    />
  );
};


export default InputArea;

