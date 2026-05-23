import {defineStore} from "pinia";
import {type Notification} from "../composables/notification";
import {WebSocketClient} from "../api/webSocketClient";
import type {Account} from "../composables/account";
import type { Product } from "../composables/product";
import type { ProductVisibility } from "../composables/productVisibility";

export const useSocketStore = defineStore("notificationStore", {
    state: () => ({
        wsClient: new WebSocketClient(`ws://${location.host.split(":")[0]}:8081/ws`),
        notifications: [] as Notification[],
    }),
    actions: {
        addNotification(msg: any) {
            //ToDo: Limit this buffer to a certain amount of messages
            this.notifications.push(msg);
        },
        sendMessage(msg: string) {
            this.wsClient.send({type: "hello", content: msg})
        },
        queryAccountExample() {
            let payload = {
                "table": "account",
                "filter": [{"column": "account_id", "operator": "gt", "value": "1"}, {
                    "column": "balance",
                    "operator": "gt",
                    "value": "12.04"
                }],
                "limit": 10,
                "sort": {"column": "account_id", "order": "desc"}
            }
            this.wsClient.send({type: "query", payload: payload})
        },
        addAccount(newAccount: Account) {
            let payload = {
                "operation": "insert",
                "table": "account",
                "values": newAccount
            }
            let msg = {type: "mutation", payload: payload}
            //console.debug(msg)
            this.wsClient.send(msg)
        },
        updateAccount(refAccount: Account){
            const table = "account";
            const operation = "update";
            const values = refAccount;
            const where = {"account_id": refAccount.id?.toString()};
            let payload = {
                "operation": operation,
                "table": table,
                "where": where,
                "values": values
            }
            let msg = {type: "mutation", payload: payload};
            this.wsClient.send(JSON.stringify(msg));
        },
        addVisibility(newVisibility: ProductVisibility) {
            let payload = {
                "operation": "insert",
                "table": "product_visibility",
                "values": newVisibility
            }
            let msg = {type: "mutation", payload: payload}
            //console.debug(msg)
            this.wsClient.send(msg)
        },
        removeVisibility(visId: number){
            let payload = {
                "operation": "delete",
                "table": "product_visibility",
                "where": {"product_visibility_id": visId.toString()} 
            }
            let msg = {type: "mutation", payload: payload}
            //console.debug(msg)
            this.wsClient.send(msg)
        },
        addProduct(newProduct: Product) {
            let payload = {
                "operation": "insert",
                "table": "product",
                "values": newProduct
            }
            let msg = {type: "mutation", payload: payload}
            //console.debug(msg)
            this.wsClient.send(msg)
        },
        addCategory(newCategory: { name: string, icon: [string, string], type: string }) {
            let payload = {
                "operation": "insert",
                "table": "category",
                "values": {
                    "name": newCategory.name,
                    "enabled": true,
                    "icon": newCategory.icon,
                    "type": newCategory.type
                }
            }
            let msg = {type: "mutation", payload: payload}
            this.wsClient.send(msg)
        },
        toggleAccountEnabled(id: number, enabled: boolean) {
            let payload = {
                "operation": "update",
                "table": "account",
                "where": { "account_id": id.toString() },
                "values": { "enabled": enabled }
            }
            this.wsClient.send({ type: "mutation", payload: payload })
        },
        toggleCategoryEnabled(id: number, enabled: boolean) {
            let payload = {
                "operation": "update",
                "table": "category",
                "where": { "category_id": id.toString() },
                "values": { "enabled": enabled }
            }
            this.wsClient.send({ type: "mutation", payload: payload })
        },
        updateCategory(cat: { id: number, name: string, icon: [string, string], type: string }) {
            let payload = {
                "operation": "update",
                "table": "category",
                "where": { "category_id": cat.id.toString() },
                "values": {
                    "name": cat.name,
                    "icon": cat.icon,
                    "type": cat.type
                }
            }
            let msg = {type: "mutation", payload: payload}
            this.wsClient.send(msg)
        },
        queryCategories() {
            let payload = {
                "table": "category",
            }
            let msg = {type: "query", payload: payload}
            //console.debug(msg)
            this.wsClient.send(msg)
        }, queryAccounts() {
            let payload = {
                "table": "account",
                "sort": {"column": "nickname", "order": "asc"}
            }
            let msg = {type: "query", payload: payload}
            //console.debug(msg)
            this.wsClient.send(msg)
        }, queryProductGroups() {
            let payload = {
                "table": "product_group",
            }
            let msg = {type: "query", payload: payload}
            //console.debug(msg)
            this.wsClient.send(msg)
        }, queryUnits() {
            let payload = {
                "table": "unit",
            }
            let msg = {type: "query", payload: payload}
            //console.debug(msg)
            this.wsClient.send(msg)
        }, queryVats() {
            let payload = {
                "table": "vat",
            }
            let msg = {type: "query", payload: payload}
            //console.debug(msg)
            this.wsClient.send(msg)
        },
        queryProducts() {
            let payload = {
                "table": "product",
            }
            let msg = {type: "query", payload: payload}
            //console.debug(msg)
            this.wsClient.send(msg)
        },
        queryLocations() {
            let payload = {
                "table": "location",
            }
            let msg = {type: "query", payload: payload}
            //console.debug(msg)
            this.wsClient.send(msg)
        },
        queryProductVisibilities() {
            let payload = {
                "table": "product_visibility",
            }
            let msg = {type: "query", payload: payload}
            //console.debug(msg)
            this.wsClient.send(msg)
        },
        getAllFromDb() {
            this.queryCategories()
            this.queryAccounts()
            this.queryProductGroups()
            this.queryUnits()
            this.queryVats()
            this.queryProducts()
            this.queryLocations()
            this.queryProductVisibilities()
        }
    }
})