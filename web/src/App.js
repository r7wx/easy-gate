import React, { useEffect, useState } from "react";
import Loading from "./components/Loading";
import Service from "./components/Service";
import Error from "./components/Error";
import { Helmet } from "react-helmet";
import Note from "./components/Note";
import axios from "axios";

function App() {
  const [loading, setLoading] = useState(true);
  const [data, setData] = useState({
    title: "Easy Gate",
    error: "",
    services: [],
    notes: [],
    theme: {
      background: "#ffffff",
      foreground: "#1d1d1d",
    },
  });

  useEffect(() => {
    const fetchData = () => {
      axios
        .get("/api/data")
        .then((res) => {
          setData(res.data);
          setLoading(false);
        })
        .catch((_) => {
          setLoading(true);
        });
    };
    fetchData();
  }, []);

  return (
    <React.Fragment>
      {!loading ? (
        <main className="py-6 px-12">
          <Helmet>
            {data.title && <title>{data.title}</title>}
            {data.theme && (
              <style>
                {`body { background-color: ${data.theme.background}; 
                color: ${data.theme.foreground}}`}
              </style>
            )}
          </Helmet>
          {data.error.length > 0 && <Error error={data.error} />}
          {data.services.length > 0 && (
            <React.Fragment>
              <div
                className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-7
                mt-4 mb-2">
                {data.services.map((service) => (
                  <Service
                    key={service.name}
                    service={service}
                    theme={data.theme}
                  />
                ))}
              </div>
            </React.Fragment>
          )}
          {data.notes.length > 0 && (
            <React.Fragment>
              <div
                className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 xl:grid-cols-4
                mt-2 mb-2">
                {data.notes.map((note) => (
                  <Note key={note.title} note={note} />
                ))}
              </div>
            </React.Fragment>
          )}
        </main>
      ) : (
        <Loading />
      )}
    </React.Fragment>
  );
}

export default App;
