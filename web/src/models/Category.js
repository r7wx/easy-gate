/**
 * Model of a category.
 */
export class Category {
    /**
     * @constructor
     * @param {object} category
     */
    constructor(category) {
        // log the type of the category
        this.id = category.id || 1000;
        this.title = category.title || "";
        this.description = category.description || "";
    }

    /**
     * Returns the list of categories as an array of Category
     * objects from the given raw category data array.
     * @param {Category[]} categoryDatas 
     * @returns Category[]
     */
    static map(categoryDatas) {
        const categories = categoryDatas.map((category) => {
            return new Category(category);
        });
        return categories;
    }
}