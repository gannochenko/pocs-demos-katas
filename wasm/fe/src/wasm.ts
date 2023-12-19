import "./wasm_exec";

let go: Go | null = null;

export type GetItemsRequest = {
    amount: number;
};

export type Item = {
    id: string;
    title: string;
    date: string;
}

export type GetItemsResponse = {
    error: string;
    items: Item[];
}

declare global {
    interface Window {
        getItems: (request: GetItemsRequest) => GetItemsResponse;
    }
}

type Methods = {
    getItems: (request: GetItemsRequest) => GetItemsResponse;
};

export const getWASM = async (): Promise<Methods> => {
    if (!go) {
        go = new Go();
        const result = await WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject);
        go.run(result.instance);
    }

    return {
        getItems: window.getItems,
    };
};
