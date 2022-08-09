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

function Dot(props) {
  const [colors, setColors] = useState({
    health_ok: "",
    health_bad: "",
    health_inactive: "",
  });

  useEffect(() => {
    setColors({
      health_ok: props.theme.health_ok,
      health_bad: props.theme.health_bad,
      health_inactive: props.theme.health_inactive,
    });
  }, [props.theme]);

  const statusDot = () => {
    switch (props.health) {
      case 1:
        return (
          <span
            className="dot"
            style={{
              backgroundColor: colors.health_ok,
            }}
          />
        );
      case 2:
        return (
          <span
            className="dot"
            style={{
              backgroundColor: colors.health_bad,
            }}
          />
        );
      default:
        return (
          <span
            className="dot"
            style={{
              backgroundColor: colors.health_inactive,
            }}></span>
        );
    }
  };

  return <React.Fragment>{statusDot()}</React.Fragment>;
}

export default Dot;
