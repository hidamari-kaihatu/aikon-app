import ReactDOM from "react-dom";

const Portal = ({ children }) => {
  const element = document.querySelector("#__next");
  return element ? ReactDOM.createPortal(children, element) : null;
};

export default Portal;