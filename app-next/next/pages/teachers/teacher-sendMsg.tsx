// import React, { useState } from "react";
// //import StartButton from "../../components/startButton";
// // import Output from "../../components/output";
// import Layout from "../teachers/teacher-layout";

// function App() {
//   return (
//     <>
//     <Layout>
//       <h1>先生が保護者に音声でメッセージを送るページ</h1>
//       <div>
// {/*       <StartButton />  */}
//       {/* <Output />  */}
//       </div>
//       </Layout>
//     </>
//   );
// }
  
// export default App;

//messageを送る人は、送る人のcenter_IDをぶちこむ
//outputコンポーネントからコピーして
import React, { useState } from "react";
import { useSpeechRecognition } from "react-speech-recognition";
import { Axios } from "../../lib/api";
import SpeechRecognition from "react-speech-recognition";
import axios from "axios";
import Layout from "../teachers/teacher-layout";


function Output({data, staffs}:any) {
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
  const stuId = Number(staffs[0].Id)
  // console.log(typeof stuId)
  const submit = () => {
    const data = {
      "Staff_id":stuId,
      "Message":transcript,
      "Student_id":Number(toPerson),
  }
  Axios.post(`api/proxy/teacherMessagePost`, data)
  .then((res) => {
    console.log(res);
  })
  .catch((error) => {
    console.log(error);
  });
  }


  const infos = data.filter(obj =>obj.Center_id  === staffs[0].Center_id)//idでふるいにかける
  // console.log(data)
  // console.log(infos)
  // console.log(staffs)
  // console.log(infos[0].Student_name)
  //現状、toPersonの値はデフォルトで１だけ
  return (
    <>
    <Layout>
      <h1>先生が保護者に音声でメッセージを送るページ</h1>
    <div className='center'>
      <label>誰へ送る: </label>
      <select value={toPerson} onChange={(e) => settoPerson(e.target.value)}>
        <option value="A">生徒の名前</option>
        {infos.map(((info:any, i: number) => {
          return (
            <>
            <option value={info.Id} key={i}>{info.Name}</option>
            </>
          )
        }))}
      </select>
      <br></br>
      <button onClick={clickHandler}>
        <span>{listen ? "Stop Listening" : "Start Listening"} 
        </span>
      </button>
      <p>{transcript}</p>
      <button onClick={submit}>送信する!</button> 
    </div>
      {staffs[0].CenterName}
        <br></br>
      {staffs[0].Name}
    </Layout>
    </>
  );
}


export async function getServerSideProps() {
  const res = await axios.get(`${process.env.API}/getAllStudents`, {
  });
  const data = await res.data;

  const staffsRes = await axios.get(`${process.env.API}/getStaffAndMiddleAndCenter`, {
  });
  const staffs = await staffsRes.data;

  return { 
      props: {
        data: data,
        staffs: staffs
      },
  };
}


export default Output;