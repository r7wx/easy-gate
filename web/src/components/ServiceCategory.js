import Service from "./Service";
import * as React from 'react'

function ServiceCategory(props) {

  return (
    <React.Fragment>
      <table>
        <tbody>
          <tr>
            <td>
              <h3 className="text-lg mt-3">
                {props.category.title}
              </h3>
            </td>
          </tr>
          <tr>
            <td>
              <div className="grid grid-cols-1 md:grid-cols-6 lg:grid-cols-8 xl:grid-cols-12 ml-2">
                {[].concat(props.services)
                  // Sort alphabetically
                  .sort(function (a, b) {
                    if (a.name.toLowerCase() < b.name.toLowerCase()) return -1;
                    if (a.name.toLowerCase() > b.name.toLowerCase()) return 1;
                    return 0;
                  })
                  .map((service) => (
                    <Service key={service.name} service={service} />
                  ))}
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </React.Fragment>
  );
}

export default ServiceCategory;
