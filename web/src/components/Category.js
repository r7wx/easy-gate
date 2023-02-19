import { useEffect, useState } from "react";
import Service from "./Service";

function Category({ category, services }) {
  const [categoryServices, setCategoryServices] = useState([]);

  useEffect(() => {
    setCategoryServices(
      services.filter((service) => service.category === category)
    );
  }, [category, services]);

  return (
    <div>
      {category !== "" && <h2 className="ml-1">{category}</h2>}
      <div
        className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 xl:grid-cols-7
                mt-2 mb-2">
        {categoryServices.map((service) => (
          <Service key={service.name} service={service} />
        ))}
      </div>
    </div>
  );
}

export default Category;
