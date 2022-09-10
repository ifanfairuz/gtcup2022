import React from "react";

export const Klasemen = ({ title, datas }) => {
  return (
    <div>
      <div className="py-2 px-4 has-background-dark has-text-white-ter has-text-weight-semibold">
        {title}
      </div>
      <div className="table-container">
        <table className="table is-striped is-fullwidth">
          <thead>
            <tr>
              <th width="70px" title="Position">
                Pos
              </th>
              <th title="Tim" style={{ minWidth: "200px" }}>
                Nama Tim
              </th>
              <th width="70px" title="Pertandingan">
                P
              </th>
              <th width="70px" title="Kemenangan">
                M
              </th>
              <th width="70px" title="Kekalahan">
                K
              </th>
              <th width="70px" title="Set Kemenangan">
                SM
              </th>
              <th width="70px" title="Set Kekalahan">
                SK
              </th>
              <th width="70px" title="Agregat Set">
                AS
              </th>
              <th width="70px" title="Skor Kemenangan">
                SC-M
              </th>
              <th width="70px" title="Skor Kekalahan">
                SC-K
              </th>
              <th width="70px" title="Agregat Skor">
                A-SC
              </th>
              <th width="100px" title="Poin" className="has-text-centered">
                Poin
              </th>
            </tr>
          </thead>
          <tbody>
            {datas.map((data, i) => (
              <tr key={data.ID}>
                <td>{i + 1}</td>
                <td>{data.Name}</td>
                <td>0</td>
                <td>0</td>
                <td>0</td>
                <td>0</td>
                <td>0</td>
                <td>0</td>
                <td>0</td>
                <td>0</td>
                <td>0</td>
                <td className="has-text-centered has-text-weight-semibold">
                  0
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};
