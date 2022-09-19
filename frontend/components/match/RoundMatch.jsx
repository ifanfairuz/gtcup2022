import React from "react";
import { Match } from "./Match";
// import { ShareButtons } from "@components/share/ShareButtons";

export const RoundMatch = ({ title, datas, ...props }) => {
  return (
    <div {...props}>
      <div className="py-2 px-4 has-background-dark has-text-white-ter has-text-weight-semibold">
        {title}
      </div>
      {Object.values(datas).map(({ title, datas }) => (
        <div key={title}>
          <div className="py-1 px-4 has-background-grey-lighter has-text-weight-medium">
            {title}
          </div>
          {/* <div className="py-1 px-4 has-background-white-ter">
            <ShareButtons size="small" />
          </div> */}
          <div className="columns is-gapless has-border mb-0">
            {datas.map((d) => (
              <div key={d.ID} id={`match-${d.ID}`} className="column">
                <Match data={d} />
              </div>
            ))}
          </div>
        </div>
      ))}
    </div>
  );
};
