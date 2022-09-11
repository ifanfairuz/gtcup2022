import React from "react";

export const Team = ({ datas }) => {
  return (
    <form
      action="/bla/team/update"
      method="post"
      className="has-background-white"
    >
      <div className="is-flex is-align-items-center is-justify-content-space-between p-2">
        <p className="title is-6 mb-0">Data Team</p>
        <button type="submit" className="button is-dark is-small">
          Save
        </button>
      </div>
      <div className="p-2">
        <div className="table-container">
          <table className="table is-striped is-fullwidth is-narrow">
            <thead>
              <tr>
                <th width="50px">Id</th>
                <th>Name</th>
                <th>Alamat</th>
                <th width="100px">Group</th>
              </tr>
            </thead>
            <tbody>
              {datas.map((t) => (
                <tr key={t.ID}>
                  <td>{t.ID}</td>
                  <td>
                    <input
                      className="input is-small"
                      type="text"
                      name={`Name-${t.ID}`}
                      defaultValue={t.Name}
                    />
                  </td>
                  <td>
                    <input
                      className="input is-small"
                      type="text"
                      name={`Alamat-${t.ID}`}
                      defaultValue={t.Alamat}
                    />
                  </td>
                  <td>
                    <div className="select is-small">
                      <select name={`Group-${t.ID}`} defaultValue={t.Group}>
                        <option value="A">A</option>
                        <option value="B">B</option>
                        <option value="C">C</option>
                        <option value="D">D</option>
                      </select>
                    </div>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </div>
    </form>
  );
};
