/*
MIT License

Copyright (c) 2022 r7wx

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

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
      health_ok: "#22c55e",
      health_bad: "#ef4444",
      health_inactive: "#d1d5db",
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
