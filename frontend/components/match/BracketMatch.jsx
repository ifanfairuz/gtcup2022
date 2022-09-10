import React, { useMemo } from "react";
import { Match } from "./Match";

export const BracketMatch = ({ title, datas, ...props }) => {
  const matches = useMemo(() => {
    const d = Object.values(datas);
    let res = [];
    for (let i = 0; i < d.length; i += 2) {
      res.push(d.slice(i, i + 2));
    }
    return res;
  }, [datas]);

  return (
    <div {...props}>
      <div className="py-2 px-4 has-background-dark has-text-white-ter has-text-weight-semibold">
        {title}
      </div>
      {matches.map((datas, i) => {
        return (
          <div key={i} className="columns is-gapless has-border mb-0">
            {datas.map(({ title, match }) => (
              <div key={match.ID} className="column">
                <div className="py-1 px-4 has-background-grey-lighter has-text-weight-medium">
                  {title}
                </div>
                <Match data={match} />
              </div>
            ))}
          </div>
        );
      })}
    </div>
  );
};
