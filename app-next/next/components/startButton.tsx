import React, { useState } from "react";
import SpeechRecognition from "react-speech-recognition";

function StartButton() {
  const [listen, setListen] = useState(false);
  let link;
  let attend;
  let download;

  const clickHandler = () => {
    if (listen === false) {
      SpeechRecognition.startListening({ continuous: true });
      setListen(true);
      navigator.mediaDevices.getUserMedia({
        video: false,
        audio: true,
      }).then(handleSuccess) 
    } else {
      SpeechRecognition.abortListening();
      setListen(false);
    }
  };

  const handleSuccess = function(stream:any) {
    let recordedChunks:any = [];
    const mediaRecorder = new MediaRecorder(stream);

    mediaRecorder.start();
    mediaRecorder.onstop = function() {
      let d = new Date();
      let fn = ((((d.getFullYear()*100 + d.getMonth()+1)*100 + d.getDate())*100
                + d.getHours())*100 + d.getMinutes())*100 + d.getMinutes();
      console.log(12)
      link = URL.createObjectURL(new Blob(recordedChunks));
      attend = "録音ファイルはこちら";
      download = fn+".webm"; 
    };

    mediaRecorder.ondataavailable = function(e) {
      console.log(13)
      if (e.data.size > 0) {
        recordedChunks.push(e.data);
    }
    };
  };


  return (
    <>
    <div>
      <button onClick={clickHandler}>
        <span>{listen ? "音声入力ストップ" : "音声入力スタート"} 
        </span>
      </button>
      <a href={link} download={download}>{attend}</a>
    </div>
    </>
  );
}

export default StartButton; 



