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

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { library } from "@fortawesome/fontawesome-svg-core";
import { far } from "@fortawesome/free-regular-svg-icons";
import { fas } from "@fortawesome/free-solid-svg-icons";
import { fab } from "@fortawesome/free-brands-svg-icons";
import Service from "./components/Service";
import Note from "./components/Note";
import React from "react";
import axios from "axios";

library.add(fas, far, fab);

function App() {
  const [data, setData] = React.useState({
    title: "",
    icon: "",
    motd: "",
    services: [],
    notes: [],
  });

  const fetchData = () => {
    axios
      .get("/api/data")
      .then((res) => {
        setData(res.data);
      })
      .catch((_) => {});
  };

  React.useEffect(() => {
    fetchData();
  }, []);

  document.title = data.title;

  return (
    <main className="py-6 px-12">
      <h1 className="text-4xl">
        <FontAwesomeIcon icon={data.icon} /> {data.title}
      </h1>
      <p className="text-base">{data.motd}</p>
      {data.services.length > 0 && (
        <React.Fragment>
          <h3 className="text-xl mt-5">
            <FontAwesomeIcon icon="fa-brands fa-buffer" /> Services
          </h3>
          <div className="grid grid-cols-1 md:grid-cols-6 lg:grid-cols-8 xl:grid-cols-12 mt-5">
            {data.services.map((service) => (
              <Service key={service.name} service={service} />
            ))}
          </div>
        </React.Fragment>
      )}
      {data.notes.length > 0 && (
        <React.Fragment>
          <h3 className="text-xl mt-5">
            <FontAwesomeIcon icon="fa-regular fa-note-sticky" /> Notes
          </h3>
          <div className="grid grid-cols-1 md:grid-cols-1 lg:grid-cols-2 xl:grid-cols-4 mt-5">
            {data.notes.map((note) => (
              <Note key={note.title} note={note} />
            ))}
          </div>
        </React.Fragment>
      )}
    </main>
  );
}

export default App;
