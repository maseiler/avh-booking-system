import type {ProductGroup} from "./productGroup"
import type {Unit} from "./unit"
import type {Vat} from "./vat.ts";
import type {Category} from "./category.ts";
import {useVatStore} from "../store/VatStore.ts";
import {useProductGroupStore} from "../store/ProductGroupStore.ts";
import {useUnitStore} from "../store/UnitStore.ts";
import {useCategoryStore} from "../store/CategoryStore.ts";

export interface Product {
    id?: number
    name: string
    price: number
    vat: Vat
    group: ProductGroup
    size: number
    unit: Unit
    category: Category
    createdAt: Date
}

export class Product implements Product {

    public constructor(prod: Product) {
        this.id = prod.id;
        this.name = prod.name;
        this.price = prod.price;
        this.group = prod.group;
        this.size = prod.size;
        this.unit = prod.unit;
        this.category = prod.category;
    }
}

export function createProduct(obj: any): Product {
    const prod = {} as Product
    prod.id = obj.id
    prod.name = obj.name
    prod.price = obj.price
    let vat = useVatStore().getById(obj.vatId)
    if (vat) {
        prod.vat = vat
    } else {
        console.error("VAT not found")
        // TODO handler error
    }
    let group = useProductGroupStore().byId(obj.productGroupId)
    if (group)
        prod.group = group
    else {
        console.error("Group not found")
        // TODO handle error
    }
    prod.size = obj.size
    let unit = useUnitStore().byId(obj.unitId)
    if (unit)
        prod.unit_id = unit
    else {
        console.error("Unit not found")
        // TODO handle error
    }
    let cat = useCategoryStore().byId(obj.categoryId)
    if (cat)
        prod.category = cat
    else {
        console.error("Category not found")
        // TODO handle error
    }

    return prod
}

export function generateTestData() {
    /*
    let p1 = {
      id: 1,
      name: "Beer",
      price: 120,
      size: 0.5,
      tax: 19,
      category: 3,
      visibility: [2]
    } as Product
    p1 = new Product(p1);

    let p2 = {
      id: 2,
      name: "Bread",
      price: 250,
      size: 1,
      tax: 0,
      category: 4,
      visibility: [1, 2]
    } as Product
    p2 = new Product(p2);

    return [p1, p2];
     */
}

// export function formatPrice(price: number){

// }