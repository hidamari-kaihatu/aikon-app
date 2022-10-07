import React, { useState } from "react";
//import StartButton from "../../components/startButton";
import Output from "../../components/output";
import Layout from "../teachers/teacher-layout";

function App() {
  return (
    <>
    <Layout>
      <h1>先生が保護者に音声でメッセージを送るページ</h1>
      <div>
{/*       <StartButton />  */}
      <Output /> 
      </div>
      </Layout>
    </>
  );
}
  
export default App;

//messageを送る人は、送る人のcenter_IDをぶちこむ