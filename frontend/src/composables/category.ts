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

export function getTestCategorys(): Category[] {
    let c1: Category = {
        id: 1,
        title: "Rooms",
        icon: ['fas', 'house'],
        type: CategoryType.ACCOUNT,
        visibility: 1
    }
    c1 = new Category(c1);

    let c2: Category = {
        id: 2,
        title: "Guests",
        icon: ['fas', 'beer'],
        type: CategoryType.ACCOUNT,
        visibility: 1
    }
    c2 = new Category(c2);

    let c3: Category = {
        id: 3,
        title: "Beverages",
        icon: ['fas', 'beer'],
        type: CategoryType.PRODUCT,
        visibility: 1
    }
    c3 = new Category(c3);

    let c4: Category = {
        id: 4,
        title: "Food",
        icon: ['fas', 'burger'],
        type: CategoryType.PRODUCT,
        visibility: 1
    }
    c4 = new Category(c4);
    return [c1, c2, c3, c4];
}

export function generateTestData(): Category[] {
    return getTestCategorys();
}
