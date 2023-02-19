function Error(props) {
  return (
    <div
      className="my-3 bg-red-100 border-l-4 border-red-500 text-red-700 p-4"
      role="alert">
      <p className="font-bold">Configuration errror</p>
      <p>{props.error}</p>
    </div>
  );
}

export default Error;
