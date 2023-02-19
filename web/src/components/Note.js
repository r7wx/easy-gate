import { faNoteSticky } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

function Note(props) {
  return (
    <div className="p-4 rounded shadow-lg m-1">
      <div className="flex">
        <FontAwesomeIcon icon={faNoteSticky} className="mr-2 mt-1 fa-sm" />
        <h3 className="text-sm font-semibold">{props.note.name}</h3>
      </div>
      <p className="text-sm">{props.note.text}</p>
    </div>
  );
}

export default Note;
