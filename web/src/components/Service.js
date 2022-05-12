import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

function Service(props) {
  const openService = (url) => {
    window.open(url, "_blank");
  };

  return (
    <div
      onClick={() => openService(props.service.url)}
      className="px-6 py-5 rounded overflow-hidden shadow-lg text-center
      cursor-pointer hover:shadow-gray-600 mr-3 mb-2">
      <FontAwesomeIcon icon={props.service.icon} className="fa-3x" />
      <h3 className="mt-2">{props.service.name}</h3>
    </div>
  );
}

export default Service;