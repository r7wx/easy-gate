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

function Service(props) {
  const openService = (url) => {
    window.open(url, "_blank");
  };

  return (
    <div
      onClick={() => openService(props.service.url)}
      className="p-8 rounded overflow-hidden shadow-lg text-center
      cursor-pointer hover:shadow-gray-600 mr-3 mb-2 grid place-items-center">
      <FontAwesomeIcon icon={props.service.icon} className="fa-3x" />
      <p className="mt-3 text-ellipsis overflow-hidden ...">
        {props.service.name}
      </p>
    </div>
  );
}

export default Service;
