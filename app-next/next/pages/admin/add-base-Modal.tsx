import React from "react";
import { useState } from "react";
import Modal from "../../components/Modal";
import CenterAddPanel from "../../components/centerAddPanel";

export default function Add() {
  const [isOpenModal, setIsOpenModal] = useState(false);

  const toggleModal = (e:any) => {
    if (e.target === e.currentTarget) {
      setIsOpenModal(!isOpenModal);
    }
  };

  return (
    <div className="App">
      <button type="button" onClick={toggleModal}>
        新規施設登録
      </button>
      {isOpenModal && (
        <Modal close={toggleModal}>
          <CenterAddPanel />
        </Modal>
      )}
    </div>
  );
}