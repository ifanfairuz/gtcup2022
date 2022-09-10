import React from "react";
import { renderToRoot, getData } from "@support/render";
import { Navbar, Footer } from "@components";

const App = ({ Teams }) => {
  return (
    <div>
      <Navbar />
      <form
        action="/bla/update-team"
        method="post"
        className="has-background-white p-4"
      >
        <table className="table is-striped is-fullwidth">
          <thead>
            <tr>
              <th>ID</th>
              <th>Name</th>
              <th>Alamat</th>
              <th>Group</th>
            </tr>
          </thead>
          <tbody>
            {Teams.map((t) => (
              <tr key={t.ID}>
                <td>{t.id}</td>
                <td>
                  <input
                    type="text"
                    name={`Name-${t.ID}`}
                    defaultValue={t.Name}
                  />
                </td>
                <td>
                  <input
                    type="text"
                    name={`Alamat-${t.ID}`}
                    defaultValue={t.Alamat}
                  />
                </td>
                <td>
                  <input
                    type="text"
                    name={`Group-${t.ID}`}
                    defaultValue={t.Group}
                  />
                </td>
              </tr>
            ))}
          </tbody>
        </table>
        <button className="button">SAVE</button>
      </form>
      <Footer />
    </div>
  );
};

const data = getData();
renderToRoot(<App {...data} />);
