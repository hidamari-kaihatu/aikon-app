import React, { useState } from "react";
import { useSpeechRecognition } from "react-speech-recognition";
import { Axios } from "../lib/api";
import SpeechRecognition from "react-speech-recognition";

function Output() {
  const [listen, setListen] = useState(false);

  const clickHandler = () => {
    if (listen === false) {
      SpeechRecognition.startListening({ continuous: true });
      setListen(true);
      // The default value for continuous is false, meaning that
      // when the user stops talking, speech recognition will end. 
    } else {
      SpeechRecognition.abortListening();
      setListen(false);
    }
  };

  const[toPerson, settoPerson]=useState("")
  const[message, setmessage]=useState("")

  //toPersonは、ログインしている先生が所属する学童の子供たちのstudent_idが入る
  //そのため、studentテーブルのすべての子供たちから、ログインしている先生の学童のcenter_idと一致する子供たちをfilterかける
  //その後、returnの中のselect文で選ばれた子供のstudent_idがtoPersonの中に入る

  const [outputMessage, setOutputMessage] = useState("");
  const { transcript, resetTranscript } = useSpeechRecognition();

  //Staff_idは、ログインしている先生が所属する学童のcenter_idにする。現状「１」でべたうち
  //Datetimeを直す
  //入力できない。音声のみ
  //reloadができてない

  const submit = () => {
    const data = {
      "Staff_id":1,
      "Message":transcript,
      "Datetime":"2022-10-01 12:12:12",
      "Student_id":1,
      "voice":"https://voice.com"
  }
  Axios.post(`api/proxy/teacherMessagePost`, data)
  .then((res) => {
    console.log(res);
  })
  .catch((error) => {
    console.log(error);
  });
  }

  //現状、toPersonの値はデフォルトで１だけ
  return (
    <div>
      <label>誰へ送る: </label>
      <select value={toPerson} onChange={(e) => settoPerson(e.target.value)}>
        <option value="A">生徒の名前</option>
        <option value={1}>HINAKO TAKAI</option>
      </select>
      <br></br>
      <button onClick={clickHandler}>
        <span>{listen ? "Stop Listening" : "Start Listening"} 
        </span>
      </button>
      <p>{transcript}</p>
      <button onClick={submit}>送信する!</button> 
    </div>
  );
}

export default Output;