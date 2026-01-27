import {ref} from "vue";
import {getApi} from "../api/api";

type Icon = [string, string];

/**
 * Type definitions
 */

export const CategoryType = {
    ACCOUNT: 0,
    PRODUCT: 1
} as const;
export type CategoryType = (typeof CategoryType)[keyof typeof CategoryType];

export const CategoryTypeToString: Record<CategoryType, string> = {
    [CategoryType.ACCOUNT]: "account",
    [CategoryType.PRODUCT]: "product",
};
export const StringToCategoryType: Record<string, CategoryType> = {
    ["account"]: CategoryType.ACCOUNT,
    ["product"]: CategoryType.PRODUCT,
};

export const CategoryVisibility = {
    HIDDEN: 0,
    SHOWN: 1
} as const;
export type CategoryVisibility = (typeof CategoryVisibility)[keyof typeof CategoryVisibility];

export interface Category {
    title: string
    type: CategoryType
    visibility: CategoryVisibility
    icon?: Icon
    id: number
}

export class Category implements Category {

    constructor(cat: Category) {
        this.title = cat.title;
        this.type = cat.type;
        this.visibility = cat.visibility;
        this.icon = cat.icon;
        this.id = cat.id;
    }
}

/**
 * Functions
 */

export function createCategory(obj: any): Category {
    const cat = {} as Category;
    cat.id = obj.id;
    cat.title = obj.name;
    cat.visibility = obj.enabled;
    cat.icon = obj.icon;
    cat.type = StringToCategoryType[obj.type];
    return new Category(cat);
}

export async function getCategorys() {
    let {response: categories, request} = getApi<Category[]>(
        "/category"
    );

    let loaded = ref(false);
    if (loaded.value === false) {
        await request();
        loaded.value = true;
    }

    return {categories};
}

