import React from "react";

const ErrorMessage = ({ message }) => {
  if (!message) return null;

  return <p className="text-red-600">{message}</p>;
};


export default ErrorMessage;

