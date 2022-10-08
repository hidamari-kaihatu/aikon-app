import React, { useState } from "react";
import SpeechRecognition from "react-speech-recognition";

function StartButton() {
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

  return (
    <div>
      <button onClick={clickHandler}>
        <span>{listen ? "Stop Listening" : "Start Listening"} 
        </span>
      </button>
    </div>
  );
}

export default StartButton;