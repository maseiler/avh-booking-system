import {defineStore} from "pinia";
import {type Notification} from "../composables/notification";
import {WebSocketClient} from "../api/webSocketClient";
import type {Account} from "../composables/account";

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
        addTestAccount() {
            let now = Date.now();
            let newAccount = {
                "id": 0,
                "firstName": "Darude",
                "nickname": now.toString(),
                "lastName": "Sandstorm",
                "email": "ohsofunny@troll.lol",
                "phone": "12345678",
                "balance": 0,
                "maxDebt": 99,
                "category": 1,
                "enabled": true,
                "createdAt": new Date(now).toISOString()
            }
            let payload = {
                "operation": "insert",
                "table": "account",
                "values": newAccount
            }
            let msg = {type: "mutation", payload: payload}
            this.wsClient.send(msg)
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
        }, queryCategories() {
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