interface WebSocketClientOptions {
    reconnectInterval?: number;
    maxReconnectAttempts?: number;
    heartbeatInterval?: number;
}

type EventHandler<T = any> = (data: T) => void;

interface EventHandlers {
    [event: string]: EventHandler[];
}

interface PongMessage {
    type: 'pong';

    [key: string]: any;
}

interface TypedMessage {
    type: string;

    [key: string]: any;
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
            console.log('Message received:', event.data);

            // Try to parse JSON messages
            let data: any = event.data;
            try {
                data = JSON.parse(event.data);
            } catch (e) {
                // Not JSON, use as-is
            }

            // Handle ping/pong for heartbeat
            if (data.type === 'pong') {
                this.lastPong = Date.now();
                return;
            }

            // Trigger custom message handlers
            this.trigger('message', data);

            // Trigger typed message handlers
            if (data.type) {
                this.trigger(data.type, data);
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
                this.send({type: 'ping', timestamp: Date.now()});

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
