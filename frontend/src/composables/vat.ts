export interface Vat {
    id?: number
    rate: number
}

export class Vat implements Vat {

    constructor(vat: Vat) {
        this.id = vat.id;
        this.rate = vat.rate;
    }
}
