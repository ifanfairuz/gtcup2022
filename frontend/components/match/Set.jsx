import React from "react";
import { Crown } from "@components/icons/Crown";
import { Ball } from "@components/icons/Ball";

export const Set = ({ data }) => {
  return (
    <div className="columns is-mobile is-vcentered is-centered is-gapless mb-0">
      <p className="column has-text-left">
        <Ball className="is-size-7 has-text-primary mr-1" />
        <span>{data.Home}</span>
        {data.winner == "home" && (
          <Crown className="is-size-7 has-text-warning ml-1" />
        )}
      </p>
      <p className="column has-text-centered">Set {data.Key}</p>
      <p className="column has-text-right">
        <Ball className="is-size-7 has-text-info mr-1" />
        <span>{data.Away}</span>
        {data.winner == "away" && (
          <Crown className="is-size-7 has-text-warning ml-1" />
        )}
      </p>
    </div>
  );
};
