import {defineStore} from "pinia";
import {type Notification} from "../composables/notification";
import {WebSocketClient} from "../api/webSocketClient";

export const useSocketStore = defineStore("notificationStore", {
    state: () => ({
        wsClient: new WebSocketClient("ws://localhost:8081/ws"),
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
        queryAccount(accountId: number) {
            //let payload = {"table": "accounts"}
            //let payload = {"table": "accounts", "filter": [{"column": "account_id", "operator": "eq", "value": "1"}]}
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
            //let payload = {"table": "accounts", "filter": [{"column": "account_id", "operator": "gt", "value": "1"}, {"column": "balance", "operator": "gt", "value": "12.04"}]}
            //let payload = {"table": "accounts", "id": accountId}
            //let payload = {"table": "accounts", "filter": {"Id": accountId}}
            this.wsClient.send({type: "query", payload: payload})
        },
        queryAccounts(){
            let payload = {
                "table": "account",
                "sort": {"column": "nickname", "order": "asc"}
            }
            this.wsClient.send({type: "query", payload: payload})
        }
    }
})