import { component$, createContextId, useStore, useContextProvider, useContext, useResource$, Resource } from "@builder.io/qwik";
import type { DocumentHead } from "@builder.io/qwik-city";

// Declare a context ID
export const inputContext = createContextId<{ type: string; value: string }>("inputText");
export const buttonContext = createContextId<{ pressed: boolean }>("buttonValue");
export const configContext = createContextId<{ blockCipherMode: string; encryptionMode: string; autofill: boolean }>("config");
export const modeSpecificConfigContext = createContextId<any>("modeSpecificConfig");

export default component$(() => {
  const input = useStore({ type: "text", value: "" });
  const button = useStore({ pressed: false });
  const config = useStore({ blockCipherMode: "cbc", encryptionMode: "encrypt", autofill: false });
  const modeSpecificConfig = useStore({});
  useContextProvider(inputContext, input);
  useContextProvider(buttonContext, button);
  useContextProvider(configContext, config);
  useContextProvider(modeSpecificConfigContext, modeSpecificConfig);

  return (
    <>
      <div class="md:h-screen flex flex-col p-3 gap-3">
        <AppHeader />

        <div class="border-2 flex flex-col md:flex-row md:justify-between p-3 gap-3 bg-gray-100 rounded-2xl flex-grow">
          {/* three section. left middle right. left for input, middle for configuration, right for output */}
          <InputBox />
          <ConfigBox />
          <OutputBox />
        </div>
      </div>
    </>
  );
});
const InputBox = component$(() => {
  const input = useContext(inputContext);
  return (
    <>
      <div class="min-h-[25vh] md:h-full md:w-1/3 rounded-2xl bg-white p-3">
        {/* heading */}
        <div>input box</div>
        <input
          type="text"
          value={input.value}
          placeholder="tesadsf"
          class="border-2 border-black"
          onChange$={(e) => {
            input.value = (e.target as HTMLInputElement).value;
          }}
        ></input>
      </div>
    </>
  );
});

const ConfigBox = component$(() => {
  const button = useContext(buttonContext);
  const config = useContext(configContext);
  // mode options = cbc, ecb, ofb, cfb
  // mode encryption options = encrypt, decrypt
  return (
    <>
      <div class="min-h-[25vh] md:h-full md:w-1/3 rounded-2xl bg-white p-3">
        config box
        {/* block cipher mode */}
        <select
          class="border-2 border-black p-2 rounded-md"
          onSelect$={(e) => {
            config.blockCipherMode = (e.target as HTMLSelectElement).value;
          }}
        >
          <option value="cbc">CBC</option>
          <option value="ecb">ECB</option>
          <option value="ofb">OFB</option>
          <option value="cfb">CFB</option>
        </select>
        {/* encryption mode */}
        <select
          class="border-2 border-black p-2 rounded-md"
          onSelect$={(e) => {
            config.encryptionMode = (e.target as HTMLSelectElement).value;
          }}
        >
          <option value="encrypt">Encrypt</option>
          <option value="decrypt">Decrypt</option>
        </select>
        {/* autofil checkbox */}
        <div class="flex items-center gap-2">
          <input type="checkbox" class="border-2 border-black p-2 rounded-md" onChange$={(e) => (config.autofill = (e.target as HTMLInputElement).checked)} />
          <label>Auto fill</label>
        </div>
        {/* mode specific input */}
        <button
          class="bg-blue-500 text-white p-2 rounded-md"
          onClick$={() => {
            button.pressed = false;
            button.pressed = true;
            console.log("button pressed", button.pressed);
          }}
        >
          run!
        </button>
      </div>
    </>
  );
});
const OutputBox = component$(() => {
  // print the input value
  const input = useContext(inputContext);
  const button = useContext(buttonContext);
  const config = useContext(configContext);
  const modeSpecificConfig = useContext(modeSpecificConfigContext);
  console.log(input.value);

  const result = useResource$<
    | {
        result: string;
        key: string;
        iv: string;
      }
    | undefined
  >(async ({ track }) => {
    // it will run first on mount (server), then re-run whenever prNumber changes (client)
    // this means this code will run on the server and the browser
    track(() => button.pressed);
    if (!button.pressed) return;
    const response = await fetch(import.meta.env.PUBLIC_BASE_URL + `/api/${config.blockCipherMode}/${config.encryptionMode}`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ text: input.value, ...modeSpecificConfig, autofill: config.autofill }),
    });

    const data = await response.json();

    return data;
  });
  return (
    <>
      <div class="min-h-[25vh] md:h-full md:w-1/3 rounded-2xl bg-white p-3">
        output box
        {input.value}
        <Resource value={result} onPending={() => <p>Loading...</p>} onResolved={(result) => <h2>{result && atob(result.result)}</h2>} onRejected={(error) => <p>Error: {error.message}</p>} />
      </div>
    </>
  );
});

function AppHeader() {
  return (
    <div class="flex justify-between">
      <div>
        <p class="text-xl">ICCBES Cipher</p>
        <p class="text-md">Next AES</p>
      </div>
      <div class="text-right text-sm">
        <div>
          <a href="https://github.com/christojeffrey">@christojeffrey</a>
        </div>
        {/* TODO: add adit and rafi */}
      </div>
    </div>
  );
}
export const head: DocumentHead = {
  title: "ICCBES",
  meta: [
    {
      name: "description",
      content: "ICCBES description",
    },
  ],
};
