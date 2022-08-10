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

import { faCircleNodes } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React from "react";

function Service(props) {
  const openService = (url) => {
    window.open(url, "_blank");
  };

  return (
    <div
      onClick={() => openService(props.service.url)}
      className="p-4 rounded shadow-lg cursor-pointer hover:shadow-gray-600 m-1">
      <div className="flex">
        <React.Fragment>
          {props.service.icon.length === 0 ? (
            <FontAwesomeIcon
              icon={faCircleNodes}
              className="mr-2 mt-0.5 fa-lg"
            />
          ) : (
            <img
              alt="service_icon"
              className="mr-2"
              width="25px"
              height="20px"
              src={props.service.icon}
            />
          )}
        </React.Fragment>
        <p className="whitespace-nowrap overflow-hidden text-ellipsis font-semibold w-5/6">
          {props.service.name}
        </p>
      </div>
    </div>
  );
}

export default Service;
