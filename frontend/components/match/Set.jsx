import React from "react";
import { Crown } from "@components/icons/Crown";

export const Set = () => {
  return (
    <div className="columns is-mobile is-vcentered is-centered is-gapless mb-0">
      <p className="column has-text-left">
        25
        {/* <Crown className="has-text-warning" /> */}
      </p>
      <p className="column has-text-centered">Set 1</p>
      <p className="column has-text-right">19</p>
    </div>
  );
};
