import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faLink } from "@fortawesome/free-solid-svg-icons";
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
            <FontAwesomeIcon icon={faLink} className="mr-2 mt-0.5 fa-lg" />
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
