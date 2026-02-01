export interface ProductGroup {
    id?: number
    name: string
}

export class ProductGroup implements ProductGroup {

    constructor(group: ProductGroup) {
        this.id = group.id;
        this.name = group.name;
    }
}
