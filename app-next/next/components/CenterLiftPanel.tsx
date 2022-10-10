import React from "react";

type Props = {
  close?: (e: any) => void;
};

const CenterLiftPanel: React.FC<Props> = props => {
  const submit = (e:any) => {
    e.preventDefault();
    //モーダルの後に登録解除フォームへ遷移
    window.location.href='https://forms.gle/ZbDAzfQR48jhHMD59'
    //第二フェーズTODO:Stripeのサブスク解除の画面に飛ぶようにする
    if (props.close) {
      props.close(e);
    }
  };

  return (
    <section>
      <header>
        <p>施設の登録解除を行うには、専用フォームでの手続きが必要です</p>
      </header>
      <div>登録解除のお問い合わせフォームへ移動しますか？</div>
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

export default CenterLiftPanel;