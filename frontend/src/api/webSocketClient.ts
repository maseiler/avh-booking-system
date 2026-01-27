import {useAccountStore} from "../store/AccountStore";
import {compileSchema, draft07} from "json-schema-library";
import type {SchemaNode} from "json-schema-library";
import messageSchema from "../../../backend/internal/models/jsonSchemas/message.json";
import {Account} from '../composables/account'
import {useCategoryStore} from "../store/CategoryStore.ts";
import {useProductGroupStore} from "../store/ProductGroupStore.ts";
import {useUnitStore} from "../store/UnitStore.ts";
import {useVatStore} from "../store/VatStore.ts";
import {useProductStore} from "../store/ProductStore.ts";
import {useLocationStore} from "../store/LocationStore.ts";
import {useProductVisibilityStore} from "../store/ProductVisibilityStore.ts";

interface WebSocketClientOptions {
    reconnectInterval?: number;
    maxReconnectAttempts?: number;
    heartbeatInterval?: number;
}

type EventHandler<T = any> = (data: T) => void;

interface EventHandlers {
    [event: string]: EventHandler[];
}

interface Message {
    type: string;
    payload: Partial<any>
}

interface queryResult {
    table: string;
    data: Partial<any> // TODO only allow Account, Product, ...
}

interface queryResultList {
    table: string;
    data: Partial<any[]> // TODO only allow Account, Product, ...
}

interface ResultMutation {
    table: string;
    operation: string;
    id: number;
}

interface Pong {
    timestamp: Partial<Date>
}

export class WebSocketClient {
    private url: string;
    private options: Required<WebSocketClientOptions>;
    private reconnectAttempts: number;
    private messageQueue: any[];
    private eventHandlers: EventHandlers;
    private isConnected: boolean;
    private ws: WebSocket | null;
    private heartbeatTimer: number | null;
    private lastPong: number;
    private schema: SchemaNode;

    constructor(url: string, options: WebSocketClientOptions = {}) {
        this.url = url;
        this.options = {
            reconnectInterval: 1000,
            maxReconnectAttempts: 5,
            heartbeatInterval: 30000,
            ...options,
        };
        this.reconnectAttempts = 0;
        this.messageQueue = [];
        this.eventHandlers = {};
        this.isConnected = false;
        this.ws = null;
        this.heartbeatTimer = null;
        this.lastPong = 0;
        this.schema = compileSchema(messageSchema, {drafts: [draft07]});

        this.connect();
    }

    connect(): void {
        let url = this.url;
        const clientId = import.meta.env.VITE_WS_CLIENT_ID;
        if (clientId.length > 0) {
            url += "?id=" + clientId;
        }
        console.log(`Connecting to ${url}...`);

        try {
            this.ws = new WebSocket(url);
            this.setupEventHandlers();
        } catch (error) {
            console.error('Failed to create WebSocket:', error);
            this.scheduleReconnect();
        }
    }

    private setupEventHandlers(): void {
        if (!this.ws) return;

        this.ws.onopen = (event: Event): void => {
            console.log('WebSocket connected');
            this.isConnected = true;
            this.reconnectAttempts = 0;

            // Send any queued messages
            while (this.messageQueue.length > 0) {
                const message = this.messageQueue.shift();
                this.send(message);
            }

            // Start heartbeat
            this.startHeartbeat();

            // Trigger custom open handlers
            this.trigger('open', event);
        };

        this.ws.onmessage = (event: MessageEvent): void => {
            if (import.meta.env.DEV) {
                // Only Log WS Messages, when running in Dev Environment
                console.debug('Message received:', event.data);
            }

            // Try to parse JSON message (could be anything)
            let parsedMsg: any;
            try {
                parsedMsg = JSON.parse(event.data);
            } catch (e) {
                console.error('Failed to parse json:', e);
                // TODO handle error?
                return
            }

            // Validate that message complies with JSON schema
            const {valid, errors} = this.schema.validate(parsedMsg);
            if (!valid) {
                console.error('Invalid message schema:', errors);
                // TODO handle error?
                return
            }

            // Create Message interface for type safety
            const message = parsedMsg as Message
            // console.log("message type: ", message.type);

            switch (message.type) {
                case 'pong': {
                    const pong = message.payload as Pong;
                    if (pong.timestamp) {
                        console.log('Received pong with timestamp', pong.timestamp)
                    } else {
                        console.log('Received pong');
                    }
                    this.lastPong = Date.now();
                    return;
                }

                case 'queryResultList': {
                    const result = message.payload as queryResultList

                    switch (result.table) {
                        case 'account': {
                            useAccountStore().patchAccounts(result.data);
                            return;
                        }

                        case 'category': {
                            useCategoryStore().patchCategories(result.data);
                            return;
                        }

                        case 'location': {
                            useLocationStore().patchLocations(result.data);
                            return;
                        }

                        case 'product': {
                            useProductStore().patchProducts(result.data);
                            return;
                        }

                        case 'product_group': {
                            useProductGroupStore().patchProductGroups(result.data);
                            return;
                        }

                        case 'product_visibility': {
                            useProductVisibilityStore().patchVisibilities(result.data);
                            return;
                        }

                        case 'unit': {
                            useUnitStore().patchUnits(result.data);
                            return;
                        }

                        case 'vat': {
                            useVatStore().patchVats(result.data);
                            return;
                        }

                        default: {
                            console.error('TODO handle table', result.table)
                        }
                    }
                    break;
                }

                case 'mutationResult': {
                    const res = message.payload as ResultMutation
                    console.info(res)
                    return
                }
                case 'broadcast' : {
                    console.info('Received broadcast');
                    // TODO check if payload exists
                    const result = message.payload as queryResult
                    switch (result.table) {
                        case 'account': {
                            const account = result.data as Account;
                            console.log(account)
                            // TODO do stuff
                            return;
                        }
                    }
                    return
                }
                default: {
                    console.log('TODO handle message type', message.type);
                }
            }

            // Trigger custom message handlers
            this.trigger('message', message);

            // Trigger typed message handlers
            if (message.type) {
                this.trigger(message.type, message);
            }
        };

        this.ws.onerror = (error: Event): void => {
            console.error('WebSocket error:', error);
            this.trigger('error', error);
        };

        this.ws.onclose = (event: CloseEvent): void => {
            console.log(`WebSocket closed: ${event.code} - ${event.reason}`);
            this.isConnected = false;
            this.stopHeartbeat();

            // Trigger custom close handlers
            this.trigger('close', event);

            // Attempt to reconnect if not a normal closure
            if (event.code !== 1000 && event.code !== 1001) {
                this.scheduleReconnect();
            }
        };
    }

    send(message: any): void {
        if (this.ws && this.ws.readyState === WebSocket.OPEN) {
            const data =
                typeof message === 'object' ? JSON.stringify(message) : message;
            this.ws.send(data);
        } else {
            // Queue message if not connected
            console.log('WebSocket not connected, queuing message');
            this.messageQueue.push(message);
        }
    }

    private startHeartbeat(): void {
        this.stopHeartbeat();
        this.heartbeatTimer = setInterval(() => {
            if (this.ws && this.ws.readyState === WebSocket.OPEN) {
                this.send({type: 'ping', payload: {timestamp: new Date(Date.now()).toISOString()}});

                // Check for pong timeout
                setTimeout(() => {
                    const timeSinceLastPong = Date.now() - (this.lastPong || 0);
                    if (timeSinceLastPong > this.options.heartbeatInterval * 2) {
                        console.log('Heartbeat timeout, reconnecting...');
                        this.ws?.close();
                    }
                }, 5000);
            }
        }, this.options.heartbeatInterval);
    }

    private stopHeartbeat(): void {
        if (this.heartbeatTimer) {
            clearInterval(this.heartbeatTimer);
            this.heartbeatTimer = null;
        }
    }

    private scheduleReconnect(): void {
        if (this.reconnectAttempts >= this.options.maxReconnectAttempts) {
            console.error('Max reconnection attempts reached');
            this.trigger('maxReconnectAttemptsReached', undefined);
            return;
        }

        this.reconnectAttempts++;
        const delay =
            this.options.reconnectInterval * Math.pow(2, this.reconnectAttempts - 1);
        console.log(
            `Reconnecting in ${delay}ms (attempt ${this.reconnectAttempts})...`
        );

        setTimeout(() => {
            this.connect();
        }, delay);
    }

    on<T = any>(event: string, handler: EventHandler<T>): void {
        if (!this.eventHandlers[event]) {
            this.eventHandlers[event] = [];
        }
        this.eventHandlers[event].push(handler);
    }

    off<T = any>(event: string, handler: EventHandler<T>): void {
        if (this.eventHandlers[event]) {
            this.eventHandlers[event] = this.eventHandlers[event].filter(
                (h) => h !== handler
            );
        }
    }

    private trigger<T = any>(event: string, data: T): void {
        if (this.eventHandlers[event]) {
            this.eventHandlers[event].forEach((handler) => {
                try {
                    handler(data);
                } catch (error) {
                    console.error(`Error in ${event} handler:`, error);
                }
            });
        }
    }

    close(): void {
        this.reconnectAttempts = this.options.maxReconnectAttempts;
        this.stopHeartbeat();
        if (this.ws) {
            this.ws.close(1000, 'Client closing connection');
        }
    }
}
