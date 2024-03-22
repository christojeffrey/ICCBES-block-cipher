import { component$, useContext } from "@builder.io/qwik";
import { buttonContext, configContext, modeSpecificConfigContext } from "~/states";

export const ConfigBox = component$(() => {
  const button = useContext(buttonContext);
  const config = useContext(configContext);
  const modeSpecificConfig = useContext(modeSpecificConfigContext);
  // mode options = cbc, ecb, ofb, cfb
  // mode encryption options = encrypt, decrypt
  return (
    <>
      <div class="min-h-[25vh] md:h-full md:w-1/3 rounded-2xl bg-white p-3 flex flex-col gap-3">
        config box
        {/* block cipher mode */}
        <select
          class="border-2 border-black p-2 rounded-md"
          onClick$={(e) => {
            config.blockCipherMode = (e.target as HTMLSelectElement).value;
            console.log(config.blockCipherMode);
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
          onClick$={(e) => {
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
        <div>
          <label>Key</label>
          <input type="text" class="border-2 border-black p-2 rounded-md" onChange$={(e) => (modeSpecificConfig.key = (e.target as HTMLInputElement).value)} />
        </div>
        {/* mode specific config */}
        {config.blockCipherMode === "cbc" && <CBCConfig />}
        {config.blockCipherMode === "ecb" && <ECBConfig />}
        {config.blockCipherMode === "ofb" && <OFBConfig />}
        {config.blockCipherMode === "cfb" && <CFBConfig />}
        {config.blockCipherMode === "counter" && <CounterConfig />}
        {/* run button */}
        <button
          class="bg-blue-500 text-white p-2 rounded-md"
          onClick$={() => {
            button.pressed = false;
            button.pressed = true; // reset state
          }}
        >
          run!
        </button>
      </div>
    </>
  );
});

const CBCConfig = component$(() => {
  const modeSpecificConfig = useContext(modeSpecificConfigContext);

  return (
    <div>
      {/* title */}
      CBC Config
      {/* configs */}
      <div>
        <label>IV</label>
        <input type="text" class="border-2 border-black p-2 rounded-md" onChange$={(e) => (modeSpecificConfig.iv = (e.target as HTMLInputElement).value)} />
      </div>
    </div>
  );
});
const ECBConfig = component$(() => {
  return <div>ECB Config</div>;
});
const OFBConfig = component$(() => {
  return <div>OFB Config</div>;
});
const CFBConfig = component$(() => {
  return <div>CFB Config</div>;
});
const CounterConfig = component$(() => {
  return <div>Counter Config</div>;
});

export default ConfigBox;
