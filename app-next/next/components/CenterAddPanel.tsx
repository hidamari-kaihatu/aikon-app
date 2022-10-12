import React from "react";

type Props = {
  close?: (e: any) => void;
};

const CenterAddPanel: React.FC<Props> = props => {
  const submit = (e:any) => {
    e.preventDefault();
    //モーダルの後に施設登録フォームへ遷移
    window.location.href='https://forms.gle/UWVFgSBWAFJT271M9'
    if (props.close) {
      props.close(e);
    }
  };

  return (
    <section>
      <header>
        <p>新しく施設を登録するには、専用フォームでの登録が必要です</p>
      </header>
      <div>スマートGAKUDO新規施設登録フォームへ移動しますか？</div>
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

export default CenterAddPanel;