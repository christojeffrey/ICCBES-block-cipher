import { createContextId } from "@builder.io/qwik";

// Declare a context ID
export const inputContext = createContextId<{ type: string; value: string }>("inputText");
export const buttonContext = createContextId<{ pressed: boolean }>("buttonValue");
export const configContext = createContextId<{ blockCipherMode: string; encryptionMode: string; autofill: boolean }>("config");
export const modeSpecificConfigContext = createContextId<any>("modeSpecificConfig");
