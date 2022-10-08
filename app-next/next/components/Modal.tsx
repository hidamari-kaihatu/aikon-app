import React, { useState } from "react";
import Portal from "./Portal";

type Props = {
  close: (e: any) => void;
  children: React.ReactNode;
};

const Modal: React.FC<Props> = props => {
  const [isMouseDown, setIsMouseDown] = useState(false);

  const onMouseDown = (e:any) => {
    if (e.target === e.currentTarget) {
      setIsMouseDown(true);
    }
  };

  const onMouseUp = (e:any) => {
    if (isMouseDown) {
      props.close(e);
    }
    setIsMouseDown(false);
  };

  return (
    <Portal>
      <div
        onMouseDown={onMouseDown}
        onMouseUp={onMouseUp}
      >
        <div>
          {React.cloneElement(props.children as any, {
            close: props.close
          })}
        </div>
      </div>
    </Portal>
  );
};

export default Modal;