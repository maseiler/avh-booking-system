import { defineStore } from "pinia";
import { type Notification } from "../composables/notification";
import { WebSocketClient } from "../api/webSocketClient";

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
      this.wsClient.send({ type: 'hello', content: msg })
    }
  }
})