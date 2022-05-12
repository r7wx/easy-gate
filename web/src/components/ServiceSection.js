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
import ServiceCategory from "./ServiceCategory";
import * as React from 'react'

function ServiceSection(props) {

  return (
    <React.Fragment>
      <h3 className="text-xl mt-5">
        <FontAwesomeIcon icon="fa-brands fa-buffer" /> Services
      </h3>
      {[].concat(props.categories.filter((category) => props.services.some((service) => service.category == category.id)))
        // Sort alphabetically
        .sort(function (a, b) {
          if (a.title.toLowerCase() < b.title.toLowerCase()) return -1;
          if (a.title.toLowerCase() > b.title.toLowerCase()) return 1;
          return 0;
        })
        .map((category) => (
          // Row with 5 margin on the left
          <row className="flex flex-wrap mx-5" key={category.id}>
            <ServiceCategory
              category={category}
              services={props.services.filter(service => service.category == category.id)} />
          </row>
        ))}
    </React.Fragment>
  );
}

export default ServiceSection;
