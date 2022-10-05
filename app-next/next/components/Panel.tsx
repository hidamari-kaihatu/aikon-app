import React from "react";

type Props = {
  close?: (e: any) => void;
};

const Panel: React.FC<Props> = props => {
  const submit = (e:any) => {
    e.preventDefault();
    if (props.close) {
      props.close(e);
    }
  };

  return (
    <section>
      <header>
        <h3>Modal Panel</h3>
      </header>
      <div>実行しますか？</div>
      <footer>
        <button type="button" onClick={props.close}>
          いいえ
        </button>
        <button type="submit" onClick={submit}>
          はい
        </button>
      </footer>
    </section>
  );
};

export default Panel;