import React from "react";

export const Banner = () => {
  return (
    <div
      className="is-relative"
      style={{
        height: 300,
        backgroundImage: "url(/assets/img/bg.jpg)",
        backgroundSize: "cover",
        backgroundPosition: "center",
      }}
    >
      <div
        className="has-background-black-bis w-full h-full is-absolute"
        style={{ top: 0, opacity: 0.2 }}
      ></div>
    </div>
  );
};
