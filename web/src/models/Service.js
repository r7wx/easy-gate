/**
 * Model of a service.
 */
export class Service {
    /**
     * @constructor
     * @param {object} service
     */
    constructor(service) {
        this.name = service.name || "";
        this.icon = service.icon || "";
        this.url = service.url || "";
        this.groups = service.groups || [];
        this.category = service.category || 1000;
        this.labels = service.labels || [];
    }

    /**
     * Returns the list of services as an array of Service
     * objects from the given raw services data array.
     * @param {Service[]} serviceDatas 
     * @returns Service[]
     */
    static map(serviceDatas) {
        const services = serviceDatas.map((service) => {
            return new Service(service);
        });
        return services;
    }
}