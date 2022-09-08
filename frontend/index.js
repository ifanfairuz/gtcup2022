import React from "react";
import { renderToRoot } from "@support/render";

const App = () => {
  return (
    <div>
      <div style={{
        margin:"10px",
        padding:"10px",
        textAlign:"center",
        backgroundColor:"greenyellow"
      }}>
        <h1>HMMM</h1>
      </div>
    </div>)
}

renderToRoot(<App />)