import type {Location} from "./location.ts";
import type {Category} from "./category.ts";
import type {Product} from "./product.ts";
import {useCategoryStore} from "../store/CategoryStore.ts";
import {useLocationStore} from "../store/LocationStore.ts";
import {useProductStore} from "../store/ProductStore.ts";

export interface ProductVisibility {
    id?: number
    categoryId: number
    locationId: number
    productId: number
    getCategory(): Category
    getProduct(): Product
    getLocation(): Location
}

export class ProductVisibility implements ProductVisibility {

    constructor(vis: ProductVisibility) {
        this.id = vis.id;
        this.categoryId = vis.categoryId;
        this.locationId = vis.locationId;
        this.productId = vis.productId;
    }

    public getCategory(): Category | undefined{
        return useCategoryStore().byId(this.categoryId)
    }

    public getProduct(): Product | undefined{
        return useProductStore().byId(this.productId)
    }

    public getLocation(): Location | undefined{
        return useLocationStore().byId(this.locationId)
    }

}

// export function createProductVisibility(obj: any): ProductVisibility {
//     const vis = {} as ProductVisibility
//     vis.id = obj.id
//     let category = useCategoryStore().byId(obj.categoryId)
//     if (category) {
//         vis.category = category
//     } else {
//         console.error("Category not found")
//         // TODO handler error
//     }
//     let location = useLocationStore().byId(obj.locationId)
//     if (location)
//         vis.location = location
//     else {
//         console.error("Group not found")
//         // TODO handle error
//     }
//     let product = useProductStore().byId(obj.productId)
//     if (product)
//         vis.product = product
//     else {
//         console.error("Unit not found")
//         // TODO handle error
//     }

//     return vis
// }
