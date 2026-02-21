import {ProductGroup} from "./productGroup"
import {Unit} from "./unit"
import {Vat} from "./vat.ts";
import {Category} from "./category.ts";
import {useVatStore} from "../store/VatStore.ts";
import {useProductGroupStore} from "../store/ProductGroupStore.ts";
import {useUnitStore} from "../store/UnitStore.ts";
import {useCategoryStore} from "../store/CategoryStore.ts";

export interface Product {
    id?: number
    name: string
    price: number
    vat: number
    group: number
    size: number
    unit: number
    category: number
    createdAt: Date
    getVat(): Vat
    getGroup(): ProductGroup
    getUnit(): Unit
    getCategory(): Category
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
        this.vat = prod.vat;
    }

    public getVat(): Vat | undefined{
        return useVatStore().byId(this.vat);
    }

    public getGroup(): ProductGroup | undefined{
        return useProductGroupStore().byId(this.group)
    }

    public getUnit(): Unit | undefined{
        return useUnitStore().byId(this.unit)
    }

    public getCategory(): Category | undefined{
        return useCategoryStore().byId(this.category)
    }

    public copy(): Product{
        let newProduct = new Product(this);
        return newProduct;
    }

    public update(reference: Product): Product{
        this.id = reference.id;
        this.name = reference.name;
        this.price = reference.price;
        this.group = reference.group;
        this.size = reference.size;
        this.unit = reference.unit;
        this.category = reference.category;
        return this;
  }
}