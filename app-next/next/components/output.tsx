import React, { useState } from "react";
import { useSpeechRecognition } from "react-speech-recognition";

function Output() {
  const [outputMessage, setOutputMessage] = useState("");
  const { transcript, resetTranscript } = useSpeechRecognition();

  return (
    <>
    <div>
{/*       <p>{transcript}</p> */}
<input type="text" value={transcript}/>
      <p>{outputMessage}</p> 
{/*       <input type="text" value={transcript} /> */}
    </div>
    </>
  );
}

export default Output;