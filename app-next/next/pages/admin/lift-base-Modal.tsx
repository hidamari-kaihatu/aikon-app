import React from "react";
import { useState } from "react";
import Modal from "../../components/Modal";
import CenterLiftPanel from "../../components/CenterLiftPanel";

export default function Lift() {
  const [isOpenModal, setIsOpenModal] = useState(false);

  const toggleModal = (e:any) => {
    if (e.target === e.currentTarget) {
      setIsOpenModal(!isOpenModal);
    }
  };

  return (
    <div className="App">
      <button type="button" onClick={toggleModal}>
        解除
      </button>
      {isOpenModal && (
        <Modal close={toggleModal}>
          <CenterLiftPanel />
        </Modal>
      )}
    </div>
  );
}