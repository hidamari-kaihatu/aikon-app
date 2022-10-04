import React, { useState } from "react";
import StartButton from "../../components/startButton";
import Output from "../../components/output";
import Layout from "../teachers/teacher-layout";

function App() {
  const[fromPerson, setFromPerson]=useState("")
  const[toPerson, setToPerson]=useState("")
  const submit = () => {
    window.location.reload()
  }
  
  return (
    <>
    <Layout>
      <h1>先生が保護者に音声でメッセージを送るページ</h1>
      <div>
      <label>誰から送る: </label>
      <select value={fromPerson} onChange={(e) => setFromPerson(e.target.value)}>
        <option value="A">先生の名前</option>
        <option value={"高井日菜子"}>高井日菜子</option>
      </select>
      <br></br>
      <label>誰へ送る: </label>
      <select value={toPerson} onChange={(e) => setToPerson(e.target.value)}>
        <option value="A">生徒の名前</option>
        <option value={"HINAKO TAKAI"}>HINAKO TAKAI</option>
      </select>
      <StartButton />
      <Output /> 
      </div>
      <button onClick={submit}>送信する</button> 
      </Layout>
    </>
  );
}
  
export default App;